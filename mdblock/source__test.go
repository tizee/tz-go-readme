package mdblock

import (
	"testing"
)

func TestData(t *testing.T) {
    var want []*DataURL
    want = append(want, &DataURL{
        Type: "wakatime",
        URL: "https://wakatime.com/api/v1/users/current/stats/last_7_days",
    },&DataURL{
        Type: "rss",
        URL: "https://tizee.github.io/rss.xml",
    })
    if got,err := GetData("../data.json"); got==nil {
        t.Errorf("invalid data.json %s",err)
    }
}