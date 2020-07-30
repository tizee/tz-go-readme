package client

import (
	"io/ioutil"
	"net/http"
)

// adaptor for sending requests using http.client

type Adaptor func(config *ClientConfig) (*http.Response,error)


func defaultAdaptor(config *ClientConfig)(*Response,error) {
	req,client := config.request,config.httpClient
	if client == nil {
		client = http.DefaultClient
		// update
		config.httpClient = client
	}
	res, err  := client.Do(req)
	if err != nil {
		return nil, err
	}
	response := &Response{
		StatusCode: res.StatusCode,
		Header: res.Header,
	}
	response.Data,err = ioutil.ReadAll(res.Body)
	defer res.Body.Close() 
	if err != nil {
		return nil, err
	}
	config.response = response
	return response,nil
}
