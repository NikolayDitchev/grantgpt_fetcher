package grant_doc_fetcher

import "testing"

func TestFetchData(t *testing.T) {
	queryFilePath := `E:\programi\GrantGPT\Fetcher\query.json`
	downloadFolderPath := `E:\programi\GrantGPT\Fetcher\grants`

	grantDocFetcher, err := NewFetcher(queryFilePath, downloadFolderPath)

	if err != nil {
		t.Errorf(err.Error())
	}

	grantDocFetcher.FetchData()
}
