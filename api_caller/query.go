package api_caller

import (
	"encoding/json"
)

const (
	type_topic   = "1"
	type_grant   = "2"
	type_cascade = "8"

	status_open        = "31094501"
	status_forthcoming = "31094502"
)

type Query struct {
	Bool struct {
		Must []struct {
			Terms map[string][]string `json:"terms,omitempty"`
		} `json:"must"`
	} `json:"bool"`
}

func GetTopicQuery() []byte {
	query := Query{}

	query.Bool.Must = append(query.Bool.Must, struct {
		Terms map[string][]string `json:"terms,omitempty"`
	}{
		Terms: map[string][]string{
			"type": {
				type_topic,
			},
		},
	})

	query.Bool.Must = append(query.Bool.Must, struct {
		Terms map[string][]string "json:\"terms,omitempty\""
	}{
		Terms: map[string][]string{
			"status": {
				status_forthcoming,
				status_open,
			},
		},
	})

	queryJson, _ := json.Marshal(query)

	return queryJson
}
