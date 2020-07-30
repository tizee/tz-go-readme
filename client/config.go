package client

import (
	"context"
	"net/http"
	"net/url"
	"time"
)

// Simple Response
type Response struct {
	Header http.Header
	StatusCode int
	Data []byte
}

type RequestConfig struct {
	// request context
	Context context.Context
    // url
    BaseURL string
    // URL string
    URL string
    // http requset method
    Method string
    // http URL parameters
    Params url.Values
    // timeout millseconds
    Timeout time.Duration
    // Custom Headers
	Headers http.Header
	// user should define the struct of Body
	Body interface{}
}

// http request config
type ClientConfig struct {
    // Http request 
	request *http.Request
	// Simplified Response
	response *Response
	// basic request config
	requestConfig *RequestConfig
	// http client
	httpClient *http.Client
}

// mergeConfig: deep merge two request configs on the second config
func mergeConfig(config1 *RequestConfig, config2 *RequestConfig) *RequestConfig{
	if config2 == nil {
		config2 = &RequestConfig{}
	}
	if config1 == nil {
		return config2
	}
    // primitives types
    if config2.BaseURL == ""{
        config2.BaseURL = config1.BaseURL
    }
    // negative and zero are invalid
    if config2.Timeout <= 0{
        config2.Timeout = config1.Timeout
	}
	if config2.BaseURL == "" {
		config2.BaseURL = config1.BaseURL
	}
    if config2.URL == "" {
        config2.URL = config1.BaseURL 
	}
	// reference types
    if config2.Headers == nil {
        config2.Headers = make(http.Header)
    }
    // do not overwirte
    for key, vals := range config1.Headers {
        for _, val := range vals {
            config2.Headers.Add(key,val)
        }
	}
	if config2.Params == nil {
		config2.Params = config1.Params
	}
	if config2.Body == nil {
		config2.Body = config1.Body
	}
	return config2
}