package eu_client

const (
	TypeTopics  = "1"
	TypeGrant   = "2"
	TypeCascade = "8"

	StatusOpen        = "31094501"
	StatusForthcoming = "31094502"
)

const (
	EndpointSearch       = "https://api.tech.ec.europa.eu/search-api/prod/rest/search"
	EndpointTopicDetails = "https://ec.europa.eu/info/funding-tenders/opportunities/data/topicDetails" // /{topicId}

	//Grant Details endpoint
	EndpointDocuments = "https://api.tech.ec.europa.eu/search-api/prod/rest/document" // /{documentID}
	GrantSufix        = "PROSPECTSEN"
)

const (
	EUWebsiteCascadeURL = "https://ec.europa.eu/info/funding-tenders/opportunities/portal/screen/opportunities/competitive-calls-cs" // /{callccm2Id}
	EUWebsiteTopicURL   = "https://ec.europa.eu/info/funding-tenders/opportunities/portal/screen/opportunities/topic-details"        // /{topicId}
	EUWebsiteGrantURL   = "https://ec.europa.eu/info/funding-tenders/opportunities/portal/screen/opportunities/prospect-details"     // /{callIdentifier}PROSPECTSEN
)

const (
	ApiKeySedia = "SEDIA"
)

const (
	DefaultTest = "***"
)

const (
	maxPageNumber = 2000
)

const (
	CascadeIdField = "callccm2Id"
	TopicIdField   = "identifier"
	GrantIdField   = "callIdentifier"
)
