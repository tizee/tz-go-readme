package mdblock

import (
	"log"
	"sync"
)

// global parsers
var parsers = make(map[string]Parser)

// Register to mdblock
func Register(parserType string, parser Parser)  {
    if _, exists := parsers[parserType]; exists {
        log.Fatalln(parserType,"Parser already registered")
    }
    log.Println("Register", parserType, "parser")
    parsers[parserType] = parser
}

type Parser interface{
// transform input string
Parse(url string) ([]byte,error) 
}

// Parse result
type Result struct {
    Type string
    Content []byte
}

// Run: entry point
func Run(filename string)  {
    data, err := GetData()
    if err != nil {
        log.Fatalln("Get data.json Error:",err)
    }
    results := make(chan *Result)

    var waitGroups sync.WaitGroup
    waitGroups.Add(len(data))

    for _, entry := range data{
        parser, exists := parsers[entry.Type]
        if !exists {
            // skip
            parser = parsers["default"]
        }

        go func(parser Parser,Type string, url string) {
            ParserRunner(parser,Type,url,results)
            waitGroups.Done()
        }(parser,entry.Type,entry.URL)
    }

    go func() {
        waitGroups.Wait()
        close(results)
    }()
    var lines = make([]*Result,0)
    // use slice instead of channel
    for res := range results {
        lines = append(lines, &Result{
            Type: res.Type,
            Content: res.Content,
        }) 
    }
    WriteToMDFile(lines,filename)
}

/*
ParserRunner: Run Parser
*/
func ParserRunner(p Parser,Type string, url string ,results chan<- *Result)  {
    log.Printf("Parse %s now \n",Type)
    res, err := p.Parse(url)
    if err != nil {
        log.Fatalln(err)
    }
    results <- &Result{
        Type: Type,
        Content: res,
    }
}