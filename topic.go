package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/NikolayDitchev/grantgpt_fetcher/eu_client"
)

var (
	topicTextFields = []string{
		"descriptionByte",
		"topicConditions",
		"supportInfo",
		"additionalInfo",
		"destinationDetails",
		"missionDetails",
		"destinationDescription",
		"missionDescription",
	}
)

func extractTopicDetails(topicResult *eu_client.Result) (*fundingBuffer, error) {

	topicId, err := eu_client.GetMetadataField(topicResult, eu_client.TopicIdField)
	if err != nil {
		fmt.Println("1")
		return nil, err
	}

	topicId = strings.ToLower(topicId)

	topicBuffer := &fundingBuffer{
		content:  &bytes.Buffer{},
		fileName: topicId,
	}

	// topicDetails, err := GetTopicDetails(topicId)
	// if err != nil {
	// 	fmt.Println(topicId)
	// 	fmt.Println("2")
	// 	return nil, err
	// }

	danswerMetadata := map[string]string{
		"link":              eu_client.EUWebsiteTopicURL + "/" + topicId,
		"file_display_name": topicResult.Metadata["title"][0],
		"status":            "OK",
	}

	metadataJson, err := json.Marshal(danswerMetadata)
	if err != nil {
		fmt.Println("3")
		return nil, err
	}

	err = topicBuffer.WriteString(`<!-- DANSWER_METADATA=` + string(metadataJson) + ` -->` + "\n\n")
	if err != nil {
		fmt.Println("4")
		return nil, err
	}

	//topicReflectValue := reflect.ValueOf(*topicDetails)

	for _, textField := range topicTextFields {

		fieldValue, ok := topicResult.Metadata[textField]
		if !ok || len(fieldValue) == 0 {
			continue
		}

		fieldValueString := StripHTMLTags(fieldValue[0])
		_ = topicBuffer.WriteString(fmt.Sprintf("%s: %s\n\n", textField, fieldValueString))

	}

	startDate, ok := topicResult.Metadata["startDate"]
	if ok {
		_ = topicBuffer.WriteString(fmt.Sprintf("Start Date: %s\n\n", startDate[0]))
	}

	endDate, ok := topicResult.Metadata["deadlineDate"]
	if ok {
		_ = topicBuffer.WriteString(fmt.Sprintf("End Date: %s\n\n", endDate[0]))
	}

	return topicBuffer, nil
}

// func GetTopicDetails(topicId string) (*TopicDetails, error) {

// 	client := &http.Client{
// 		Timeout: 5 * time.Second,
// 	}

// 	var url string = eu_client.EndpointTopicDetails + "/" + topicId + ".json"

// 	body := make([]byte, 1024)
// 	var err error

// 	for i := 0; i < 10; i++ {

// 		resp, err := client.Get(url)
// 		if err != nil {
// 			continue
// 		}

// 		body, err = io.ReadAll(resp.Body)
// 		resp.Body.Close()

// 		if err == nil {
// 			break
// 		}
// 	}

// 	if err != nil {
// 		return nil, err
// 	}

// 	var topicJson *tdResponse

// 	err = json.Unmarshal(body, &topicJson)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &topicJson.TopicDetails, nil
// }

// type tdResponse struct {
// 	TopicDetails TopicDetails `json:"TopicDetails"`
// }

// type TopicDetails struct {
// 	Type                int    `json:"type"`
// 	Ccm2ID              int    `json:"ccm2Id"`
// 	CftID               int    `json:"cftId"`
// 	Identifier          string `json:"identifier"`
// 	Title               string `json:"title"`
// 	PublicationDateLong int64  `json:"publicationDateLong"`
// 	CallIdentifier      string `json:"callIdentifier"`
// 	CallTitle           string `json:"callTitle"`
// 	Callccm2ID          int    `json:"callccm2Id"`
// 	AllowPartnerSearch  bool   `json:"allowPartnerSearch"`
// 	FrameworkProgramme  struct {
// 		ID           int    `json:"id"`
// 		Abbreviation string `json:"abbreviation"`
// 		Description  string `json:"description"`
// 	} `json:"frameworkProgramme"`
// 	ProgrammeDivision []struct {
// 		ID           int    `json:"id"`
// 		Abbreviation string `json:"abbreviation"`
// 		Description  string `json:"description"`
// 	} `json:"programmeDivision"`
// 	DestinationDetails     string   `json:"destinationDetails"`
// 	DestinationDescription string   `json:"destinationDescription"`
// 	MissionDetails         string   `json:"missionDetails"`
// 	MissionDescription     string   `json:"missionDescription"`
// 	TopicMGAs              []any    `json:"topicMGAs"`
// 	Tags                   []string `json:"tags"`
// 	Keywords               []string `json:"keywords"`
// 	Flags                  []string `json:"flags"`
// 	Sme                    bool     `json:"sme"`
// 	Actions                []struct {
// 		Status struct {
// 			ID           int    `json:"id"`
// 			Abbreviation string `json:"abbreviation"`
// 			Description  string `json:"description"`
// 		} `json:"status"`
// 		Types []struct {
// 			TypeOfAction string `json:"typeOfAction"`
// 			TypeOfMGA    []struct {
// 				ID           int    `json:"id"`
// 				Abbreviation string `json:"abbreviation"`
// 				Description  string `json:"description"`
// 			} `json:"typeOfMGA"`
// 		} `json:"types"`
// 		PlannedOpeningDate  string `json:"plannedOpeningDate"`
// 		SubmissionProcedure struct {
// 			ID           int    `json:"id"`
// 			Abbreviation string `json:"abbreviation"`
// 			Description  string `json:"description"`
// 		} `json:"submissionProcedure"`
// 		DeadlineDates []string `json:"deadlineDates"`
// 	} `json:"actions"`
// 	LatestInfos            []any `json:"latestInfos"`
// 	BudgetOverviewJSONItem struct {
// 		BudgetTopicActionMap map[string][]Action `json:"budgetTopicActionMap"`
// 		BudgetYearsColumns   []string            `json:"budgetYearsColumns"`
// 	} `json:"budgetOverviewJSONItem"`
// 	Description         string              `json:"description"`
// 	Conditions          string              `json:"conditions"`
// 	SupportInfo         string              `json:"supportInfo"`
// 	SepTemplate         string              `json:"sepTemplate"`
// 	Links               []map[string]string `json:"links"`
// 	AdditionalDossiers  []any               `json:"additionalDossiers"`
// 	InfoPackDossiers    []any               `json:"infoPackDossiers"`
// 	CallDetailsJSONItem struct {
// 		LatestInfos          []any `json:"latestInfos"`
// 		HasForthcomingTopics bool  `json:"hasForthcomingTopics"`
// 		HasOpenTopics        bool  `json:"hasOpenTopics"`
// 		AllClosedTopics      bool  `json:"allClosedTopics"`
// 	} `json:"callDetailsJSONItem"`
// }

// type Action struct {
// 	Action               string             `json:"action"`
// 	PlannedOpeningDate   string             `json:"plannedOpeningDate"`
// 	DeadlineModel        string             `json:"deadlineModel"`
// 	DeadlineDates        []string           `json:"deadlineDates"`
// 	BudgetYearMap        map[string]float64 `json:"budgetYearMap"`
// 	ExpectedGrants       float64            `json:"expectedGrants"`
// 	MinContribution      float64            `json:"minContribution"`
// 	MaxContribution      float64            `json:"maxContribution"`
// 	BudgetTopicActionMap struct {
// 	} `json:"budgetTopicActionMap"`
// }
