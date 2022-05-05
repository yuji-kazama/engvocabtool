package words

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"reflect"
)
const (
	apiURL = "https://wordsapiv1.p.rapidapi.com"
)

type ClientOption func(*Client)

type Client struct {
	baseUrl *url.URL
	httpClient *http.Client
	xRapidApiHost string
	xRapidApiKey string
}

type Response struct {
	Word    string `json:"word"`
	Results []Result `json:"results"`
	Frequency float64 `json:"frequency"`
	Pronunciation Pronunciation `json:"pronunciation"`
}

type Result struct {
	Definition   string   `json:"definition"`
	PartOfSpeech string   `json:"partOfSpeech"`
	Synonyms     []string `json:"synonyms,omitempty"`
	Examples     []string `json:"examples,omitempty"`
}

type Pronunciation struct {
	All string `json:"all,omitempty"`
	Noun string `json:"noun,omitempty"`
	Verb string `json:"verb,omitempty"`
	Adjective string `json:"adjective,omitempty"`
	Adverb string `json:"adverb,omitempty"`
	Conjunction string `json:"conjunction,omitempty"`
}

func NewClient(opts ...ClientOption) (*Client) {
	u, err := url.Parse(apiURL)
	if err != nil {
		panic(err)
	}
	c := &Client{
		httpClient: http.DefaultClient,
		baseUrl: u,
		xRapidApiHost: os.Getenv("WORDS_XRAPIDAPI_HOST"),
		xRapidApiKey: os.Getenv("WORDS_XRAPIDAPI_KEY"),
	}
	c.httpClient = new(http.Client)

	for _, opt := range opts {
		opt(c)
	}
	return c
}

func WithHttpClient(client *http.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = client
	}
}

func (c *Client) request(ctx context.Context, method string, urlStr string, requestBody interface{}) (*http.Response, error) {
	u, err := c.baseUrl.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if requestBody != nil && !reflect.ValueOf(requestBody).IsNil() {
		body, err := json.Marshal(requestBody)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(body)
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-RapidAPI-Host", c.xRapidApiHost)
	req.Header.Set("X-RapidAPI-Key", c.xRapidApiKey)

	res, err := c.httpClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error code: %v", res.StatusCode)
	}
	return res, nil
}

func (c *Client) GetEverything(ctx context.Context, word string) (*Response, error) {
	res, err := c.request(ctx, http.MethodGet, fmt.Sprintf("/words/%s", word), nil)
	if err != nil {
		return nil, err
	}

	defer func() {
		if errClose := res.Body.Close(); errClose != nil {
			fmt.Println("failed to close body, should never happen")
		}
	}()

	var response Response

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}