package mdblock

import (
	"encoding/json"
	"os"
)

const dataFile = "./data.json"

type DataURL struct {
    Type string  `json:"type"`
    URL string `json:"url"`
}

func GetData() ([]*DataURL,error)   {
    file, err := os.Open(dataFile)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    var data []*DataURL
    err = json.NewDecoder(file).Decode(&data)

    return data, err
}