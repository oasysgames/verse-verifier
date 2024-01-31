package ipc

import (
	"errors"
	"fmt"
	"io"
	"time"

	goipc "github.com/james-barrow/golang-ipc"
)

func NewClient(sockname string, msgType int) (c *Client, err error) {
	if msgType == EOM {
		return nil, fmt.Errorf("message type %d is reserved as EOM", EOM)
	}

	client, err := goipc.StartClient(sockname, nil)
	if err != nil {
		return nil, err
	}

	c = &Client{client, msgType}
	if !c.ready() {
		return nil, errors.New("IPC is not ready")
	}

	return c, nil
}

type Client struct {
	client  *goipc.Client
	msgType int
}

func (c *Client) ready() bool {
	for i := 0; i < 10; i++ {
		c.client.Read()
		if c.client.Status() == "Connected" {
			return true
		}
		time.Sleep(time.Second / 4)
	}
	return false
}

func (c *Client) Read() ([]byte, error) {
	for {
		msg, err := c.client.Read()
		if err != nil {
			return nil, err
		}

		switch msg.MsgType {
		case -1:
			continue // changed server status
		case -2:
			return nil, errors.New("error on server")
		case EOM:
			return nil, io.EOF
		case c.msgType:
			return msg.Data, nil
		default:
			return nil, fmt.Errorf("%d is unknown message type", msg.MsgType)
		}
	}
}

func (c *Client) Write(message []byte) error {
	return c.client.Write(c.msgType, message)
}

func (c *Client) Close() error {
	c.client.Close()
	return nil
}
