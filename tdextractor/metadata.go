package main

import (
	"encoding/json"
	"strings"
)

const (
	eu_topic_url = `https://ec.europa.eu/info/funding-tenders/opportunities/portal/screen/opportunities/topic-details/`
)

type DanswerMetadata struct {
	Link             string `json:"link"`
	File_dislay_name string `json:"file_display_name"`
	Status           string `json:"status"`
}

func getMetadataJson(topicDetails *TopicDetails) (string, error) {
	metadata := DanswerMetadata{
		Link:             eu_topic_url + strings.ToLower(topicDetails.Identifier),
		File_dislay_name: topicDetails.Title,
		Status:           "OK",
	}

	metadataJson, err := json.Marshal(metadata)
	if err != nil {
		return "", err
	}

	metadataString := `<!-- DANSWER_METADATA=` + string(metadataJson) + ` -->`

	return metadataString, nil
}
