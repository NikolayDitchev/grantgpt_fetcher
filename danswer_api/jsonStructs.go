package danswer_api

import "time"

type Connector struct {
	Name                    string `json:"name"`
	Source                  string `json:"source"`
	InputType               string `json:"input_type"`
	ConnectorSpecificConfig struct {
		BaseURL          string   `json:"base_url"`
		WebConnectorType string   `json:"web_connector_type"`
		FileLocations    []string `json:"file_locations"`
		FolderPaths      []string `json:"folder_paths"`
		IncludeShared    bool     `json:"include_shared"`
		OnlyOrgPublic    bool     `json:"only_org_public"`
		FollowShortcuts  bool     `json:"follow_shortcuts"`
	} `json:"connector_specific_config,omitempty"`
	RefreshFreq   int       `json:"refresh_freq"`
	Disabled      bool      `json:"disabled"`
	ID            int       `json:"id"`
	CredentialIds []int     `json:"credential_ids"`
	TimeCreated   time.Time `json:"time_created"`
	TimeUpdated   time.Time `json:"time_updated"`
}
