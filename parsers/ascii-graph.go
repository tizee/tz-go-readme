package parsers

import (
	"tz-go-readme/mdblock"

	"github.com/guptarohit/asciigraph"
)

type AsciiGraphParser struct {}

func init() {
    var parser AsciiGraphParser 
    mdblock.Register("ascii-graph",parser)
}

func (p AsciiGraphParser) Parse(url string) ([]byte, error)  {
    data := []float64{3, 4, 9, 6, 2, 4, 5, 8, 5, 10, 2, 7, 2, 5, 6}
    graph := asciigraph.Plot(data)
    return []byte(graph), nil
}