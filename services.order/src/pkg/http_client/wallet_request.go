package httpclient

import (
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

func NewCustomerClient() *Client {
	return &Client{
		hostURL: config.AppConfig.WalletServiceURL,
	}
}

func (c *Client) GetWalletInfo(userId int, token string) (*dto.WalletDto, error) {

	req, err := http.NewRequest("GET", c.hostURL+fmt.Sprintf("/%d", 1), nil)
	req.Header.Add("access_token", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response -", err)
		return nil, errors.New("Error occured when fetching user wallet AAAA.")
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
