package eu_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"testing"
	"time"
)

// func TestGetTopicIDs(t *testing.T) {

// 	topicIds := make(chan string)
// 	errChan := make(chan error)

// 	apc := NewAPI_Caller(10 * time.Second)

// 	go apc.GetTopicIDs(topicIds, errChan)

// 	counter := 0

// 	go func() {
// 		for err := range errChan {
// 			t.Error(err)
// 		}
// 	}()

// 	for topicId := range topicIds {
// 		fmt.Println(topicId)
// 		counter++
// 	}

// 	fmt.Println(counter)
// }

func TestGetQuery(t *testing.T) {

	query := NewQuery(
		WithTypes(TypeTopics),
		WithStatus(StatusForthcoming, StatusOpen),
	)

	queryJson, err := json.Marshal(query)
	if err != nil {
		t.Error(err)
	}

	var prettyQuery bytes.Buffer
	err = json.Indent(&prettyQuery, queryJson, " ", "\t")
	if err != nil {
		t.Error(err)
	}

	prettyQuery.WriteTo(os.Stdout)
}

func TestSendRequest(t *testing.T) {

	query := NewQuery(WithTypes(TypeTopics), WithStatus(StatusOpen, StatusForthcoming))

	req, err := NewEURequest(
		"POST",
		EndpointSearch,
		[]RequestBodyOption{
			WithQuery(query),
			WithLanguages("en"),
		},
		[]RequestURLOption{
			WithApiKey(ApiKeySedia),
			WithPageSize(50),
			WithText(DefaultTest),
			WithPageNumber(1),
		},
	)

	if err != nil {
		t.Error(err)
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	var prettyQuery bytes.Buffer
	err = json.Indent(&prettyQuery, body, " ", "\t")
	if err != nil {
		t.Error(err)
	}

	prettyQuery.WriteTo(os.Stdout)

}

func TestGetPages(t *testing.T) {
	query := NewQuery(WithTypes(TypeTopics), WithStatus(StatusOpen, StatusForthcoming))

	req, err := NewEURequest(
		"POST",
		EndpointSearch,
		[]RequestBodyOption{
			WithQuery(query),
			WithLanguages("en"),
		},
		[]RequestURLOption{
			WithApiKey(ApiKeySedia),
			WithPageSize(50),
			WithText(DefaultTest),
			WithPageNumber(1),
		},
	)

	if err != nil {
		t.Error(err)
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	pages, err := GetPages(req, client)
	if err != nil {
		t.Error(err)
	}

	for i, page := range pages {
		file, err := os.Create("pages" + strconv.Itoa(i) + ".txt")
		if err != nil {
			t.Error(err)
		}
		_, err = file.WriteString(fmt.Sprintf("%+v\n", page))
		if err != nil {
			t.Error(err)
		}
	}
}
