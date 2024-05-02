package main

type ResponseBody struct {
	APIVersion    string //`json:"apiVersion"`
	Terms         string //`json:"terms"`
	ResponseTime  int    //`json:"responseTime"`
	TotalResults  int    //`json:"totalResults"`
	PageNumber    int    //`json:"pageNumber"`
	PageSize      int    //`json:"pageSize"`
	Sort          string //`json:"sort"`
	GroupByField  string //`json:"groupByField"`
	QueryLanguage struct {
		Language    string  //`json:"language"`
		Probability float64 //`json:"probability"`
	} //`json:"queryLanguage"`
	SpellingSuggestion any      //`json:"spellingSuggestion"`
	BestBets           []any    //`json:"bestBets"`
	Results            []Result //`json:"results"`
	Warnings           []any    //`json:"warnings"`
}

type Result struct {
	APIVersion        string  //`json:"apiVersion"`
	Reference         string  //`json:"reference"`
	URL               string  //`json:"url"`
	Title             any     //`json:"title"`
	ContentType       string  //`json:"contentType"`
	Language          string  //`json:"language"`
	DatabaseLabel     string  //`json:"databaseLabel"`
	Database          string  //`json:"database"`
	Summary           string  //`json:"summary"`
	Weight            float64 //`json:"weight"`
	GroupByID         string  //`json:"groupById"`
	Content           string  //`json:"content"`
	AccessRestriction bool    //`json:"accessRestriction"`
	Pages             any     //`json:"pages"`
	Checksum          string  //`json:"checksum"`
	Metadata          map[string][]string
	EnrichedMetadata  struct {
	} //`json:"enrichedMetadata"`
	Children             []any //`json:"children"`
	HighlightedFragments []any //`json:"highlightedFragments"`
}

// Metadata = map[string][]string {
// 	UpdateDate                []string `json:"updateDate"`
// 	EsSortDate                []string `json:"es_SortDate"`
// 	BeneficiaryAdministration []string `json:"beneficiaryAdministration"`
// 	SortStatus                []string `json:"sortStatus"`
// 	ContractType              []string `json:"contractType"`
// 	Language                  []string `json:"language"`
// 	Title                     []string `json:"title"`
// 	Type                      []string `json:"type"`
// 	EsSTChecksum              []string `json:"esST_checksum"`
// 	EsSTFileName              []string `json:"esST_FileName"`
// 	CallIdentifier            []string `json:"callIdentifier"`
// 	Datasource                []string `json:"DATASOURCE"`
// 	Currency                  []string `json:"currency"`
// 	FrameworkProgramme        []string `json:"frameworkProgramme"`
// 	Budget                    []string `json:"budget"`
// 	Identifier                []string `json:"identifier"`
// 	CaName                    []string `json:"caName"`
// 	EsContentType             []string `json:"es_ContentType"`
// 	GeographicalZones         []string `json:"geographicalZones"`
// 	PublicationDocuments      []string `json:"publicationDocuments"`
// 	ProgrammePeriod           []string `json:"programmePeriod"`
// 	DeadlineDate              []string `json:"deadlineDate"`
// 	EsDAIngestDate            []string `json:"esDA_IngestDate"`
// 	URL                       []string `json:"url"`
// 	EsCombine                 []string `json:"es_Combine"`
// 	EsSTURL                   []string `json:"esST_URL"`
// 	EsDAQueueDate             []string `json:"esDA_QueueDate"`
// 	ProgrammeDivisionProspect []string `json:"programmeDivisionProspect"`
// 	GeographicalZone          []string `json:"geographicalZone"`
// 	StartDate                 []string `json:"startDate"`
// 	Status                    []string `json:"status"`
// } `json:"metadata,omitempty"`
