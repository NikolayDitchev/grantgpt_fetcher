package tdextractor

import "testing"

func TestFetchData(t *testing.T) {

	fetcher := NewTopicDetailsFetcher(`E:\programi\GrantGPT\Fetcher\tdextractor\topics`)

	fetcher.FetchData()
}

func TestCreateZip(t *testing.T) {
	fetcher := NewTopicDetailsFetcher(`topicZip`)

	err := fetcher.CreateZip()
	if err != nil {
		t.Error(err.Error())
	}
}
