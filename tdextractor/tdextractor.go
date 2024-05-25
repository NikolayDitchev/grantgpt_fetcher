package tdextractor

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fetcher/api_caller"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
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

}

func (tdf *TopicDetailsFetcher) CreateZip() error {
	topicsZip, err := os.Create(tdf.folderPath + ".zip")
	if err != nil {
		return err
	}

	zipWriter := zip.NewWriter(topicsZip)
	topicBuffersChan := make(chan *topicBuffer)

	go func() {
		eu_client := api_caller.NewAPI_Caller()
		topicIDsChan := make(chan string)
		errChan := make(chan error)

		go eu_client.GetTopicIDs(topicIDsChan, errChan)

		go func() {
			for err := range errChan {
				log.Println(err)
			}
		}()

		for topicID := range topicIDsChan {

			go func(id string) {

				topicBuffer, err := tdf.ExtractTopicDetails(id)
				if err != nil {
					log.Fatal(err)
				}

				topicBuffersChan <- topicBuffer

			}(topicID)
		}

		close(topicBuffersChan)

	}()

	counter := 0

	for topicBuffer := range topicBuffersChan {

		zipFileWriter, err := zipWriter.Create(topicBuffer.GetTopicId() + ".txt")
		if err != nil {
			return err
		}

		_, err = io.Copy(zipFileWriter, topicBuffer.GetContent())
		if err != nil {
			return err
		}

		counter++
	}

	err = zipWriter.Close()
	if err != nil {
		return err
	}

	fmt.Println(counter)

	return nil
}

func (tdf *TopicDetailsFetcher) ExtractTopicDetails(topicID string) (*topicBuffer, error) {
	topicID = strings.ToLower(topicID)

	topicBuffer := &topicBuffer{
		content: &bytes.Buffer{},
		id:      topicID,
	}

	var url string = base_url + topicID + json_suffix

	resp, err := tdf.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var topicJson tdResponse

	err = json.Unmarshal(body, &topicJson)
	if err != nil {
		return nil, err
	}

	metadata, err := getMetadataJson(&topicJson.TopicDetails)
	if err != nil {
		return nil, err
	}

	err = topicBuffer.WriteString(metadata + "\n\n")
	if err != nil {
		return nil, err
	}

	regex, err := regexp.Compile("<.+?>")
	if err != nil {
		return nil, err
	}

	description := "Description: \n" + regex.ReplaceAllString(topicJson.TopicDetails.Description, "") + "\n\n"
	err = topicBuffer.WriteString(description)
	if err != nil {
		return nil, err
	}

	conditions := "Conditions: \n" + regex.ReplaceAllString(topicJson.TopicDetails.Conditions, "") + "\n\n"
	err = topicBuffer.WriteString(conditions)
	if err != nil {
		return nil, err
	}

	supportInfo := "Support Info: \n" + regex.ReplaceAllString(topicJson.TopicDetails.SupportInfo, "") + "\n\n"
	err = topicBuffer.WriteString(supportInfo)
	if err != nil {
		return nil, err
	}

	return topicBuffer, nil
}

type topicBuffer struct {
	content *bytes.Buffer
	id      string
}

func (tb *topicBuffer) GetTopicId() string {
	return tb.id
}

func (tb *topicBuffer) SetTopicId(topicId string) {
	tb.id = topicId
}

func (tb *topicBuffer) GetContent() *bytes.Buffer {
	return tb.content
}

func (tb *topicBuffer) SetContent(content *bytes.Buffer) {
	tb.content = content
}

func (tb *topicBuffer) WriteString(data string) error {
	_, err := tb.content.WriteString(data)
	return err
}
