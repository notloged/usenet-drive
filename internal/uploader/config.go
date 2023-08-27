package uploader

import (
	"strings"
)

type Config struct {
	Host           string
	Port           int
	Username       string
	Password       string
	Groups         []string
	SSL            bool
	MaxConnections int
	fileWhiteList  []string
	nyuuPath       string
	articleSize    string
}

func (c *Config) getGroups() string {
	return strings.Join(c.Groups, ",")
}

type Option func(*Config)

func defaultConfig() *Config {
	return &Config{
		SSL:         false,
		articleSize: "750K",
	}
}

func WithHost(host string) Option {
	return func(c *Config) {
		c.Host = host
	}
}

func WithPort(port int) Option {
	return func(c *Config) {
		c.Port = port
	}
}

func WithUsername(username string) Option {
	return func(c *Config) {
		c.Username = username
	}
}

func WithPassword(password string) Option {
	return func(c *Config) {
		c.Password = password
	}
}

func WithGroups(groups []string) Option {
	return func(c *Config) {
		c.Groups = groups
	}
}

func WithSSL(ssl bool) Option {
	return func(c *Config) {
		c.SSL = ssl
	}
}

func WithMaxConnections(maxConnections int) Option {
	return func(c *Config) {
		c.MaxConnections = maxConnections
	}
}

func WithFileWhiteList(fileWhiteList []string) Option {
	return func(c *Config) {
		c.fileWhiteList = fileWhiteList
	}
}

func WithNyuuPath(nyuuPath string) Option {
	return func(c *Config) {
		c.nyuuPath = nyuuPath
	}
}

func WithArticleSize(articleSize string) Option {
	return func(c *Config) {
		c.articleSize = articleSize
	}
}
