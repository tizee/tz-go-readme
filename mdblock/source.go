package mdblock

import (
	"encoding/json"
	"os"
)


type DataURL struct {
    Type string  `json:"type"`
    URL string `json:"url"`
}

func GetData(filepath string) ([]*DataURL,error)   {
    file, err := os.Open(filepath)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    var data []*DataURL
    err = json.NewDecoder(file).Decode(&data)

    return data, err
}