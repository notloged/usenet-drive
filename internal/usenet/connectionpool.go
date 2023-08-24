//Copyright 2013, Daniel Morsing
//For licensing information, See the LICENSE file

// This file contains a muxer that will limit the amount of connections
// that are concurrently running.

package usenet

import (
	"log"
	"net"
	"sync"

	"github.com/chrisfarms/nntp"
)

type connectionPool struct {
	mu               sync.Mutex
	connectionsTaken int
	ch               chan *nntp.Conn
	maxConnections   int
	config           *Config
}

type connectionError struct {
	c   *nntp.Conn
	err error
}

func NewConnectionPool(options ...Option) *connectionPool {
	config := defaultConfig()
	for _, option := range options {
		option(config)
	}

	return &connectionPool{
		ch:             make(chan *nntp.Conn, config.MaxConnections),
		maxConnections: config.MaxConnections,
		config:         config,
	}
}

func (p *connectionPool) GetConnection() (*nntp.Conn, error) {
	// check if there's a free conn we can get
	select {
	case c := <-p.ch:
		return c, nil
	default:
	}
	p.mu.Lock()
	if p.connectionsTaken == p.maxConnections {
		p.mu.Unlock()
		// wait for idle conn
		c := <-p.ch
		return c, nil
	}
	p.connectionsTaken++
	p.mu.Unlock()
	ch := make(chan connectionError)
	cancelCh := make(chan struct{})

	// dial with this connection.
	// if we manage to get a connection from
	// a client done with theirs, we will use that one
	// and put the idle conn
	go func() {
		c, err := p.dialNNTP()
		select {
		case <-cancelCh:
			if err == nil {
				p.FreeConnection(c)
				return
			}
			// ignore error
			p.mu.Lock()
			p.connectionsTaken--
			p.mu.Unlock()
		case ch <- connectionError{c, err}:
		}
	}()
	select {
	case ce := <-ch:
		return ce.c, ce.err
	case c := <-p.ch:
		close(cancelCh)
		return c, nil
	}
}

func (p *connectionPool) CloseConnection(c *nntp.Conn) error {
	err := c.Quit()
	if err != nil {
		return err
	}
	p.mu.Lock()
	p.connectionsTaken--
	p.mu.Unlock()

	return nil
}

func (p *connectionPool) FreeConnection(c *nntp.Conn) {
	p.ch <- c
}

func (p *connectionPool) dialNNTP() (*nntp.Conn, error) {
	dialStr := p.config.getConnectionString()
	var err error
	var c *nntp.Conn

	for {
		if p.config.TLS {
			c, err = nntp.DialTLS("tcp", dialStr, p.config.TLSConfig)
		} else {
			c, err = nntp.Dial("tcp", dialStr)
		}
		if err != nil {
			// if it's a timeout, ignore and try again
			e, ok := err.(net.Error)
			if ok && e.Timeout() {
				log.Default().Printf("timeout connecting to %s, retrying", dialStr)
				continue
			}
			return nil, err
		}

		// auth
		if err := c.Authenticate(p.config.Username, p.config.Password); err != nil {
			return nil, err
		}

		break
	}
	return c, nil
}
