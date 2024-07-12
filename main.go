package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/NikolayDitchev/grantgpt_fetcher/eu_client"
)

type detailsExtractor func(*eu_client.Result) (*fundingBuffer, error)

var (
	extratorsMap = map[string]detailsExtractor{
		eu_client.TypeTopics:  extractTopicDetails,
		eu_client.TypeCascade: extractCascadeDetails,
	}
)

var wgFunding sync.WaitGroup

func main() {

	fundingBufferChan := make(chan *fundingBuffer)
	errChan := make(chan error)
	done := make(chan struct{})

	t := time.Now().UTC().Format(time.DateOnly)

	danswerZip, err := os.Create(fmt.Sprintf("danswer-fundings_%s.zip", t))
	if err != nil {
		return
	}
	defer danswerZip.Close()

	zipWriter := zip.NewWriter(danswerZip)
	defer zipWriter.Close()

	go FetchData(fundingBufferChan, errChan, done)

	// for fundingBuffer := range fundingBufferChan {

	// 	zipFileWrites, err := zipWriter.Create(fundingBuffer.fileName + ".txt")
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	_, err = io.Copy(zipFileWrites, fundingBuffer.content)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	for {

		select {
		case fundingBuffer := <-fundingBufferChan:

			if fundingBuffer == nil {
				continue
			}

			zipFileWrites, err := zipWriter.Create(fundingBuffer.fileName + ".txt")
			if err != nil {
				panic(err)
			}
			_, err = io.Copy(zipFileWrites, fundingBuffer.content)
			if err != nil {
				panic(err)
			}

			continue

		case err := <-errChan:
			fmt.Println("error:", err)
			continue
		case <-done:
			fmt.Println("done")
		}

		break
	}

}

func getPages(types ...string) ([]*eu_client.Page, error) {

	query := eu_client.NewQuery(
		eu_client.WithTypes(types...),
		eu_client.WithStatus(eu_client.StatusOpen, eu_client.StatusForthcoming),
	)

	req, err := eu_client.NewEURequest(
		"POST",
		eu_client.EndpointSearch,
		[]eu_client.RequestBodyOption{
			eu_client.WithQuery(query),
			eu_client.WithLanguages("en"),
		},
		[]eu_client.RequestURLOption{
			eu_client.WithApiKey(eu_client.ApiKeySedia),
			eu_client.WithPageSize(100),
			eu_client.WithText(""),
			eu_client.WithPageNumber(1),
		},
	)

	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Timeout: 30 * time.Minute,
	}

	return eu_client.GetPages(req, client)
}

func FetchData(fundingBufferChan chan *fundingBuffer, errChan chan error, done chan struct{}) {

	pages, err := getPages(eu_client.TypeTopics, eu_client.TypeCascade)
	if err != nil {
		panic(err)
	}

	for _, page := range pages {
		for inx := range page.Results {

			fundingType, err := eu_client.GetMetadataField(&page.Results[inx], "type")
			if err != nil {
				log.Println(err)
				continue
			}

			extractor, ok := extratorsMap[fundingType]
			if !ok {
				continue
			}

			wgFunding.Add(1)
			go func(wgFunding *sync.WaitGroup) {
				defer wgFunding.Done()

				fundingBuffer, err := extractor(&page.Results[inx])
				if err != nil {
					errChan <- err
				}

				fundingBufferChan <- fundingBuffer

			}(&wgFunding)
		}
	}

	wgFunding.Wait()

	done <- struct{}{}
	// close(done)
	// close(errChan)
	// close(fundingBufferChan)
}
