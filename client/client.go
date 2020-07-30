package client

import (
	"bytes"
	"context"
	"log"
	"net/http"
	"tz-go-readme/client/util"
)

type Client struct{
    config *ClientConfig
}

// a really simple http client based on go's standard libraries
var DefaultClient *Client

func init()  {
    // setup default http client
    DefaultClient = NewClient(nil)
}

// create a http client
func NewClient(config *ClientConfig) *Client {
    if config == nil {
        // using default config
        config = &ClientConfig{}
    }
    return &Client{
        config: config,
    }
}

// request builder
func RequestFactory(config *RequestConfig)(*http.Request, error)  {
	// 1. build url
    requestUrl := util.JoinBaseURL(config.BaseURL,config.URL)
	requestUrl,err := util.JoinParams(requestUrl,config.Params)
	if err != nil {
		log.Fatalf("%s is not a valid URL",requestUrl)
	}
	// 2. setup method
	if config.Method == ""{
		config.Method = http.MethodGet
	}
	// 3. transform body to []byte with default request transformers
	bodyData, err := transformData(config.Body,config.Headers,DefaultRequestTransformers)
	if err != nil {
		log.Fatalf("invalid body, %s",err)
	}
	reader := bytes.NewReader(bodyData.([]byte))
	request, err := http.NewRequest(config.Method,requestUrl,reader) 
	if config.Timeout != 0 {
		ctx := config.Context
		if ctx == nil {
			ctx = context.Background()
		}
		ctx, cancel := context.WithTimeout(ctx,config.Timeout)
		defer cancel()
		config.Context = ctx
	}
	// setup Context
	if config.Context != nil {
		request = request.WithContext(config.Context)
	}
	// setup Header
	for key,vals := range config.Headers{
		for _, val := range vals {
			request.Header.Add(key,val)
		}
	}
	// setup user-agent
	request.Header.Set(HeaderUserAgent,UserAgent)
	// set accept Encoding
	request.Header.Set(HeaderAccept,DefaultAcceptType)
	return request,err
}

    // merge with client's request config
func (instance *Client) Request(config *RequestConfig) (*Response,error) {
    // 1. merge config
    mergeConfig(instance.config.requestConfig,config)
	// 2. create http request using request config
	request, err := RequestFactory(config)
	if err != nil {
		log.Fatalf("Create Request failed: %s", err)
	}
	instance.config.request = request
	// 3. send
	resp, err := defaultAdaptor(instance.config) 
	return resp,err
}

func (instance *Client) Get(url string,config *RequestConfig)  (*Response,error){
	getConfig := &RequestConfig{
		URL: url,
		Method: http.MethodGet,
	}
	getConfig = mergeConfig(config,getConfig)
	return instance.Request(getConfig)
}

func (instance  *Client) Post(url string,data interface{},config *RequestConfig)(*Response,error){
	postConfig := &RequestConfig{
		URL: url,
		Method: http.MethodPost,
		Body: data,
	}
	postConfig= mergeConfig(config,postConfig)
	return instance.Request(postConfig)
}
