package main

type tdResponse struct {
	TopicDetails TopicDetails `json:"TopicDetails"`
}

type TopicDetails struct {
	Type                int    `json:"type"`
	Ccm2ID              int    `json:"ccm2Id"`
	CftID               int    `json:"cftId"`
	Identifier          string `json:"identifier"`
	Title               string `json:"title"`
	PublicationDateLong int64  `json:"publicationDateLong"`
	CallIdentifier      string `json:"callIdentifier"`
	CallTitle           string `json:"callTitle"`
	Callccm2ID          int    `json:"callccm2Id"`
	AllowPartnerSearch  bool   `json:"allowPartnerSearch"`
	FrameworkProgramme  struct {
		ID           int    `json:"id"`
		Abbreviation string `json:"abbreviation"`
		Description  string `json:"description"`
	} `json:"frameworkProgramme"`
	ProgrammeDivision []struct {
		ID           int    `json:"id"`
		Abbreviation string `json:"abbreviation"`
		Description  string `json:"description"`
	} `json:"programmeDivision"`
	DestinationDetails     string   `json:"destinationDetails"`
	DestinationDescription string   `json:"destinationDescription"`
	MissionDetails         string   `json:"missionDetails"`
	TopicMGAs              []any    `json:"topicMGAs"`
	Tags                   []string `json:"tags"`
	Keywords               []string `json:"keywords"`
	Flags                  []string `json:"flags"`
	Sme                    bool     `json:"sme"`
	Actions                []struct {
		Status struct {
			ID           int    `json:"id"`
			Abbreviation string `json:"abbreviation"`
			Description  string `json:"description"`
		} `json:"status"`
		Types []struct {
			TypeOfAction string `json:"typeOfAction"`
			TypeOfMGA    []struct {
				ID           int    `json:"id"`
				Abbreviation string `json:"abbreviation"`
				Description  string `json:"description"`
			} `json:"typeOfMGA"`
		} `json:"types"`
		PlannedOpeningDate  string `json:"plannedOpeningDate"`
		SubmissionProcedure struct {
			ID           int    `json:"id"`
			Abbreviation string `json:"abbreviation"`
			Description  string `json:"description"`
		} `json:"submissionProcedure"`
		DeadlineDates []string `json:"deadlineDates"`
	} `json:"actions"`
	LatestInfos            []any `json:"latestInfos"`
	BudgetOverviewJSONItem struct {
		BudgetTopicActionMap map[string][]Action `json:"budgetTopicActionMap"`
		BudgetYearsColumns   []string            `json:"budgetYearsColumns"`
	} `json:"budgetOverviewJSONItem"`
	Description         string `json:"description"`
	Conditions          string `json:"conditions"`
	SupportInfo         string `json:"supportInfo"`
	SepTemplate         string `json:"sepTemplate"`
	Links               []any  `json:"links"`
	AdditionalDossiers  []any  `json:"additionalDossiers"`
	InfoPackDossiers    []any  `json:"infoPackDossiers"`
	CallDetailsJSONItem struct {
		LatestInfos          []any `json:"latestInfos"`
		HasForthcomingTopics bool  `json:"hasForthcomingTopics"`
		HasOpenTopics        bool  `json:"hasOpenTopics"`
		AllClosedTopics      bool  `json:"allClosedTopics"`
	} `json:"callDetailsJSONItem"`
}

type Action struct {
	Action               string             `json:"action"`
	PlannedOpeningDate   string             `json:"plannedOpeningDate"`
	DeadlineModel        string             `json:"deadlineModel"`
	DeadlineDates        []string           `json:"deadlineDates"`
	BudgetYearMap        map[string]float64 `json:"budgetYearMap"`
	ExpectedGrants       float64            `json:"expectedGrants"`
	MinContribution      float64            `json:"minContribution"`
	MaxContribution      float64            `json:"maxContribution"`
	BudgetTopicActionMap struct {
	} `json:"budgetTopicActionMap"`
}
