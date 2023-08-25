package webdav

import "github.com/chrisfarms/nntp"

type UsenetConnectionPool interface {
	Get() (*nntp.Conn, error)
	Close(c *nntp.Conn) error
	Free(c *nntp.Conn) error
}
