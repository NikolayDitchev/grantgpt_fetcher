package main

import (
	"fetcher/cp_urls_fetcher"
	"fetcher/grant_doc_fetcher"
	"io"
	"os"
)

type Fetcher interface {
	FetchData()
}

func FetcherFactory(fetcherType string, data map[string]string) (fetcher Fetcher, err error) {

	file, err := os.Open(data["queryFilePath"])
	if err != nil {
		return nil, err
	}

	query, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	switch fetcherType {
	case "grantDocFetcher":
		fetcher, err = grant_doc_fetcher.NewFetcher(query, data["downloadFolderPath"])
		return
	case "urlFetcher":
		fetcher, err = cp_urls_fetcher.NewFetcher(query, data["downloadFolderPath"])
	default:
		return
	}

	return
}
