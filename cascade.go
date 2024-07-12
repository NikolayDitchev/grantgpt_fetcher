package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/NikolayDitchev/grantgpt_fetcher/eu_client"
)

var (
	cascadeTextFields = []string{
		"beneficiaryAdministration",
		"description",
		"duration",
		"furtherInformation",
		"caName",
		"deadlineDate",
		"startDate",
	}
)

func extractCascadeDetails(cascadeResult *eu_client.Result) (*fundingBuffer, error) {

	cascadeId, err := eu_client.GetMetadataField(cascadeResult, eu_client.CascadeIdField)
	if err != nil {
		fmt.Println("5")
		return nil, err
	}

	cascadeBuffer := &fundingBuffer{
		content:  &bytes.Buffer{},
		fileName: cascadeId,
	}

	danswerMetadata := map[string]string{
		"link":             eu_client.EUWebsiteCascadeURL + "/" + cascadeId,
		"file_dislay_name": cascadeResult.Summary,
		"status":           "OK",
	}

	metadataJson, err := json.Marshal(danswerMetadata)
	if err != nil {
		fmt.Println("6")
		return nil, err
	}

	err = cascadeBuffer.WriteString(`<!-- DANSWER_METADATA=` + string(metadataJson) + ` -->` + "\n\n")
	if err != nil {
		fmt.Println("7")
		return nil, err
	}

	for _, textField := range cascadeTextFields {

		fieldValue, _ := eu_client.GetMetadataField(cascadeResult, textField)
		if fieldValue == "" {
			continue
		}

		fieldValue = StripHTMLTags(fieldValue)
		_ = cascadeBuffer.WriteString(fmt.Sprintf("%s: %s\n\n", textField, fieldValue))

	}

	return cascadeBuffer, nil
}
