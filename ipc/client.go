package ipc

import (
	"errors"
	"time"

	goipc "github.com/james-barrow/golang-ipc"
)

type ReaderFunc func(*goipc.Client, []byte)

func NewClient(listen string, id int) (*Client, error) {
	client, err := goipc.StartClient(listen, nil)
	if err != nil {
		return nil, err
	}

	c := &Client{Client: client, id: id}
	if !c.ready() {
		return nil, errors.New("IPC is not ready")
	}

	return c, nil
}

type Client struct {
	*goipc.Client

	id int
}

func (c *Client) ready() bool {
	for i := 0; i < 10; i++ {
		c.Client.Read()
		if c.Client.Status() == "Connected" {
			return true
		}
		time.Sleep(time.Second / 4)
	}
	return false
}

func (c *Client) Read() ([]byte, error) {
	for {
		msg, err := c.Client.Read()
		if err != nil {
			return nil, err
		}

		if msg.MsgType == -1 {
			// changed server status
			continue
		}

		if msg.MsgType == -2 {
			return nil, errors.New("error on server")
		}

		if msg.MsgType == c.id {
			return msg.Data, nil
		}
	}
}
