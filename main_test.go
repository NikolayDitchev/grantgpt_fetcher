package main

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/NikolayDitchev/grantgpt_fetcher/eu_client"
)

func TestGetTopicDetails(t *testing.T) {

	pages, err := getPages(eu_client.TypeTopics)
	if err != nil {
		t.Error(err)
	}

	for _, page := range pages {
		for inx := range page.Results {

			topicId, err := eu_client.GetMetadataField(&page.Results[inx], eu_client.TopicIdField)
			if err != nil {
				t.Error(err)
			}

			topicId = strings.ToLower(topicId)

			_, err = GetTopicDetails(topicId)
			if err != nil {
				fmt.Printf("index: %v\n, topicId: %v", inx, topicId)
				t.Error(err)
			}
		}
	}
}

func TestGetPages(t *testing.T) {

	topicIdsMap := map[string]int{}
	totalResults := 0

	pages, err := getPages(eu_client.TypeTopics)
	if err != nil {
		t.Error(err)
	}

	for i := 0; i < 50; i++ {

		for _, page := range pages {

			file, err := os.Create(fmt.Sprintf("page_%v_%v.txt", i, page.PageNumber))
			if err != nil {
				t.Error(err)
			}
			defer file.Close()

			totalResults = page.TotalResults

			for inx := range page.Results {

				topicId, err := eu_client.GetMetadataField(&page.Results[inx], eu_client.TopicIdField)
				if err != nil {
					t.Error(err)
				}

				if _, ok := topicIdsMap[topicId]; ok {
					continue
				}

				file.WriteString(topicId + "\n")

				topicIdsMap[topicId] = 1
			}
		}

		fmt.Println(len(topicIdsMap))

		if len(topicIdsMap) == totalResults {
			break
		}
	}
}
