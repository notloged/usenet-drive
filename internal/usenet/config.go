package usenet

import (
	"crypto/tls"
	"fmt"
)

type Config struct {
	Host           string
	Port           int
	Username       string
	Password       string
	Group          string
	TLS            bool
	TLSConfig      *tls.Config
	MaxConnections int
}

func (c *Config) getConnectionString() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

type Option func(*Config)

func defaultConfig() *Config {
	return &Config{
		TLS: false,
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

func WithGroup(group string) Option {
	return func(c *Config) {
		c.Group = group
	}
}

func WithTLS(tls bool) Option {
	return func(c *Config) {
		c.TLS = tls
	}
}

func WithTLSConfig(tlsConfig *tls.Config) Option {
	return func(c *Config) {
		c.TLSConfig = tlsConfig
	}
}

func WithMaxConnections(maxConnections int) Option {
	return func(c *Config) {
		c.MaxConnections = maxConnections
	}
}
