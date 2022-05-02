package words

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type ClientOption func(*Client)

type Client struct {
	baseURL string
	httpClient *http.Client
}

type Response struct {
	Word    string `json:"word"`
	Results []Result `json:"results"`
	Frequency float64 `json:"frequency"`
}

type Result struct {
	Definition   string   `json:"definition"`
	PartOfSpeech string   `json:"partOfSpeech"`
	Synonyms     []string `json:"synonyms,omitempty"`
	Examples     []string `json:"examples,omitempty"`
}

func NewClient(opts ...ClientOption) (*Client) {
	c := new(Client)
	c.baseURL = "https://wordsapiv1.p.rapidapi.com/words/"
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

func (c *Client) newRequest(method, spath string, body io.Reader) (*http.Request, error) {
	url := c.baseURL + spath
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-RapidAPI-Host", os.Getenv("WORDS_XRAPIDAPI_HOST"))
	req.Header.Set("X-RapidAPI-Key", os.Getenv("WORDS_XRAPIDAPI_KEY"))
	return req, nil
}

func (c *Client) GetEverything(word string) (*Response, error) {
	req, err := c.newRequest(http.MethodGet, word, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode == 404 {
		return nil, fmt.Errorf("no matching word was found")
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("request error: %v", res)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("read error: %v", err)
		return nil, err
	}

	var result Response
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Printf("cannot unmarshal json: %v", err)
		return nil, err
	}
	return &result, nil
}