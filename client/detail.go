package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/secr3t/atp-client/model"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
)

const detailApiName = "item_get"

type DetailClient struct {
	apiKey string
}

func NewDetailClient(apiKey string) *DetailClient {
	return &DetailClient{
		apiKey: apiKey,
	}
}

func (c *DetailClient) GetItems(itemIds []string) []model.DetailItem {
	items := make([]model.DetailItem, 0)
	itemChans := make(chan *model.DetailItem)

	wg := sync.WaitGroup{}
	for _, itemId := range itemIds {
		itemId := itemId
		go func() {
			wg.Add(1)
			result, err := c.getItem(itemId)
			if err == nil {
				result.DetailItem.SetOptions()
				itemChans <- result.DetailItem
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(itemChans)
	}()

	for item := range itemChans {
		items = append(items, *item)
	}

	return items
}

func (c *DetailClient) GetDetails(itemIds []string) chan *model.DetailItem {
	itemLen := len(itemIds)

	var wg sync.WaitGroup
	wg.Add(itemLen)

	itemChans := make(chan *model.DetailItem, itemLen)

	for _, itemId := range itemIds {
		itemId := itemId
		go func() {
			result, err := c.getItem(itemId)
			if err == nil {
				result.DetailItem.SetOptions()
				itemChans <- result.DetailItem
			}
			wg.Done()
		}()
	}

	defer func() {
		go func() {
			wg.Wait()
			close(itemChans)
		}()
	}()

	return itemChans
}

func (c *DetailClient) getItem(itemId string) (model.DetailResult, error) {
	query := c.getDetailQueryParam(itemId)

	reqUri := GetUri(query)

	res, _ := http.Get(reqUri)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var result model.DetailResult

	err := json.Unmarshal(body, &result)

	if err != nil {
		return model.DetailResult{}, err
	}

	if result.IsError() {
		err = errors.New(fmt.Sprintf("%s fetch failed.", itemId))
		return model.DetailResult{}, err
	}

	return result, nil
}

func (c *DetailClient) getDetailQueryParam(itemId string) string {
	p := url.Values{}

	p.Add("api_name", detailApiName)
	p.Add("route", route)
	p.Add("lang", lang)
	p.Add("is_promotion", "!")
	p.Add("key", c.apiKey)
	p.Add("num_iid", itemId)

	return p.Encode()
}
