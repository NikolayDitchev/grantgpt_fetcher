package danswer_api

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type DanswerAPIClient struct {
	client *http.Client
}

func NewDanswerAPIClient() *DanswerAPIClient {
	return &DanswerAPIClient{
		client: &http.Client{},
	}
}

func (dc DanswerAPIClient) GetConnectors() (*[]Connector, error) {

	url, err := url.Parse(base_url + get_connectors)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := dc.sendRequest(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()

	connectors := []Connector{}
	err = json.Unmarshal(body, &connectors)
	if err != nil {
		return nil, err
	}

	return &connectors, nil
}

func (dc DanswerAPIClient) sendRequest(req *http.Request) (*http.Response, error) {

	cookie, err := dc.getAuthCookie()

	if err != nil {
		return nil, err
	}

	req.AddCookie(cookie)

	resp, err := dc.client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (dc DanswerAPIClient) getAuthCookie() (*http.Cookie, error) {

	cookie := &http.Cookie{
		Name:  "fastapiusersauth",
		Value: "nJeXEy6WHbFsLPnwIvdxUwq9mKdsccLERL3Cz8q3qZ4",
	}

	return cookie, nil
}
