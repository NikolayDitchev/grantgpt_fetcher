package main

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	url       *url.URL
	urlParams url.Values
	formData  struct {
		query []byte
	}

	pageNumber int
}

func NewAPI_Caller(query []byte) (apc *API_Caller, err error) {
	apc = &API_Caller{}

	apc.urlParams = url.Values{
		"apiKey":   []string{"SEDIA"},
		"text":     []string{"***"},
		"pageSize": []string{"100"},
	}

	apc.url, err = url.Parse(API_ENDPOINT)
	if err != nil {
		return nil, err
	}

	apc.formData.query = query
	apc.pageNumber = 1

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

	if len(jsonResponse.Results) == 0 {
		close(resultsChan)
		return
	}

	fmt.Println("here")
	resultsChan <- jsonResponse.Results
	apc.pageNumber++
	apc.GetResults(resultsChan)
}

func (apc *API_Caller) createRequest() *http.Request {

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	var header textproto.MIMEHeader = make(textproto.MIMEHeader)
	header.Add("Content-Disposition", `form-data; name="query";`)
	header.Add("Content-Type", "application/json")

	part, err := writer.CreatePart(header)
	if err != nil {
		log.Fatalln(err)
	}

	part.Write(apc.formData.query)

	err = writer.Close()
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
