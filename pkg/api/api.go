package api

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

type Response struct {
	StatusCode int         `json:"statusCode"`
	Body       interface{} `json:"body"`
}

func GetRequest(endpoint string) (*Response, error) {
	/*
		http get request wrapper
	*/

	var response Response

	client := &http.Client{}

	req, err := http.NewRequest("GET", endpoint, nil)

	if err != nil {
		return &response, errors.New(fmt.Sprintf("ERROR : GetRequest : %v", err.Error()))
	}

	req.Header.Set("Accept", "application/vnd.heroku+json; version=3")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", os.Getenv("HEROKU_API_KEY")))

	resp, err := client.Do(req)

	if err != nil {
		return &response, errors.New(fmt.Sprintf("ERROR : GetRequest : %v", err.Error()))
	}

	var body interface{}

	json.NewDecoder(resp.Body).Decode(&body)

	response.StatusCode = resp.StatusCode
	response.Body = body

	return &response, nil
}

func PostRequest(endpoint string, payload string) (*Response, error) {
	/*
		http post request wrapper
	*/

	var response Response

	client := &http.Client{}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer([]byte(payload)))

	if err != nil {
		return &response, errors.New(fmt.Sprintf("ERROR : PostRequest : %v", err.Error()))
	}

	req.Header.Set("Accept", "application/vnd.heroku+json; version=3")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", os.Getenv("HEROKU_API_KEY")))

	resp, err := client.Do(req)

	if err != nil {
		return &response, errors.New(fmt.Sprintf("ERROR : GetRequest : %v", err.Error()))
	}

	var body interface{}

	json.NewDecoder(resp.Body).Decode(&body)

	response.StatusCode = resp.StatusCode
	response.Body = body

	return &response, nil
}

func StreamRequest(endpoint string, buffer chan string, apiSignal chan bool) error {
	/*
		http get request with output stream
	*/

	client := &http.Client{}

	req, err := http.NewRequest("GET", endpoint, nil)

	if err != nil {
		return errors.New(fmt.Sprintf("ERROR : StreamRequest : %v", err.Error()))
	}

	resp, err := client.Do(req)
	if err != nil {
		return errors.New(fmt.Sprintf("ERROR : StreamRequest : %v", err.Error()))
	}

	reader := bufio.NewReader(resp.Body)

	go func() {
		select {
		case <-apiSignal:
			resp.Body.Close()
		}
	}()

	for {
		line, err := reader.ReadBytes('\n')

		if err != nil {
			return errors.New(fmt.Sprintf("ERROR : StreamRequest : %v", err.Error()))
		}

		buffer <- string(line)
	}
}

// func main() {
// os.Setenv("HEROKU_API_KEY", "00a2fdf8-66c8-4925-b55e-d87fd79bc803")

// resp, err := GetRequest("https://api.heroku.com/apps")
// if err != nil {
// fmt.Println(err)
// }

// fmt.Println(*resp)
// }
