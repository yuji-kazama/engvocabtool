package notion

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
}

type Page struct {
	Object     string `json:"object"`
	ID         string `json:"id"`
	Properties struct {
		Meaning struct {
			ID       string `json:"id"`
			Type     string `json:"type"`
			RichText []struct {
				Type string `json:"type"`
				Text struct {
					Content string      `json:"content"`
					Link    interface{} `json:"link"`
				} `json:"text"`
				Annotations struct {
					Bold          bool   `json:"bold"`
					Italic        bool   `json:"italic"`
					Strikethrough bool   `json:"strikethrough"`
					Underline     bool   `json:"underline"`
					Code          bool   `json:"code"`
					Color         string `json:"color"`
				} `json:"annotations"`
				PlainText string      `json:"plain_text"`
				Href      interface{} `json:"href"`
			} `json:"rich_text"`
		} `json:"Meaning"`
		Class struct {
			ID     string `json:"id"`
			Type   string `json:"type"`
			Select struct {
				ID    string `json:"id"`
				Name  string `json:"name"`
				Color string `json:"color"`
			} `json:"select"`
		} `json:"Class"`
		Status struct {
			ID     string `json:"id"`
			Type   string `json:"type"`
			Select struct {
				ID    string `json:"id"`
				Name  string `json:"name"`
				Color string `json:"color"`
			} `json:"select"`
		} `json:"Status"`
		StudyDate struct {
			ID   string `json:"id"`
			Type string `json:"type"`
			Date struct {
				Start    string      `json:"start"`
				End      interface{} `json:"end"`
				TimeZone interface{} `json:"time_zone"`
			} `json:"date"`
		} `json:"Study Date"`
		Sentence struct {
			ID       string `json:"id"`
			Type     string `json:"type"`
			RichText []struct {
				Type string `json:"type"`
				Text struct {
					Content string      `json:"content"`
					Link    interface{} `json:"link"`
				} `json:"text"`
				Annotations struct {
					Bold          bool   `json:"bold"`
					Italic        bool   `json:"italic"`
					Strikethrough bool   `json:"strikethrough"`
					Underline     bool   `json:"underline"`
					Code          bool   `json:"code"`
					Color         string `json:"color"`
				} `json:"annotations"`
				PlainText string      `json:"plain_text"`
				Href      interface{} `json:"href"`
			} `json:"rich_text"`
		} `json:"Sentence"`
		ReviewDate struct {
			ID      string `json:"id"`
			Type    string `json:"type"`
			Formula struct {
				Type string `json:"type"`
				Date struct {
					Start    string      `json:"start"`
					End      interface{} `json:"end"`
					TimeZone interface{} `json:"time_zone"`
				} `json:"date"`
			} `json:"formula"`
		} `json:"Review Date"`
		Note struct {
			ID       string `json:"id"`
			Type     string `json:"type"`
			RichText []struct {
				Type string `json:"type"`
				Text struct {
					Content string      `json:"content"`
					Link    interface{} `json:"link"`
				} `json:"text"`
				Annotations struct {
					Bold          bool   `json:"bold"`
					Italic        bool   `json:"italic"`
					Strikethrough bool   `json:"strikethrough"`
					Underline     bool   `json:"underline"`
					Code          bool   `json:"code"`
					Color         string `json:"color"`
				} `json:"annotations"`
				PlainText string      `json:"plain_text"`
				Href      interface{} `json:"href"`
			} `json:"rich_text"`
		} `json:"Note"`
		Image struct {
			ID    string `json:"id"`
			Type  string `json:"type"`
			Files []struct {
				Name     string `json:"name"`
				Type     string `json:"type"`
				External struct {
					URL string `json:"url"`
				} `json:"external"`
			} `json:"files"`
		} `json:"Image"`
		CheckNum struct {
			ID     string `json:"id"`
			Type   string `json:"type"`
			Select struct {
				ID    string `json:"id"`
				Name  string `json:"name"`
				Color string `json:"color"`
			} `json:"select"`
		} `json:"Check Num"`
		Frequency struct {
			ID     string `json:"ID"`
			Type   string `json:"type"`
			Number int    `json:"number"`
		} `json:"Frequency"`
		Name struct {
			ID    string `json:"id"`
			Type  string `json:"type"`
			Title []struct {
				Type string `json:"type"`
				Text struct {
					Content string      `json:"content"`
					Link    interface{} `json:"link"`
				} `json:"text"`
				Annotations struct {
					Bold          bool   `json:"bold"`
					Italic        bool   `json:"italic"`
					Strikethrough bool   `json:"strikethrough"`
					Underline     bool   `json:"underline"`
					Code          bool   `json:"code"`
					Color         string `json:"color"`
				} `json:"annotations"`
				PlainText string      `json:"plain_text"`
				Href      interface{} `json:"href"`
			} `json:"title"`
		} `json:"Name"`
	} `json:"properties"`
	URL string `json:"url"`
}

type CreateResult struct {
	Object         string    `json:"object"`
	ID             string    `json:"id"`
	URL string `json:"url"`
}

type SearchResult struct {
	Object  string `json:"object"`
	Results []struct {
		Object         string    `json:"object"`
		ID             string    `json:"id"`
		CreatedBy      struct {
			Object string `json:"object"`
			ID     string `json:"id"`
		} `json:"created_by"`
		LastEditedBy struct {
			Object string `json:"object"`
			ID     string `json:"id"`
		} `json:"last_edited_by"`
		Cover  interface{} `json:"cover"`
		Icon   interface{} `json:"icon"`
		Parent struct {
			Type       string `json:"type"`
			DatabaseID string `json:"database_id"`
		} `json:"parent"`
		Archived   bool `json:"archived"`
		Properties struct {
			Meaning struct {
				ID       string `json:"id"`
				Type     string `json:"type"`
				RichText []struct {
					Type string `json:"type"`
					Text struct {
						Content string      `json:"content"`
						Link    interface{} `json:"link"`
					} `json:"text"`
					Annotations struct {
						Bold          bool   `json:"bold"`
						Italic        bool   `json:"italic"`
						Strikethrough bool   `json:"strikethrough"`
						Underline     bool   `json:"underline"`
						Code          bool   `json:"code"`
						Color         string `json:"color"`
					} `json:"annotations"`
					PlainText string      `json:"plain_text"`
					Href      interface{} `json:"href"`
				} `json:"rich_text"`
			} `json:"Meaning"`
			Class struct {
				ID     string `json:"id"`
				Type   string `json:"type"`
				Select struct {
					ID    string `json:"id"`
					Name  string `json:"name"`
					Color string `json:"color"`
				} `json:"select"`
			} `json:"Class"`
			Status struct {
				ID     string `json:"id"`
				Type   string `json:"type"`
				Select struct {
					ID    string `json:"id"`
					Name  string `json:"name"`
					Color string `json:"color"`
				} `json:"select"`
			} `json:"Status"`
			StudyDate struct {
				ID   string `json:"id"`
				Type string `json:"type"`
				Date struct {
					Start    string      `json:"start"`
					End      interface{} `json:"end"`
					TimeZone interface{} `json:"time_zone"`
				} `json:"date"`
			} `json:"Study Date"`
			Media struct {
				ID       string        `json:"id"`
				Type     string        `json:"type"`
				Relation []interface{} `json:"relation"`
			} `json:"Media"`
			Sentence struct {
				ID       string        `json:"id"`
				Type     string        `json:"type"`
				RichText []interface{} `json:"rich_text"`
			} `json:"Sentence"`
			ReviewDate struct {
				ID      string `json:"id"`
				Type    string `json:"type"`
				Formula struct {
					Type string `json:"type"`
					Date struct {
						Start    string      `json:"start"`
						End      interface{} `json:"end"`
						TimeZone interface{} `json:"time_zone"`
					} `json:"date"`
				} `json:"formula"`
			} `json:"Review Date"`
			Note struct {
				ID       string        `json:"id"`
				Type     string        `json:"type"`
				RichText []interface{} `json:"rich_text"`
			} `json:"Note"`
			Image struct {
				ID    string        `json:"id"`
				Type  string        `json:"type"`
				Files []interface{} `json:"files"`
			} `json:"Image"`
			CheckNum struct {
				ID     string `json:"id"`
				Type   string `json:"type"`
				Select struct {
					ID    string `json:"id"`
					Name  string `json:"name"`
					Color string `json:"color"`
				} `json:"select"`
			} `json:"Check Num"`
			Book struct {
				ID       string        `json:"id"`
				Type     string        `json:"type"`
				Relation []interface{} `json:"relation"`
			} `json:"Book"`
			CreatedTime struct {
				ID          string    `json:"id"`
				Type        string    `json:"type"`
			} `json:"Created time"`
			CoursesTraining struct {
				ID       string        `json:"id"`
				Type     string        `json:"type"`
				Relation []interface{} `json:"relation"`
			} `json:"Courses & Training"`
			Frequency struct {
				ID     string `json:"id"`
				Type   string `json:"type"`
				Number int    `json:"number"`
			} `json:"Frequency"`
			Name struct {
				ID    string `json:"id"`
				Type  string `json:"type"`
				Title []struct {
					Type string `json:"type"`
					Text struct {
						Content string      `json:"content"`
						Link    interface{} `json:"link"`
					} `json:"text"`
					Annotations struct {
						Bold          bool   `json:"bold"`
						Italic        bool   `json:"italic"`
						Strikethrough bool   `json:"strikethrough"`
						Underline     bool   `json:"underline"`
						Code          bool   `json:"code"`
						Color         string `json:"color"`
					} `json:"annotations"`
					PlainText string      `json:"plain_text"`
					Href      interface{} `json:"href"`
				} `json:"title"`
			} `json:"Name"`
		} `json:"properties"`
		URL string `json:"url"`
	} `json:"results"`
	NextCursor interface{} `json:"next_cursor"`
	HasMore    bool        `json:"has_more"`
}

func NewClient() *Client {
	c := new(Client)
	c.baseURL = "https://api.notion.com/v1"
	c.httpClient = new(http.Client)
	return c
}

func (c *Client) newRequest(method, spath string, body io.Reader) (*http.Request, error) {
	url := c.baseURL + spath
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+os.Getenv("NOTION_INTEGRATION_TOKEN"))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Notion-Version", "2021-05-13")
	return req, nil
}

func (c *Client) GetPage(pageId string) (*Page, error) {
	req, err := c.newRequest(http.MethodGet, "/pages/"+pageId, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.httpClient.Do(req)
	if verifyResponse(res, err) != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Read Error:", err)
		return nil, err
	}
	// fmt.Print(string(body))
	var page Page
	if err := json.Unmarshal(body, &page); err != nil {
		fmt.Printf("Can not unmarshal JSON: %v", err)
		return nil, err
	}
	return &page, nil
}

func (c *Client) UpdatePage(pageId string, data string) (*CreateResult, error) {
	req, err := c.newRequest(http.MethodPatch, "/pages/" + pageId, bytes.NewBuffer([]byte(data)))
	if err != nil {
		return nil, err
	}
	return doRequest(c, req)
}

func (c *Client) CreatePage(data string) (*CreateResult, error) {
	req, err := c.newRequest(http.MethodPost, "/pages", bytes.NewBuffer(([]byte(data))))
	if err != nil {
		return nil, err
	}
	return doRequest(c, req)
}

func doRequest(c *Client, req *http.Request) (*CreateResult, error) {
	res, err := c.httpClient.Do(req)
	if verifyResponse(res, err) != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var result CreateResult
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Printf("cannot unmarshal json: %v", err)
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetAllPages() (*SearchResult, error) {
	query := `{
		"filter": {
			"property": "Name",
			"title": {
				"is_not_empty": true
			}
		}
	}`
	return search(c, query)
}

func (c *Client) GetPageByName(name string) (*SearchResult, error) {
	query := `{
		"filter": {
			"property": "Name",
			"title": {
				"equals": "`+ name + `"
			}
		}
	}`
	return search(c, query)
}

func search(c *Client, query string) (*SearchResult, error) {
	req, err := c.newRequest(
		http.MethodPost,
		"/databases/" + os.Getenv("NOTION_DATABASE_ID") + "/query",
		bytes.NewBuffer([]byte(query)))
	if err != nil {
		return nil, err
	}
	res, err := c.httpClient.Do(req)
	if verifyResponse(res, err) != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var result SearchResult
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Printf("cannot unmarshal json: %v", err)
		return nil, err
	}
	return &result, nil
}

func (c *Client) Exist(name string) bool {
	res, _ := c.GetPageByName(name)
	return len(res.Results) > 0
}

func verifyResponse(res *http.Response, err error) (error) {
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		return fmt.Errorf("request error: %v", res)
	}
	return nil
}