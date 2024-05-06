package api_caller

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestGetResults(t *testing.T) {

	resultsChan := make(chan []Result)
	bodyParams := make(map[string][]byte)
	urlParams := make(map[string][]string)

	file, err := os.Open(`E:\programi\GrantGPT\Fetcher\query.json`)
	if err != nil {
		fmt.Println(err)
		return
	}

	query, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	bodyParams["query"] = query
	bodyParams["language"] = []byte(`["en"]`)

	urlParams["apiKey"] = []string{"SEDIA"}
	urlParams["text"] = []string{"***"}
	urlParams["pageSize"] = []string{"100"}

	apicaller, err := NewAPI_Caller(bodyParams, urlParams)
	if err != nil {
		fmt.Println(err)
		return
	}

	go apicaller.GetResults(resultsChan)

	for resultArr := range resultsChan {
		fmt.Printf("results in this array: %v\n\n\n----------------\n\n\n-----------", len(resultArr))
		for _, result := range resultArr {
			fmt.Println(result.Metadata["identifier"])
		}
	}
}
