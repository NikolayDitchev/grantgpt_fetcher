package tdextractor

import "testing"

func TestFetchData(t *testing.T) {

	fetcher := NewTopicDetailsFetcher(`E:\programi\GrantGPT\Fetcher\tdextractor\topics`)

	fetcher.FetchData()
}
