package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type BasicResponse struct {
	Status  int                    `json:"status"`
	Success bool                   `json:"success"`
	Data    map[string]interface{} `json:"data"`
}

type Response struct {
	Url string `json:"url"`
}

func uploadAttachedFile(r *http.Request) ([]byte, error) {
	file, header, err := r.FormFile("media")
	if err != nil {
		return nil, err
	}

	defer file.Close()

	req, err := request(file, header.Filename)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var resp BasicResponse
	json.Unmarshal(body, &resp)

	link := resp.Data["link"].(string)

	if res.StatusCode != 200 || link == "" {
		err = errors.New("imgur response status code " + string(res.StatusCode))
		return nil, err
	}

	response := Response{}
	response.Url = link

	js, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}

	return js, nil
}
