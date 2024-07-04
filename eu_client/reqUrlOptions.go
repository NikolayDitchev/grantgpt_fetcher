package eu_client

import (
	"net/url"
	"strconv"
)

type RequestURLOption func(string) (string, error)

func AddPath(addition string) RequestURLOption {
	return func(s string) (string, error) {
		return s + "/" + addition, nil
	}
}

func WithPageNumber(pageNumber int) RequestURLOption {
	return func(reqUrl string) (string, error) {
		pageNumberString := strconv.Itoa(pageNumber)
		return SetUrlParam("pageNumber", pageNumberString, reqUrl)
	}
}

func WithApiKey(apiKey string) RequestURLOption {
	return func(reqUrl string) (string, error) {
		return SetUrlParam("apiKey", apiKey, reqUrl)
	}
}

func WithText(text string) RequestURLOption {
	return func(reqUrl string) (string, error) {
		return SetUrlParam("text", text, reqUrl)
	}
}

func WithPageSize(pageSize int) RequestURLOption {
	return func(reqUrl string) (string, error) {
		pageSizeString := strconv.Itoa(pageSize)
		return SetUrlParam("pageSize", pageSizeString, reqUrl)
	}
}

func SetUrlParam(name string, value string, reqUrl string) (string, error) {
	url, err := url.Parse(reqUrl)
	if err != nil {
		return reqUrl, err
	}

	query := url.Query()
	query.Set(name, value)

	url.RawQuery = query.Encode()
	return url.String(), nil
}
