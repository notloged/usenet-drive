package usenet

import (
	"fmt"

	"github.com/chrisfarms/nntp"
)

type client struct {
	connection nntp.Conn
}

func NewClient(options ...Option) (*client, error) {
	config := defaultConfig()
	for _, option := range options {
		option(config)
	}

	var conn *nntp.Conn
	if config.SSL {
		c, err := nntp.DialTLS("tcp", fmt.Sprintf("%s:%d", config.Host, config.Port), config.TLSConfig)
		if err != nil {
			return nil, fmt.Errorf("connection failed: %v", err)
		}
		conn = c
	} else {
		c, err := nntp.Dial("tcp", fmt.Sprintf("%s:%d", config.Host, config.Port))
		if err != nil {
			return nil, fmt.Errorf("connection failed: %v", err)
		}
		conn = c
	}

	// auth
	if err := conn.Authenticate(config.Username, config.Password); err != nil {
		return nil, fmt.Errorf("could not authenticate: %v", err)
	}

	// connect to a news group
	_, l, _, err := conn.Group(config.Group)
	if err != nil {
		return nil, fmt.Errorf("could not connect to group %s: %v %d", config.Group, err, l)
	}

	return &client{
		connection: *conn,
	}, nil
}

func (c *client) GetArticle(id string) (*nntp.Article, error) {
	// get article
	article, err := c.connection.Article(id)
	if err != nil {
		return nil, fmt.Errorf("could not get article %s: %v", id, err)
	}

	return article, nil
}
