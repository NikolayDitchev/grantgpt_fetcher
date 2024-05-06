package api_caller

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"strconv"
)

const (
	METHOD       = "POST"
	API_ENDPOINT = "https://api.tech.ec.europa.eu/search-api/prod/rest/search"
)

type API_Caller struct {
	url        *url.URL
	urlParams  url.Values
	bodyParams map[string][]byte

	pageNumber int
	pageSize   int
}

func NewAPI_Caller(bodyParams map[string][]byte, urlParams url.Values) (apc *API_Caller, err error) {
	apc = &API_Caller{
		urlParams:  urlParams,
		bodyParams: bodyParams,
		pageNumber: 1,
	}

	apc.url, err = url.Parse(API_ENDPOINT)
	if err != nil {
		return nil, err
	}

	apc.pageSize, err = strconv.Atoi(apc.urlParams.Get("pageSize"))
	if err != nil {
		return nil, err
	}

	return
}

func (apc *API_Caller) GetResults(resultsChan chan<- []Result) {

	apc.urlParams.Set("pageNumber", strconv.Itoa(apc.pageNumber))
	apc.url.RawQuery = apc.urlParams.Encode()

	client := &http.Client{}

	res, err := client.Do(apc.createRequest())
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	jsonResponse := &ResponseBody{}

	err = json.Unmarshal(body, jsonResponse)
	if err != nil {
		log.Fatalln(err)
	}

	resultsChan <- jsonResponse.Results

	if apc.pageNumber*apc.pageSize >= jsonResponse.TotalResults {
		close(resultsChan)
		return
	}

	apc.pageNumber++
	apc.GetResults(resultsChan)
}

func (apc *API_Caller) createRequest() *http.Request {

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	for key, value := range apc.bodyParams {

		var header textproto.MIMEHeader = make(textproto.MIMEHeader)
		header.Add("Content-Disposition", `form-data; name="`+key+`";`)
		header.Add("Content-Type", "application/json")

		part, err := writer.CreatePart(header)
		if err != nil {
			log.Fatalln(err)
		}

		part.Write(value)
	}

	err := writer.Close()
	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest(METHOD, apc.url.String(), payload)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	return req
}
