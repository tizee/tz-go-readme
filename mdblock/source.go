package mdblock

import (
	"encoding/json"
	"os"
)

type Source struct {
    Type string  `json:"type"`
    Src string `json:"source"`
}

// Retrieve data from json file
func GetData(filepath string) ([]*Source,error)   {
    file, err := os.Open(filepath)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    var data []*Source
    err = json.NewDecoder(file).Decode(&data)

    return data, err
}