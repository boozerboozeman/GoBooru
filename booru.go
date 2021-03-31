package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
)

type Booru struct {
	ImageUrl string `json:"file_url"`
}

func readJSONFromUrl(url string) ([]Booru, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var image_url []Booru
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	respByte := buf.Bytes()
	if err := json.Unmarshal(respByte, &image_url); err != nil {
		return nil, err
	}

	return image_url, nil
}

func main() {
	var query string
	flag.StringVar(&query, "query", "makise_kurisu", "your tag (danbooru format)")
	flag.Parse()
	url := "https://safebooru.donmai.us/posts.json?random=true&tags=" + query + "&rating=safe&limit=10"
	waifu, err := readJSONFromUrl(url)
	if err != nil {
		fmt.Println(nil, err)
	}
	for res := range waifu {
		fmt.Println(waifu[res].ImageUrl)
	}
}
