package usenet

import "crypto/tls"

type Config struct {
	Host      string
	Port      int
	Username  string
	Password  string
	Group     string
	SSL       bool
	TLSConfig *tls.Config
}

type Option func(*Config)

func defaultConfig() *Config {
	return &Config{
		SSL: false,
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

func WithSSL(ssl bool) Option {
	return func(c *Config) {
		c.SSL = ssl
	}
}

func WithTLSConfig(tlsConfig *tls.Config) Option {
	return func(c *Config) {
		c.TLSConfig = tlsConfig
	}
}
