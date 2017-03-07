package calil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const endpoint = "https://api.calil.jp"

type (
	Client struct {
		AppKey string
		Client *http.Client
	}

	SearchLibraryParams struct {
		Pref     string `json:"pref"`
		City     string `json:"city"`
		SystemID string `json:"systemid"`
		Geocode  string `json:"geocode"` // format: `lat,lng`
		Format   string `json:"format"`
		Callback string `json:"callback"`
		Limit    int    `json:"limit"`
	}

	SearchLibraryResult struct {
		Libraries []Library `json:"libraries"`
	}

	Library struct {
		SystemID   string `json:"systemid"`
		SystemName string `json:"systemname"`
		LibKey     string `json:"libkey"`
		LibID      string `json:"libid"`
		Short      string `json:"short"`
		Formal     string `json:"formal"`
		URLPC      string `json:"url_pc"`
		Address    string `json:"address"`
		Pref       string `json:"pref"`
		City       string `json:"city"`
		Post       string `json:"post"`
		Tel        string `json:"tel"`
		Geocode    string `json:"geocode"`
		Category   string `json:"category"`
		Image      string `json:"image"`
		Distance   string `json:"distance"`
	}
)

func NewClient(appkey string, client *http.Client) *Client {
	return &Client{AppKey: appkey, Client: client}
}

func (c *Client) SearchLibrary(p SearchLibraryParams) (SearchLibraryResult, error) {
	const method = "library"

	values := url.Values{
		"appkey":   {c.AppKey},
		"pref":     {p.Pref},
		"city":     {p.City},
		"systemid": {p.SystemID},
		"geocode":  {p.Geocode},
		"format":   {p.Format},
		"callback": {p.Callback},
		"limit":    {fmt.Sprint(p.Limit)},
	}

 	req, err := c.newRequest(http.MethodGet, method, bytes.NewBufferString(values.Encode()))
	if err != nil {
		return SearchLibraryResult{}, err
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return SearchLibraryResult{}, err
	}
	defer resp.Body.Close()

	var ret SearchLibraryResult
	if err := json.NewDecoder(resp.Body).Decode(&ret.Libraries); err != nil {
		return SearchLibraryResult{}, err
	}

	return ret, nil
}

func (c *Client) newRequest(method, spath string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method,  endpoint+"/"+spath, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return req, nil
}
