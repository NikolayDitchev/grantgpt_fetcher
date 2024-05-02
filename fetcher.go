package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
)

type Fetcher struct {
	query      []byte
	folderPath string

	resultsChan chan []Result
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

	fetcher.resultsChan = make(chan []Result)

	return
}

func (f *Fetcher) FetchData() {

	var wgGrants sync.WaitGroup

	apiCaller, err := NewAPI_Caller(f.query)
	if err != nil {
		log.Fatalln(err)
	}

	go apiCaller.GetResults(f.resultsChan)

	for results := range f.resultsChan {

		fmt.Println(len(results))

		for inx := range results {

			wgGrants.Add(1)

			go f.handleGrant(&results[inx], &wgGrants)
			fmt.Println(results[inx].Metadata["identifier"][0])
		}
	}

	wgGrants.Wait()
}

func (f *Fetcher) handleGrant(grant *Result, wgGrant *sync.WaitGroup) {
	defer wgGrant.Done()

	if len(grant.Metadata["callIdentifier"]) == 0 {
		log.Printf("no callIdentifier on %v\n", grant.Metadata["identifier"][0])
		return
	}

	if len(grant.Metadata["publicationDocuments"]) == 0 {
		log.Printf("no publicationDocuments on %v\n", grant.Metadata["identifier"][0])
		return
	}

	grantFolderPath := filepath.Join(f.folderPath, grant.Metadata["callIdentifier"][0])

	err := os.MkdirAll(grantFolderPath, 0777)
	if err != nil {
		log.Fatalln(err, grant.Metadata["identifier"][0])
	}

	var documents []Document

	err = json.Unmarshal([]byte(grant.Metadata["publicationDocuments"][0]), &documents)
	if err != nil {
		log.Fatalln(err, grant.Metadata["identifier"][0])
	}

	var wgDocs sync.WaitGroup
	for inx := range documents {
		wgDocs.Add(1)

		go func(doc *Document, grantFolderPath string) {
			defer wgDocs.Done()

			if doc.LanguageDoc != "EN" || (doc.TypeDoc != "pdf" && doc.TypeDoc != "docx") {
				return
			}

			err := doc.DownloadFile(grantFolderPath)
			if err != nil {
				log.Fatalln(err, filepath.Base(grantFolderPath), doc.NameDoc)
			}

		}(&documents[inx], grantFolderPath)
	}

	wgDocs.Wait()
}
