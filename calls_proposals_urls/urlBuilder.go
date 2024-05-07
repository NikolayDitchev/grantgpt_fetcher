package calls_proposals_urls

import (
	"fetcher/api_caller"
	"strings"
)

const (
	topicUrl       string = `https://ec.europa.eu/info/funding-tenders/opportunities/portal/screen/opportunities/topic-details/`
	grantUrl       string = `https://ec.europa.eu/info/funding-tenders/opportunities/portal/screen/opportunities/prospect-details/`
	cascadeFundUrl string = `https://ec.europa.eu/info/funding-tenders/opportunities/portal/screen/opportunities/competitive-calls-cs/`
)

type UrlBuilingFunc func(result *api_caller.Result) (string, error)

func getUrlBuilders() map[string]UrlBuilingFunc {
	return map[string]UrlBuilingFunc{
		"1": topicUrlBuilder,
		"2": grantUrlBuilder,
		"8": cascadeGrantUrlBuilder,
	}
}

func topicUrlBuilder(result *api_caller.Result) (string, error) {

	identifier := strings.ToLower(result.Metadata["identifier"][0])
	url := topicUrl + identifier

	return url, nil
}

func grantUrlBuilder(result *api_caller.Result) (string, error) {

	callIdentifier := result.Metadata["callIdentifier"][0]

	url := grantUrl + callIdentifier + "PROSPECTSEN"

	return url, nil
}

func cascadeGrantUrlBuilder(result *api_caller.Result) (string, error) {

	url := cascadeFundUrl + result.Metadata["callccm2Id"][0]

	return url, nil
}
