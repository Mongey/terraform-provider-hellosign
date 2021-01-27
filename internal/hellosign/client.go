package hellosign

import (
	"fmt"
	"sync"

	hs "github.com/StefanNyman/hellosign"
)

type Config struct {
	APIKey string
}

type Client struct {
	once   sync.Once
	config *Config
	client *hs.APIAppAPI
}

func NewClient(config *Config) *Client {
	return &Client{
		config: config,
		client: hs.NewAPIAppAPI(config.APIKey),
	}
}

func (c *Client) init() error {
	var err error

	c.once.Do(func() {
		c.client = hs.NewAPIAppAPI(c.config.APIKey)
	})

	return err
}

func (c *Client) App(clientID string) (*hs.APIApp, error) {
	err := c.init()
	if err != nil {
		return nil, err
	}
	return c.client.Get(clientID)
}

func (c *Client) Create(params hs.APIAppCreateParms) (*hs.APIApp, error) {
	err := c.init()
	if err != nil {
		return nil, err
	}
	return c.client.Create(params)
}

func (c *Client) UpdateApp(clientID string, params hs.APIAppUpdateParms) (*hs.APIApp, error) {
	err := c.init()
	if err != nil {
		return nil, err
	}
	return c.client.Update(clientID, params)
}

func (c *Client) DeleteApp(clientID string) error {
	err := c.init()
	if err != nil {
		return err
	}

	deleted, err := c.client.Delete(clientID)
	if err != nil {
		return err
	}
	if !deleted {
		return fmt.Errorf("Unable to delete")
	}

	return nil
}
