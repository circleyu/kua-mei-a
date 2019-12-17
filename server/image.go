package main

import (
	"encoding/json"
	"net/http"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

// ImageData ImageData
type ImageData struct {
	Code   string
	Imgurl string
}

func getImangURL() (string, string) {
	data := ImageData{}
	getJSON("https://api.ooopn.com/image/beauty/api.php?type=json", &data)
	return data.Imgurl, data.Code
}

func getJSON(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
