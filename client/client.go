package client

import "net/http"

type Client struct{
    config *ClientConfig
	id uint64
}

// a really simple http client based on go's standard libraries
var defaultClient *Client

func init()  {
    // setup default http client
    defaultClient = NewClient(nil)
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

    // merge with client's request config
func (instance *Client) Request(config *RequestConfig) (*http.Response,error) {
    config.mergeConfig(instance.config.requestConfig)
    // 1. merge config
    // 2. set config method
	// 3. send request
	return nil,nil
}

func (instance *Client) Get(url string,config *RequestConfig)  (*http.Response,error){
	return nil,nil
}

func (instance  *Client) Post(url string,config *RequestConfig)(*http.Response,error){
	return nil,nil
}


