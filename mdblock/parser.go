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
Parse(source string) ([]byte,error) 
}

// Parse result
type Result struct {
    Type string
    Content []byte
}

// Run: entry point
func Run(filename string,datafile string)  {
    data, err := GetData(datafile)
    if err != nil {
        log.Fatalf("Get %s Error: %s\n",datafile,err)
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

        go func(parser Parser,Type string, src string) {
            ParserRunner(parser,Type,src,results)
            waitGroups.Done()
        }(parser,entry.Type,entry.Src)
    }

    go func() {
        waitGroups.Wait()
        close(results)
	}()
    // use slice instead of channel
	var lines = make([]*Result,0)
    for res := range results {
        lines = append(lines, &Result{
            Type: res.Type,
            Content: res.Content,
        }) 
    }

	// synchronous operation
    WriteToMDFile(lines,filename)
}

/*
ParserRunner: Run Parser
*/
func ParserRunner(p Parser,Type string, src string ,results chan<- *Result)  {
    log.Printf("Parse %s now \n",Type)
    res, err := p.Parse(src)
    if err != nil {
        log.Fatalln(err)
    }
    results <- &Result{
        Type: Type,
        Content: res,
    }
}