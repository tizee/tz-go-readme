package parsers

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"time"
	"tz-go-readme/client"
	"tz-go-readme/mdblock"
)

// only need title, description, link, time for displaying on Github
type (
    // <item>
    item struct{
        XMLNAME xml.Name `xml:"item"`
        PubDate *string `xml:"pubDate"`
        Description *string `xml:"description"`
        Title *string `xml:"title"`
        Link *string `xml:"xlink"`
        URL *string `xml:"url"`
        GUID *string `xml:"guid"`
    }
    // <image>
    image struct{
        XMLNAME xml.Name `xml:"image"`
        Title *string `xml:"title"`
        Link *string `xml:"xlink"`
        URL *string `xml:"url"`
    }
    // <channel>
    channel struct{
        XMLNAME xml.Name `xml:"channel"`
        Title *string `xml:"title"`
        Description *string `xml:"description"`
        Link *string `xml:"link"`
        PubDate *string `xml:"pubDate"`
        Image *image `xml:"image"`
        Item []*item `xml:"item"`
    }
    // <rss>
    rssDocument struct{
        XMLNAME xml.Name `xml:"rss"`
        Channel *channel `xml:"channel"`
    }
)

type rssParser struct {}

func init() {
    var parser rssParser
    mdblock.Register("rss",parser)
}

func createAnchorLink(link *string,txt *string) string {
	return "<a href=\"" + *link + "\" target=\"_blank\">" + *txt + "</a>" 
}

// fetch and parse rss 
func (p rssParser) Parse(src string) ([]byte,error) {
	resp, err := client.DefaultClient.Get(src,&client.RequestConfig{})
	if err != nil {
		log.Fatalf("[Parser]: Error on get %s",src)
	}
	var doc rssDocument
	err = xml.Unmarshal(resp.Data,&doc)
	if err != nil {
		log.Fatalf("[Parser]: Error on decoding xml %s, error: %s",src,err)
	}
	// latest 5, gang of five LOL
	maxLen := 5
	if maxLen > len(doc.Channel.Item){
		maxLen = len(doc.Channel.Item)
	}
	posts := doc.Channel.Item[:maxLen]
	// display("posts",reflect.ValueOf(posts))
	var buf bytes.Buffer
	for _, post := range posts {
		log.Printf("%s - %s - %s\n",*post.Title,*post.PubDate,*post.GUID)
		t, err := time.Parse(http.TimeFormat,*post.PubDate)
		if err != nil {
			log.Fatalf("Error on format %s",err)
		}
		timeStr := t.Format(time.RFC3339)[:10] 
		buf.WriteString("- " + fmt.Sprintf("%-s - %-s\n",createAnchorLink(post.GUID,post.Title), timeStr))
	}
    return buf.Bytes(),nil
}
