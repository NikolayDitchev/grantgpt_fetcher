package tdextractor

import (
	"encoding/json"
	"fetcher/api_caller"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"
)

const (
	base_url    = `https://ec.europa.eu/info/funding-tenders/opportunities/data/topicDetails/`
	json_suffix = `.json`
)

type TopicDetailsFetcher struct {
	folderPath string
	client     *http.Client
}

func NewTopicDetailsFetcher(folderPath string) *TopicDetailsFetcher {

	fetcher := &TopicDetailsFetcher{
		folderPath: folderPath,
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}

	return fetcher
}

func (tdf *TopicDetailsFetcher) FetchData() {

	err := os.MkdirAll(tdf.folderPath, 0777)
	if err != nil {
		log.Fatalln(err)
	}

	apc, _ := api_caller.NewAPI_Caller()
	var wgTopics sync.WaitGroup

	topicIDs := apc.GetTopicIDs()
	if err != nil {
		log.Fatalln(err)
	}

	for topicID := range topicIDs {

		wgTopics.Add(1)

		go func(topicID string, wg *sync.WaitGroup) {
			defer wg.Done()

			err := tdf.ExtractTopicDetails(topicID)
			if err != nil {
				log.Fatalln(err)
			}

		}(topicID, &wgTopics)

	}

	wgTopics.Wait()

}

func (tdf *TopicDetailsFetcher) ExtractTopicDetails(topicID string) error {
	topicID = strings.ToLower(topicID)

	filePath := filepath.Join(tdf.folderPath, topicID+".txt")

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	var url string = base_url + topicID + json_suffix

	resp, err := tdf.client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var topicJson tdResponse

	err = json.Unmarshal(body, &topicJson)
	if err != nil {
		return err
	}

	metadata, err := getMetadataJson(&topicJson.TopicDetails)
	if err != nil {
		return err
	}

	_, err = file.WriteString(metadata + "\n\n")
	if err != nil {
		return err
	}

	regex, err := regexp.Compile("<.+?>")
	if err != nil {
		log.Fatalln(err)
	}

	description := "Description: \n" + regex.ReplaceAllString(topicJson.TopicDetails.Description, "") + "\n\n"
	_, err = file.WriteString(description)
	if err != nil {
		log.Fatalln(err)
	}

	conditions := "Conditions: \n" + regex.ReplaceAllString(topicJson.TopicDetails.Conditions, "") + "\n\n"
	_, err = file.WriteString(conditions)
	if err != nil {
		log.Fatalln(err)
	}

	supportInfo := "Support Info: \n" + regex.ReplaceAllString(topicJson.TopicDetails.SupportInfo, "") + "\n\n"
	_, err = file.WriteString(supportInfo)
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}
