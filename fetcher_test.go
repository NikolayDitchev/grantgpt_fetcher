package main

import (
	"fmt"
	"testing"
)

func TestNewFetcher(t *testing.T) {
	fetcher, err := NewFetcher(`E:\programi\GrantGPT\Fetcher\query.json`,
		`E:\programi\GrantGPT\Fetcher\grants`)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", fetcher)
}
