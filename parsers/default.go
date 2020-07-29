package parsers

import "tz-go-readme/mdblock"


type DefaultParser struct {
}

func (p DefaultParser) Parse(url string) ([]byte,error)  {
    return nil,nil
}

func init()  {
    var parser DefaultParser 
    mdblock.Register("default",parser)
}