package newsapi

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

const (
	baseUrl              = "https://newsapi.org/v2/"
	everythingEndpoint   = "everything"
	topHeadlinesEndpoint = "top-headlines"
	sourcesEndpoint      = "sources"
	apiKeyHeader         = "X-Api-Key"
)

type Client struct {
	client *http.Client
	apiKey string
}

func NewClient(apiKey string) *Client {
	c := &Client{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		apiKey: apiKey,
	}
	return c
}

func (c *Client) GetTopHeadlines(params TopHeadlinesRequestParams) (*NewsResponse, error) {
	req, err := c.buildRequest(topHeadlinesEndpoint, params)
	if err != nil {
		return nil, err
	}

	res := struct {
		*ErrorResponse
		*NewsResponse
	}{}

	_, err = c.sendRequestAndDecode(req, &res)
	if err != nil {
		return nil, err
	}

	if res.ErrorResponse != nil {
		return nil, res.ErrorResponse
	}

	return res.NewsResponse, nil
}

func (c *Client) GetEverything(params EverythingRequestParams) (*NewsResponse, error) {
	req, err := c.buildRequest(everythingEndpoint, params)
	if err != nil {
		return nil, err
	}

	res := struct {
		*ErrorResponse
		*NewsResponse
	}{}

	_, err = c.sendRequestAndDecode(req, &res)
	if err != nil {
		return nil, err
	}

	if res.ErrorResponse != nil {
		return nil, res.ErrorResponse
	}

	return res.NewsResponse, nil
}

func (c *Client) GetSources(params SourcesRequestParams) (*SourcesResponse, error) {
	req, err := c.buildRequest(sourcesEndpoint, params)
	if err != nil {
		return nil, err
	}

	res := struct {
		*ErrorResponse
		*SourcesResponse
	}{}

	_, err = c.sendRequestAndDecode(req, &res)
	if err != nil {
		return nil, err
	}

	if res.ErrorResponse != nil {
		return nil, res.ErrorResponse
	}

	return res.SourcesResponse, nil
}

func (c *Client) sendRequestAndDecode(req *http.Request, v interface{}) (*http.Response, error) {
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(res.Body).Decode(v)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) buildRequest(endpoint string, params interface{}) (*http.Request, error) {

	v, err := query.Values(params)
	if err != nil {
		return nil, err
	}
	u := baseUrl + endpoint + "?" + v.Encode()

	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set(apiKeyHeader, c.apiKey)
	return req, nil
}
