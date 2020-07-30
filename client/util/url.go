package util

import (
	"net/url"
	"strings"
)

func JoinBaseURL(BaseURL string, URL string) string  {
	if BaseURL == "" {
		return URL
	}
	if strings.HasPrefix(URL,"https://") || strings.HasPrefix(URL,"http://"){
		return URL
	}
	if strings.HasSuffix(BaseURL,"/") && strings.HasPrefix(URL,"/") {
		URL = URL[1:]
	}
	return BaseURL + URL 
}

func JoinParams(URL string, Params url.Values) (string,error)  {
    urlObj,err := url.Parse(URL)
    if err != nil {
        return URL, err
    }
	res := urlObj.String()
	n := len(Params)
    if n > 0{
		query := Params.Encode(); 
		if strings.HasSuffix(res,"?") {
			res += query
		}else {
			res += "?"+ query
		}
    }
    return res, nil
}
