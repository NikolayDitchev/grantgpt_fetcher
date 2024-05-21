package cp_urls_fetcher

// import (
// 	"io"
// 	"os"
// 	"testing"
// )

// func TestUrlFetch(t *testing.T) {
// 	queryFilePath := `E:\programi\GrantGPT\Fetcher\query.json`
// 	downloadFolderPath := `E:\programi\GrantGPT\Fetcher\url.txt`

// 	file, err := os.Open(queryFilePath)
// 	if err != nil {
// 		t.Errorf("%v", err)
// 	}

// 	query, err := io.ReadAll(file)
// 	if err != nil {
// 		t.Errorf("%v", err)
// 	}

// 	urlFetcher, err := NewFetcher(query, downloadFolderPath)

// 	if err != nil {
// 		t.Errorf("%v", err)
// 	}

// 	urlFetcher.FetchData()

// }
