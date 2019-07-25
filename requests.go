package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// SendRequest make a request to retrieve the data from the API
func SendRequest(method, URL string, body interface{}) (data []byte) {
	client := &http.Client{}

	requestBody, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest(method, URL, bytes.NewBuffer(requestBody))
	if method != http.MethodDelete {
		req.Header.Set("Content-Type", "application/json")
	}

	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	data, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return
}
