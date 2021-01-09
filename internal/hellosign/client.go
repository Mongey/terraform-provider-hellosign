package hellosign

import (
	"fmt"

	hs "github.com/StefanNyman/hellosign"
)

type Config struct {
	APIKey string
}

type Client struct {
	config *Config
	client *hs.APIAppAPI
}

func NewClient(config *Config) *Client {
	return &Client{
		config: config,
		client: hs.NewAPIAppAPI(config.APIKey),
	}
}

func (c *Client) App(clientID string) (*hs.APIApp, error) {
	return c.client.Get(clientID)
}

func (c *Client) Create(params hs.APIAppCreateParms) (*hs.APIApp, error) {
	return c.client.Create(params)
}

func (c *Client) UpdateApp(clientID string) (interface{}, error) {
	return nil, nil
}

func (c *Client) DeleteApp(clientID string) error {
	deleted, err := c.client.Delete(clientID)
	if err != nil {
		return err
	}
	if !deleted {
		return fmt.Errorf("Unable to delete")
	}

	return nil
}
