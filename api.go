package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type relatedData struct {
	RelatedTitle  string `json:"title"`
	RelatedAuthor string `json:"author"`
	RelatedLink   string `json:"link"`
	RelatedDate   string `json:"date"`
}

type data struct {
	Title   string        `json:"title"`
	Author  string        `json:"author"`
	Link    string        `json:"link"`
	Date    string        `json:"date"`
	Related []relatedData `json:"related"`
}

var m map[string][]data

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	resp, err := http.Get("http://127.0.0.1:1323/scrape")
	check(err)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	check(err)

	err = json.Unmarshal(body, &m)
	check(err)

	fmt.Println(m["business"][0].Related[0].RelatedTitle)
}
