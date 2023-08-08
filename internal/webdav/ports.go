package webdav

import "github.com/chrisfarms/nntp"

type UsenetConnectionPool interface {
	GetConnection() (*nntp.Conn, error)
	CloseConnection(c *nntp.Conn) error
}
