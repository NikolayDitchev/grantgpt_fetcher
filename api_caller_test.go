package main

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestGetResults(t *testing.T) {

	resultsChan := make(chan []Result)

	file, err := os.Open(``)
	if err != nil {
		fmt.Println(err)
		return
	}

	query, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	apicaller, err := NewAPI_Caller(query)
	if err != nil {
		fmt.Println(err)
		return
	}

	go apicaller.GetResults(resultsChan)

	for resultArr := range resultsChan {
		fmt.Printf("results in this array: %v", len(resultArr))
		for _, result := range resultArr {
			fmt.Println(result.Metadata["identifier"])
		}
	}
}
