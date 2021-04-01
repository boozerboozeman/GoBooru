package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
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

func downloadFile(url string, path string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	return err
}

func main() {
	var query string
	var limit string
	flag.StringVar(&query, "query", "makise_kurisu", "your tag (danbooru format)")
	flag.StringVar(&limit, "limit", "1", "the number of picture you want to see")
	flag.Parse()
	url := "https://safebooru.donmai.us/posts.json?random=true&tags=" + query + "&rating=safe&limit=" + limit
	waifu, err := readJSONFromUrl(url)
	if err != nil {
		fmt.Println(nil, err)
	}
	for res := range waifu {
		fmt.Println("Downloading:" + waifu[res].ImageUrl)
		fileUrl := waifu[res].ImageUrl
		temp := strings.Split(fileUrl, "/")
		fileName := temp[len(temp)-1]
		downloadFile(waifu[res].ImageUrl, "./"+fileName)
	}
}
