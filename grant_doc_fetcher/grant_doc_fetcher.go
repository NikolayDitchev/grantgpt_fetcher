package grant_doc_fetcher

import (
	"encoding/json"
	"fetcher/api_caller"
	"io"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"sync"
)

type GrantDocFetcher struct {
	query      []byte
	folderPath string

	resultsChan chan []api_caller.Result
	apiCaller   *api_caller.API_Caller
}

func NewFetcher(queryFilePath string, downloadFolderPath string) (fetcher *GrantDocFetcher, err error) {
	queryFile, err := os.Open(queryFilePath)
	if err != nil {
		return nil, err
	}

	query, err := io.ReadAll(queryFile)
	if err != nil {
		return nil, err
	}

	fetcher = &GrantDocFetcher{
		query:      query,
		folderPath: downloadFolderPath,
	}

	fetcher.resultsChan = make(chan []api_caller.Result)

	bodyParams := map[string][]byte{
		"query":    query,
		"language": []byte(`["en"]`),
	}

	urlParams := url.Values{
		"apiKey":   []string{"SEDIA"},
		"pageSize": []string{"50"},
		"text":     []string{"***"},
	}

	fetcher.apiCaller, err = api_caller.NewAPI_Caller(bodyParams, urlParams)
	if err != nil {
		return nil, err
	}

	return
}

func (f *GrantDocFetcher) FetchData() {

	var wgGrants sync.WaitGroup

	go f.apiCaller.GetResults(f.resultsChan)

	for results := range f.resultsChan {
		for inx := range results {

			wgGrants.Add(1)
			go f.handleGrant(&results[inx], &wgGrants)
		}
	}

	wgGrants.Wait()
}

func (f *GrantDocFetcher) handleGrant(grant *api_caller.Result, wgGrant *sync.WaitGroup) {
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
