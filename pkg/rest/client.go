package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/sss-eda/lemi-011b/pkg/acquisition"
	"github.com/sss-eda/lemi-011b/pkg/registration"
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
	jsonData, err := json.Marshal(datum)
	if err != nil {
		log.Println(err)
	}

	resp, err := client.api.Post(
		client.url+"/datum",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	log.Println(resp)

	// _, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Println(err)
	// }

	return nil
}

// RegisterInstrument TODO
func (client *Client) RegisterInstrument(
	ctx context.Context,
	instrument registration.Instrument,
) error {
	jsonData, err := json.Marshal(instrument)
	if err != nil {
		log.Println(err)
	}

	resp, err := client.api.Post(
		client.url+"/instrument",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	log.Println(resp)

	// _, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Println(err)
	// }

	return nil
}
