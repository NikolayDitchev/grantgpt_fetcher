package main

import (
	"fmt"
	"testing"
)

func TestFetchData(t *testing.T) {
	fetcher, err := NewFetcher(``, ``)
	if err != nil {
		fmt.Println(err)
	}

	fetcher.FetchData()
	//fmt.Printf("%+v\n", fetcher)
}
