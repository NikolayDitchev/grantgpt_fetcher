package grant_doc_fetcher

import (
	"io"
	"os"
	"testing"
)

func TestFetchData(t *testing.T) {
	queryFilePath := `E:\programi\GrantGPT\Fetcher\query.json`
	downloadFolderPath := `E:\programi\GrantGPT\Fetcher\grants`

	file, err := os.Open(queryFilePath)
	if err != nil {
		t.Errorf("%v", err)
	}

	query, err := io.ReadAll(file)
	if err != nil {
		t.Errorf("%v", err)
	}

	grantDocFetcher, err := NewFetcher(query, downloadFolderPath)

	if err != nil {
		t.Errorf("%v", err)
	}

	grantDocFetcher.FetchData()
}
