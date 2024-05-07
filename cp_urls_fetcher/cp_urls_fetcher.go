package cp_urls_fetcher

import (
	"fetcher/api_caller"
	"log"
	"net/url"
	"os"
)

type URLFetcher struct {
	file *os.File

	apiCaller   *api_caller.API_Caller
	resultsChan chan []api_caller.Result

	urlBuilders map[string]UrlBuilingFunc
}

func NewFetcher(query []byte, filePath string) (fetcher *URLFetcher, err error) {

	file, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}

	bodyParams := map[string][]byte{
		"query":     query,
		"languages": []byte(`["en"]`),
	}

	urlParams := url.Values{
		"apiKey":   []string{"SEDIA"},
		"pageSize": []string{"50"},
		"text":     []string{"***"},
	}

	resultsChan := make(chan []api_caller.Result)

	apiCaller, err := api_caller.NewAPI_Caller(bodyParams, urlParams)
	if err != nil {
		return nil, err
	}

	urlBuilders := getUrlBuilders()

	fetcher = &URLFetcher{
		file:        file,
		resultsChan: resultsChan,
		apiCaller:   apiCaller,
		urlBuilders: urlBuilders,
	}

	return
}

func (uf *URLFetcher) FetchData() {
	go uf.apiCaller.GetResults(uf.resultsChan)

	for results := range uf.resultsChan {
		for inx := range results {

			if len(results[inx].Metadata["type"]) == 0 {
				log.Println("no type on CoP" + results[inx].Metadata["identifier"][0])
				continue
			}

			cpType := results[inx].Metadata["type"][0]
			cpURL, err := uf.urlBuilders[cpType](&results[inx])

			if err != nil {
				log.Println(err)
			}

			uf.file.WriteString(cpURL)
			uf.file.Write([]byte{'\n'})

			//fmt.Println(inx)
		}
	}
}
