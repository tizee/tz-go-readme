package client

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
        return "", err
    }
	res := urlObj.String()
	n := len(Params)
    if n > 0{
		count := 0
        res += "?"
        for key, vals := range Params {
            for _, val := range vals {
				if count < n-1{
				res += key + "=" + val + "&"
				}else{
				res += key + "=" + val
				}
				count += 1
            }
        }
    }
    return res, nil
}
