package client

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

// basic Transformer
type Transformer func(body interface{}, headers http.Header) ([]byte, error)
// request body transformer
type RequestTransformer = Transformer
// response data transformer
type ReponseTransformer = Transformer
// user should handle response transformation in parser

// transform data
func transformData(body interface{}, headers http.Header, fns []Transformer) (interface{},error) {
	data := body
	for _, fn := range fns {
		res, err:= fn(data,headers)
		if err!=nil{
			return body, err
		}
		data = res
	}
	return data, nil
}

func basicRequestTransformer(body interface{},headers http.Header) ([]byte,error) {
	var res []byte
	switch body := body.(type) {
    // binary data
	case []byte:
		res = body
	// string
	case string:
		res = []byte(body)
    // stream buffer
	case bytes.Buffer:
		res = body.Bytes()
    // url parameters
	case url.Values:
		res = []byte(body.Encode())
		headers.Set(ContentType,ContentTypeURLParameters)
	// json
	default:
		res2,err := json.Marshal(body)
		headers.Set(ContentType,ContentTypeJSON)
		if err != nil {
			return nil,err
		}
		res = res2
	}
	return res,nil
}