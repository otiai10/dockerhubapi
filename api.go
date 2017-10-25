package dockerhubapi

import (
	"encoding/json"
	"net/http"
	"path"
)

const (
	// BaseURL ...
	BaseURL = "https://hub.docker.com/"
)

// API ...
type API struct {
	HTTPClient *http.Client
	Version    string
	BaseURL    string
	Verbose    bool
}

// New ...
func New(client ...*http.Client) *API {
	if len(client) == 0 {
		client = append(client, http.DefaultClient)
	}
	return &API{
		HTTPClient: client[0],
		Version:    "v2",
		BaseURL:    BaseURL,
	}
}

// Fetch ...
func (api *API) Fetch(resource Resource) error {
	location := api.BaseURL + path.Join(api.Version, resource.Path()) + "/"
	req, err := http.NewRequest("GET", location, nil)
	if err != nil {
		return err
	}
	res, err := api.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(resource); err != nil {
		return err
	}
	return nil
}
