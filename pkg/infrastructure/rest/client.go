package rest

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/sss-eda/lemi-011b/pkg/domain/acquisition"
)

// Client TODO
type Client struct {
	api *http.Client
	url string
}

// NewClient TODO
func NewClient(url string) (*Client, error) {
	return &Client{
		api: &http.Client{},
		url: url,
	}, nil
}

// AcquireDatum TODO
func (client *Client) AcquireDatum(
	ctx context.Context,
	datum acquisition.Datum,
) error {
	req, err := http.NewRequest("POST", client.url+"/datum", nil)
	if err != nil {
		log.Println(err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.api.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	log.Println(bodyBytes)

	return nil
}
