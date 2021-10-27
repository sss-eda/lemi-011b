package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/sss-eda/lemi-011b/pkg/domain/acquisition"
	"github.com/sss-eda/lemi-011b/pkg/domain/registration"
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

// RegisterSensor TODO
func (client *Client) RegisterSensor(
	ctx context.Context,
	sensor registration.Sensor,
) error {
	jsonData, err := json.Marshal(sensor)
	if err != nil {
		log.Println(err)
	}

	resp, err := client.api.Post(
		client.url+"/sensor",
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
