package heartbeat

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
)

type client struct{}

func Client() *client {
	return &client{}
}

func (c *client) Check(addr string) (Message, error) {
	resp, err := http.Get(addr)
	if err != nil {
		return Message{}, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Message{}, err
	}
	defer resp.Body.Close()

	var m Message
	err = json.Unmarshal(b, &m)
	if err != nil {
		return Message{}, err
	}

	return m, nil
}
