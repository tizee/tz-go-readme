package mdblock

import (
	"testing"
)

func TestData(t *testing.T) {
    want := []*Source{{Type: "wakatime",
        Src: "https://wakatime.com/api/v1/users/current/stats/last_7_days",
    },{
        Type: "rss",
        Src: "https://tizee.github.io/rss.xml",
    }}
    if got,err := GetData("../data.json"); got==nil {
        t.Errorf("GetData want %d but got %d, err: %s",len(want),len(got),err)
    }
}