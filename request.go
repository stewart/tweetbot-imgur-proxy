package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
)

func request(file multipart.File, name string) (*http.Request, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", name)

	if err != nil {
		return nil, err
	}

	_, err = io.Copy(part, file)
	err = writer.Close()

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", ENDPOINT, body)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Client-ID "+CLIENT_ID)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	return req, nil
}
