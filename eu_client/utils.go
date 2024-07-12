package eu_client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
)

func GetPages(req *http.Request, client *http.Client) ([]*Page, error) {

	pages := make([]*Page, 0)

	for i := 0; i < maxPageNumber; i++ {

		temp := req.Clone(context.Background())
		tempBody, err := req.GetBody()
		if err != nil {
			return nil, err
		}
		temp.Body = tempBody

		resp, err := client.Do(temp)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		page := &Page{}
		err = json.Unmarshal(body, page)
		if err != nil {
			return nil, err
		}
		pages = append(pages, page)

		if page.PageSize*page.PageNumber >= page.TotalResults {
			return pages, nil
		}

		err = increasePageNumber(req)
		if err != nil {
			return nil, err
		}
	}

	return pages, nil
}

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

func increasePageNumber(req *http.Request) error {
	urlParams := req.URL.Query()
	pageNumber, err := strconv.Atoi(urlParams.Get("pageNumber"))
	if err != nil {
		return err
	}
	pageNumber++
	urlParams.Set("pageNumber", strconv.Itoa(pageNumber))
	req.URL.RawQuery = urlParams.Encode()
	return nil
}

func GetMetadataField(result *Result, field string) (string, error) {

	fieldArr, ok := result.Metadata[field]
	if !ok {
		return "", errors.New("metadata field not found")
	}

	if len(fieldArr) == 0 {
		return "", errors.New("metadata field is empty")
	}

	return fieldArr[0], nil
}
