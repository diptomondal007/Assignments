package hn

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	apiBase = "https://hacker-news.firebaseio.com/v0"
)

//Client ...
type Client struct {
	apiBase string
}

func (c *Client) defaultify() {
	if c.apiBase == "" {
		c.apiBase = apiBase
	}
}

//TopItems ...
func (c *Client) TopItems() ([]int, error) {
	c.defaultify()
	log.Println(c.apiBase)
	resp, err := http.Get(fmt.Sprintf("%s/topstories.json", c.apiBase))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var ids []int
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&ids)
	if err != nil {
		return nil, err
	}
	return ids, nil
}

//GetItem ...
func (c *Client) GetItem(id int) (Item, error) {
	c.defaultify()
	var item Item
	resp, err := http.Get(fmt.Sprintf("%s/item/%d.json", c.apiBase, id))
	if err != nil {
		return item, err
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&item)
	if err != nil {
		return item, err
	}
	return item, nil
}

//Item ...
type Item struct {
	By          string `json:"by"`
	Descendants int    `json:"descendants"`
	ID          int    `json:"id"`
	Kids        []int  `json:"kids"`
	Score       int    `json:"score"`
	Time        int    `json:"time"`
	Title       string `json:"title"`
	Type        string `json:"type"`

	// Only one of these should exist
	Text string `json:"text"`
	URL  string `json:"url"`
}
