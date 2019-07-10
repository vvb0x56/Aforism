package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func getAforism(siteUrl string, values *url.Values) (string, error) {
	rs, err := http.PostForm(siteUrl, *values)
	if err != nil {
		return "", err
	}
	defer rs.Body.Close()

	bodyBytes, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		return "", err
	}

	bodyString := string(bodyBytes)

	return bodyString, nil
}

func main() {
	site := "http://api.forismatic.com/api/1.0/"
	vals := url.Values{}
	vals.Set("method", "getQuote")
	vals.Add("format", "text")
	vals.Add("lang", "ru")

	aforism, err := getAforism(site, &vals)
	if err != nil {
		panic(err)
	}
	fmt.Println(aforism)
}
