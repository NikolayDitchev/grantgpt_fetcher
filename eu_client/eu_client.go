package eu_client

import (
	"bytes"
	"encoding/json"
	"errors"
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

	uniqueResults map[string]int
}

func NewAPI_Caller(timeout time.Duration) (apc *API_Caller) {
	apc = &API_Caller{
		client: &http.Client{
			Timeout: timeout,
		},

		uniqueResults: make(map[string]int),
	}

	apc.url, _ = url.Parse(API_ENDPOINT)

	return
}

func (apc *API_Caller) Reset() {
	clear(apc.uniqueResults)
}

/*

	The pagination of the EU funding and tenders API is broken.

	If the results you want to fetch are more that the pageSize,
	then it duplicates results which some results not present in any page.
	This behaviour is unidentified, so the function below is requesting the same pages
	multiple times, hoping that every unique result will be fetched.

	If the fetched unique results don't match the totalResults value,
	the function will return an error.

*/

// func (apc *API_Caller) GetUIDs(query Query, uidKey string, tries int) error {
// 	UIDsChan := make(chan string)
// 	errChan := make(chan error)

// 	queryJSON, err := json.Marshal(query)
// 	if err != nil {
// 		return err
// 	}

// 	var bodyParams map[string][]byte = map[string][]byte{
// 		"query":     queryJSON,
// 		"languages": []byte(`["en"]`),
// 	}

// 	var urlParams url.Values = url.Values{
// 		"apiKey":     []string{"SEDIA"},
// 		"text":       []string{"***"},
// 		"pageSize":   []string{"100"},
// 		"pageNumber": []string{"1"},
// 	}

// 	url, _ := url.Parse(API_ENDPOINT)
// 	url.RawQuery = urlParams.Encode()

// 	IDsMap := make(map[string]int)
// 	totalResults := 0

// 	for i := 0; i < tries; i++ {

// 		pageChan := make(chan *Page)

// 		go func() {

// 			err := apc.getPages(bodyParams, url, pageChan)
// 			if err != nil {
// 				errChan <- err
// 			}

// 			close(pageChan)
// 		}()

// 		for page := range pageChan {
// 			totalResults = page.TotalResults

// 			for inx := range page.Results {

// 				id := page.Results[inx].Metadata[uidKey][0]

// 				if _, exists := IDsMap[id]; !exists {

// 					UIDsChan <- id
// 					IDsMap[id] = 1
// 				}
// 			}
// 		}

// 		fmt.Println(len(IDsMap))

// 		if len(IDsMap) >= totalResults {
// 			close(UIDsChan)
// 			close(errChan)
// 			return nil
// 		}
// 	}

// 	close(UIDsChan)
// 	errChan <- errors.New("not every item was fetched")
// 	close(errChan)

// 	return nil
// }

func (apc *API_Caller) GetTopicIDs(topicIDsChan chan string, errChan chan error) {
	defer close(errChan)

	query := NewQuery(WithTypes(TypeTopics), WithStatus(StatusOpen, StatusForthcoming))
	queryJson, err := json.Marshal(query)
	if err != nil {
		errChan <- err
		return
	}

	var bodyParams map[string][]byte = map[string][]byte{
		"query":     queryJson,
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

	topicIDsMap := make(map[string]int)
	totalResults := 0

	for i := 0; i < 50; i++ {

		pageChan := make(chan *Page)

		go func() {

			err := apc.getPages(bodyParams, url, pageChan)
			if err != nil {
				errChan <- err
			}

			close(pageChan)
		}()

		for page := range pageChan {
			totalResults = page.TotalResults

			for inx := range page.Results {

				id := page.Results[inx].Metadata["identifier"][0]

				if _, exists := topicIDsMap[id]; !exists {

					topicIDsChan <- id
					topicIDsMap[id] = 1
				}
			}
		}

		fmt.Println(len(topicIDsMap))

		if len(topicIDsMap) >= totalResults {
			close(topicIDsChan)
			return
		}
	}

	close(topicIDsChan)
	errChan <- errors.New("not every topic was fetched")
}

func (apc *API_Caller) getPages(bodyParams map[string][]byte, url *url.URL, pageChan chan *Page) error {

	body, err := apc.sendRequest(bodyParams, url.String())
	if err != nil {
		return err
	}

	page := &Page{}

	err = json.Unmarshal(body, page)
	if err != nil {
		return err
	}

	pageChan <- page

	if page.PageSize*page.PageNumber >= page.TotalResults {
		return err
	}

	err = apc.increasePageNumber(url)
	if err != nil {
		return err
	}

	apc.getPages(bodyParams, url, pageChan)

	return nil
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
