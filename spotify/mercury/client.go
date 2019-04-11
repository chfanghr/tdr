package mercury

import "bytes"

type Client struct{}

func (c *Client) Handle(cmd uint8, reader *bytes.Reader) error { return nil } //TODO
