package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type booklogInfo struct {
	Tana struct {
		Account  string `json:"account"`
		Name     string `json:"name"`
		ImageUrl string `json:"image_url"`
	}
	Category []string `json:"category"`
	Books    []struct {
		Url     string `json:"url"`
		Title   string `json:"title"`
		Image   string `json:"image"`
		Catalog string `json:"catalog"`
	}
}

func main() {
	url := "http://api.booklog.jp/v2/json/vtryo"

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	jsonpStr := string(body)

	var s booklogInfo
	json.Unmarshal([]byte(jsonpStr), &s)
	fmt.Println(s)
}
