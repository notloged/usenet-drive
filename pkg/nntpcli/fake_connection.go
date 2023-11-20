package nntpcli

import (
	"bytes"
	"io"
)

type fakeConnection struct {
	connectionType ConnectionType
	providerId     int
	isClosed       bool
}

func NewFakeConnection(host string, providerId int, connectionType ConnectionType) Connection {
	return &fakeConnection{}
}

func (c *fakeConnection) IsClosed() bool {
	return c.isClosed
}

func (c *fakeConnection) ProviderID() int {
	return c.providerId
}
func (c *fakeConnection) GetConnectionType() ConnectionType {
	return c.connectionType
}

func (c *fakeConnection) Authenticate(username, password string) error {
	return nil
}

func (c *fakeConnection) Body(id string) (io.Reader, error) {
	return bytes.NewBuffer([]byte{}), nil
}

func (c *fakeConnection) SelectGroup(group string) (number int, low int, high int, err error) {
	return 0, 0, 0, nil
}

func (c *fakeConnection) Post(p []byte, chunkSize int64) error {
	return nil
}

func (c *fakeConnection) Quit() error {
	c.isClosed = true
	return nil
}
