package danswer_api

import "testing"

func TestGetConnectors(t *testing.T) {
	apiClient := NewDanswerAPIClient()

	connectors, err := apiClient.GetConnectors()

	if err != nil {
		t.Errorf("error in getting connectors : %v", err)
	}

	_ = connectors
}
