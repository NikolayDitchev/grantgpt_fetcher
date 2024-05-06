package main

import (
	"fetcher/grant_doc_fetcher"
)

type Fetcher interface {
	FetchData()
}

func FetcherFactory(fetcherType string, data map[string]string) (fetcher Fetcher, err error) {

	switch fetcherType {
	case "grantDocFetcher":
		fetcher, err = grant_doc_fetcher.NewFetcher(data["queryFilePath"], data["downloadFolderPath"])
		return
	default:
		return
	}

}
