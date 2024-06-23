package eu_client

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"
)

func TestGetQuery(t *testing.T) {

	query := NewQuery(
		WithTypes([]string{TypeTopics}),
		WithStatus([]string{StatusForthcoming, StatusOpen}),
	)

	queryJson, err := json.Marshal(query)
	if err != nil {
		t.Error(err)
	}

	var prettyQuery bytes.Buffer
	err = json.Indent(&prettyQuery, queryJson, " ", "\t")
	if err != nil {
		t.Error(err)
	}

	prettyQuery.WriteTo(os.Stdout)
}
