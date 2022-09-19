package httpclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aliaydins/microservice/service.order/src/config"
	"github.com/aliaydins/microservice/service.order/src/dto"
	"io/ioutil"
	"log"
	"net/http"
)

type Client struct {
	hostURL string
}

type walletUpdate struct {
	USD int `json:"usd"`
	BTC int `json:"btc"`
}

func NewCustomerClient() *Client {
	return &Client{
		hostURL: config.AppConfig.WalletServiceURL,
	}
}

func (c *Client) GetWalletInfo(userId int, token string) (*dto.WalletDto, error) {

	req, err := http.NewRequest("GET", c.hostURL+fmt.Sprintf("/%d", userId), nil)
	req.Header.Add("access_token", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response -", err)
		return nil, errors.New("Error occured when fetching user wallet.")
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("Error occured when fetching user wallet info.")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var w dto.WalletDto
	json.Unmarshal(body, &w)
	return &w, nil
}

func (c *Client) WalletUpdate(userId int, usd int, btc int, token string) error {
	body := walletUpdate{
		USD: usd,
		BTC: btc,
	}

	byteBody, err := json.Marshal(body)

	req, err := http.NewRequest("PUT", c.hostURL+fmt.Sprintf("/%d", userId), bytes.NewBuffer(byteBody))
	req.Header.Add("access_token", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response -", err)
		return errors.New("Error occured when updating user wallet.")
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New("Error occured when updating user wallet info.")
	}

	return nil
}
