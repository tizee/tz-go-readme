package parsers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"tz-go-readme/client"
	"tz-go-readme/mdblock"
	"unicode/utf8"
)

type wakaParser struct{}

func init() {
    var parser wakaParser
    mdblock.Register("wakatime",parser)
}
/*
see https://wakatime.com/developers#stats
	{
"machine_name_id": <string: unique id of this machine>,
"date": <string: day with most coding time logged as Date string in YEAR-MONTH-DAY format>,
"name": <string: language name>,
"total_seconds": <float: total coding activity spent in this language as seconds>,
"percent": <float: percent of time spent in this language>,
"digital": <string: total coding activity for this language in digital clock format>,
"text": <string: total coding activity in human readable format>,
"hours": <integer: hours portion of coding activity for this language>,
"minutes": <integer: minutes portion of coding activity for this language>,
"seconds": <integer: seconds portion of coding activity for this language>
	}
*/

// wakatime api response

type WakaStatDataItem struct {
	Name *string `json:"name,omitempty"`
	TotalSeconds *float64 `json:"total_seconds,omitempty"`
	Percent *float64 `json:"percent,omitempty"`
	Digital *string `json:"digital,omitempty"`
	Text *string `json:"text,omitempty"`
	Hours *uint8 `json:"hours,omitempty"`
	Minutes *uint8 `json:"minutes,omitempty"`
	Seconds *uint8 `json:"seconds,omitempty"`
	Date *string `json:"data,omitempty"`
	MachineNameId *string `json:"machine_name_id,omitempty"`
}

type WakaData struct {
	TotalSeconds *float64 `json:"total_seconds,omitempty"`
	HumanReadableTotal *string `json:"human_readable_total,omitempty"`
	DailyAverage *float64 `json:"daily_average,omitempty"`
	HumanReadableDailyAverage *string `json:"human_readable_daily_average,omitempty"`
	Languages []*WakaStatDataItem `json:"languages,omitempty"`
	Machines []*WakaStatDataItem `json:"machines,omitempty"` 
	Dependencies []*WakaStatDataItem  `json:"dependencies,omitempty"`
	OperatingSystem []*WakaStatDataItem `json:"operating_systems,omitempty"`
	Editors []*WakaStatDataItem `json:"editors,omitempty"`
	Projects []*WakaStatDataItem `json:"projects,omitempty"`

}

type WakaJSON struct {
	Data WakaData `json:"data,omitempty"`
}


// fetch and parse wakatime  
func (p wakaParser) Parse(src string) ([]byte,error) {
	resp, err := client.DefaultClient.Get(src,&client.RequestConfig{
		Params: map[string][]string{"api_key":{os.Getenv("WAKATIME_APIKEY")}},
	})
	if err != nil {
		log.Fatalf("[Parse]: wakatime error %s",err)
	}
	var stat WakaJSON
	err = json.Unmarshal(resp.Data,&stat)
	if err !=nil {
		log.Fatalf("[Parse]: parse jsong %s",err)
	}
	lines := getWakaBox(&stat)
    return lines,nil
}

type BoxLine struct {
	Name string
	Percent float32
	Time string
}

/* wakatime-box */
/*
Language1 time [====...] percent
Language2 time [====...] percent
Language3 time [====...] percent
Language4 time [====...] percent
*/

const (
	WakaTimeNameMaxWidth = 12 // e.g TypeScript 
	WakaTimeTimeMaxWidth = 6 // e.g 23h59m
	WakaTimePerceMaxWidth = 5 // e.g 99.9%
)

// createBarChart
// size: length of the bar 
func createBarChart(percent float64, size int) string  {
	// 2 state
	syms := []rune(`•=`)
	frac := math.Floor((float64(size) * 2 * percent)/100)
	barFull := int(math.Floor(frac / 2))
	// check whether is full
	if barFull >= size {
		return strings.Repeat(string(syms[1:2]),size)
	}
	// no other semi-state symbol
	barStr := strings.Repeat(string(syms[1:2]),barFull) 
	return  padEnd(barStr,string(syms[0:1]),size)
}


func padEnd(src string,pattern string,size int) string {
		repeatNum := size - utf8.RuneCountInString(src)
		if repeatNum < 0 {
			return src 
		}
		// need padding
		return src + strings.Repeat(pattern,repeatNum)
}


// The same notation before a '*' for a width or precision selects the argument index holding the value.
func createWakaBoxLine(data *WakaStatDataItem) string {
	name,hours,minutes,percent := data.Name,data.Hours,data.Minutes,data.Percent
	// time
	timeStr := fmt.Sprintf("%dh%dm",*hours,*minutes)
	// percent
	perceStr := fmt.Sprintf("%4.1f%%",*percent)
	return fmt.Sprintf("%-*s %-*s [%s] %s",WakaTimeNameMaxWidth,*name,WakaTimeTimeMaxWidth,timeStr,createBarChart(*percent,21),perceStr)
}

func getRankSymbol(num int)string  {
	switch num {
	case 0:
		return "🦄"
	case 1:
		return "🥇"
	case 2:
		return "🥈"
	case 3:
		return "🥉"
	default:
		return "🙏"		
	}	
}

func getWakaBox(data *WakaJSON) []byte {
	// do whatever you want with the data
	// only print 4 languages
	buf := bytes.NewBuffer(nil)
	// gang of four languages 🚀
	maxLen := 4
	if maxLen > len(data.Data.Languages) {
		maxLen = len(data.Data.Languages)
	}
	gof := data.Data.Languages[:maxLen]
	// Yooo~~
	buf.WriteString("🖖 Code long and prosper\n")	
	buf.WriteString("```text\n")
	for i, val := range gof {
		line := createWakaBoxLine(val)
		buf.WriteString(getRankSymbol(i)+" "+line+"\n")
	}
	buf.WriteString("```\n")
	return buf.Bytes()
}