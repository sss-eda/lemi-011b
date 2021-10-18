package rest

import (
	"context"
	"encoding/json"
	"io"
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
	r, w := io.Pipe()
	enc := json.NewEncoder(w)

	req, err := http.NewRequest("POST", client.url+"/datum", r)
	if err != nil {
		log.Println(err)
	}

	err = enc.Encode(datum)
	if err != nil {
		log.Println(err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.api.Do(req.WithContext(ctx))
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
