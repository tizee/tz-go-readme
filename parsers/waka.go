package parsers

import "tz-go-readme/mdblock"

type wakaParser struct{}

func init() {
    var parser wakaParser
    mdblock.Register("wakatime",parser)
}

// fetch and parse wakatime  
func (p wakaParser) Parse(url string) ([]byte,error) {
    return []byte("wakatime"),nil
}
