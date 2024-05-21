package api_caller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"strconv"
	"time"
)

const (
	METHOD       = "POST"
	API_ENDPOINT = "https://api.tech.ec.europa.eu/search-api/prod/rest/search"
)

type API_Caller struct {
	client *http.Client
	url    *url.URL
}

func NewAPI_Caller() (apc *API_Caller, err error) {
	apc = &API_Caller{
		client: &http.Client{
			Timeout: 4 * time.Second,
		},
	}

	apc.url, _ = url.Parse(API_ENDPOINT)

	return
}

func (apc *API_Caller) GetTopicIDs() (topicIDs chan string) {

	var bodyParams map[string][]byte = map[string][]byte{
		"query":     GetTopicQuery(),
		"languages": []byte(`["en"]`),
	}

	var urlParams url.Values = url.Values{
		"apiKey":     []string{"SEDIA"},
		"text":       []string{"***"},
		"pageSize":   []string{"100"},
		"pageNumber": []string{"1"},
	}

	url, _ := url.Parse(API_ENDPOINT)
	url.RawQuery = urlParams.Encode()

	topicIDs = make(chan string)

	go func() {

		topicIDsMap := make(map[string]int)
		totalResults := 0

		for {
			pagesChan, _ := apc.getPages(bodyParams, url)

			for page := range pagesChan {
				totalResults = page.TotalResults

				for inx := range page.Results {

					id := page.Results[inx].Metadata["identifier"][0]

					if _, exists := topicIDsMap[id]; !exists {
						topicIDs <- id
						topicIDsMap[id] = 1
					}
				}
			}

			if len(topicIDsMap) >= totalResults {
				close(topicIDs)
				return

			}

			fmt.Println(len(topicIDsMap))

			if len(topicIDsMap) < totalResults {
				time.Sleep(3 * time.Second)
			}

		}
	}()

	return topicIDs
}

func (apc *API_Caller) getPages(bodyParams map[string][]byte, url *url.URL) (pageChan chan *Page, errorsChan chan error) {

	pageChan = make(chan *Page)
	errorsChan = make(chan error)

	go func() {

		for {

			body, err := apc.sendRequest(bodyParams, url.String())
			if err != nil {
				errorsChan <- err
				return
			}

			page := &Page{}

			err = json.Unmarshal(body, page)
			if err != nil {
				errorsChan <- err
				return
			}

			pageChan <- page

			if page.PageSize*page.PageNumber >= page.TotalResults {
				close(pageChan)
				break
			}

			err = apc.increasePageNumber(url)
			if err != nil {
				errorsChan <- err
				return
			}
		}
	}()

	return
}

func (apc *API_Caller) sendRequest(bodyParams map[string][]byte, url string) ([]byte, error) {

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	for key, value := range bodyParams {

		var header textproto.MIMEHeader = make(textproto.MIMEHeader)
		header.Add("Content-Disposition", `form-data; name="`+key+`";`)
		header.Add("Content-Type", "application/json")

		part, err := writer.CreatePart(header)
		if err != nil {
			return nil, err
		}

		part.Write(value)
	}

	err := writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(METHOD, url, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := apc.client.Do(req)
	if err != nil {
		return nil, err
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}

	return respBody, nil
}

func (apc *API_Caller) increasePageNumber(url *url.URL) error {

	urlParams := url.Query()
	pageNumber, err := strconv.Atoi(urlParams.Get("pageNumber"))
	if err != nil {
		return err
	}

	pageNumber = pageNumber + 1

	urlParams.Set("pageNumber", strconv.Itoa(pageNumber))
	url.RawQuery = urlParams.Encode()

	return nil
}
