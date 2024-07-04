package main

import (
	"io"
	"os"
	"testing"
)

func TestFetchData(t *testing.T) {

	fetcher := NewTopicDetailsFetcher(`E:\programi\GrantGPT\Fetcher\tdextractor\topics`)

	fetcher.FetchData()
}

func TestCreateZip(t *testing.T) {
	fetcher := NewTopicDetailsFetcher(`topicZip`)

	err := fetcher.CreateZip("topicDetails")
	if err != nil {
		t.Error(err.Error())
	}
}

func TestExtractorDetails(t *testing.T) {

	fetcher := NewTopicDetailsFetcher("topic")

	topicId := "horizon-eic-2024-accelerator-01"

	topicBuffer, err := fetcher.ExtractTopicDetails(topicId)
	if err != nil {
		t.Errorf("error extracting details: %v", err)
	}

	file, err := os.Create(topicId + ".txt")
	if err != nil {
		t.Errorf("erro creating file: %v", err)
	}

	_, err = io.Copy(file, topicBuffer.GetContent())
	if err != nil {
		t.Errorf("error writing to file: %v", err)
	}

}
