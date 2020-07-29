package parsers

import (
	"encoding/xml"
	"tz-go-readme/mdblock"
)

// only need title, description, link, time for displaying on Github
type (
    // <item>
    item struct{
        XMLNAME xml.Name `xml:"item"`
        PubDate string `xml:"pubDate"`
        Description string `xml:"description"`
        Title string `xml:"title"`
        Link string `xml:"xlink"`
    }
    // <image>
    image struct{
        XMLNAME xml.Name `xml:"image"`
        Title string `xml:"title"`
        Link string `xml:"xlink"`
        URL string `xml:"url"`
    }
    // <channel>
    channel struct{
        XMLNAME xml.Name `xml:"channel"`
        Title string `xml:"title"`
        Description string `xml:"description"`
        Link string `xml:"link"`
        PubDate string `xml:"pubDate"`
        Image image `xml:"image"`
        Item []item `xml:"item"`
    }
    // <rss>
    rssDocument struct{
        XMLNAME xml.Name `xml:"rssDocument"`
        Channel channel `xml:"channel"`
    }
)

type rssParser struct {}

func init() {
    var parser rssParser
    mdblock.Register("rss",parser)
}

// fetch and parse wakatime  
func (p rssParser) Parse(url string) ([]byte,error) {
    return []byte("rss"),nil
}
