package mdblock

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
)

/*
Tag
start: <!-- Type-Start -->
end: <!-- Type-End -->
*/
func Tag(name string) []byte{
    return []byte("<!-- " + name + " -->")
}

/* 
WriteToMDFile
write results into markdown file
*/
func WriteToMDFile(results []*Result, filename string) {
// this function should run synchronously
    content, err := ioutil.ReadFile(filename)
    if err != nil {
        log.Fatalln("Invalid markdown file")
	}
	if len(results) == 0 {
        log.Println("empty results")
	} 
    // replace slots with newer content 
    for _,lines := range results {
        start,end:= Tag(lines.Type+"-start"), Tag(lines.Type+"-end")
        startIndx,endIndex := bytes.Index(content,start),bytes.Index(content,end)
        // use update
        beforeStart,afterEnd := content[:startIndx+len(start)] ,content[endIndex:] 
        updatedContent := bytes.NewBuffer(nil)
        // insert after <!-- xxx-start -->
        updatedContent.Write(beforeStart)
        updatedContent.WriteString("\n")
        updatedContent.Write(lines.Content)
        updatedContent.WriteString("\n")
        // insert before <!-- xxx-end -->
        updatedContent.Write(afterEnd)
        // update content
        content = updatedContent.Bytes()
    }
    err = ioutil.WriteFile(filename,content,os.ModeAppend)
    if err != nil {
        log.Fatalln("Can't wirte into " + filename)
    }
}