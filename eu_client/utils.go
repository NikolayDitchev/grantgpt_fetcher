package eu_client

import (
	"bytes"
	"mime/multipart"
	"net/http"
)

func NewEURequest(method string, endpoint string,
	bodyOptions []RequestBodyOption,
	urlOptions []RequestURLOption) (*http.Request, error) {

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	for _, bodyOption := range bodyOptions {
		err := bodyOption(writer)
		if err != nil {
			return nil, err
		}
	}

	err := writer.Close()
	if err != nil {
		return nil, err
	}

	for _, urlOption := range urlOptions {
		endpoint, err = urlOption(endpoint)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, endpoint, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	return req, nil
}
