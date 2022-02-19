package weatherstack

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

const Url = "http://api.weatherstack.com/"

type Client struct {
	log    *logrus.Logger
	client *http.Client
}

func NewClient(log *logrus.Logger) *Client {
	return &Client{
		client: http.DefaultClient,
		log:    log,
	}
}

func (c *Client) GetCurrent(ctx context.Context, accessKey string, city string) (*GetCurrentResponse, error) {
	result := &GetCurrentResponse{}
	resp, err := c.client.Get(fmt.Sprintf("%s?access_key=%s&query=%s", Url+"/current", accessKey, city))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non 200 HTTP code")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err = json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
}
