package main

import (
	"io"
	"os"
)

const (
	METHOD       = "POST"
	API_ENDPOINT = "https://api.tech.ec.europa.eu/search-api/prod/rest/search"
)

type Fetcher struct {
	query      []byte
	folderPath string
}

func NewFetcher(queryFilePath string, downloadFolderPath string) (fetcher *Fetcher, err error) {
	queryFile, err := os.Open(queryFilePath)
	if err != nil {
		return nil, err
	}

	query, err := io.ReadAll(queryFile)
	if err != nil {
		return nil, err
	}

	fetcher = &Fetcher{
		query:      query,
		folderPath: downloadFolderPath,
	}

	//fmt.Println(string(fetcher.query))

	return
}

func (f *Fetcher) FetchData() {

}
