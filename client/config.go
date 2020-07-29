package client

import (
	"log"
	"net/http"
	"net/url"
	"time"
)

type RequestConfig struct {
    // URL string
    URL string
    // http requset method
    Method string
    // url
    BaseURL string
    // http URL parameters
    Params url.Values
    // timeout millseconds
    Timeout time.Duration
    // Custom Headers
    Headers http.Header
}

// http request config
type ClientConfig struct {
    // Http request 
    request *http.Request
    // Http response
    response *http.Response
    // base url for
	requestConfig *RequestConfig
	// http client
	httpClient *http.Client
    // result
    data map[string]interface{}
}

// mergeConfig: deep merge two request configs on the second config
func (config *RequestConfig) mergeConfig(other *RequestConfig) (*RequestConfig){
    // primitives types
    if other.BaseURL == ""{
        other.BaseURL = config.BaseURL
    }
    if other.Headers == nil {
        other.Headers = make(http.Header)
    }
    // negative and zero are invalid
    if other.Timeout <= 0{
        other.Timeout = config.Timeout
    }
    if other.URL == "" {
        other.URL = config.BaseURL 
    }
    // do not overwirte
    for key, vals := range config.Headers {
        for _, val := range vals {
            other.Headers.Add(key,val)
        }
    }
    fullURL,err := JoinParams(other.URL,other.Params)
    if err != nil {
        log.Fatalf("Invalid URL or parameters: %s",err)
    }
    other.URL = fullURL
    return other;
}