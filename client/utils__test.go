package client

import (
	"net/url"
	"testing"
)

func TestJoinParams(t *testing.T)  {
    want := "https://www.google.com?a=1&b=2"
    params := make(url.Values)
    params.Set("a","1")
    params.Set("b","2")
    if got,_ := JoinParams("https://www.google.com",params); got != want{
        t.Errorf("buildURL want %s but got %s",want,got)
    }
}

func TestJoinURL(t *testing.T)  {
    want := "https://www.google.com/query?a=1"
    if got := JoinBaseURL("https://www.google.com","/query?a=1"); got != want{
        t.Errorf("buildURL want %s but got %s",want,got)
    }
}