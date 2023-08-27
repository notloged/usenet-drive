package webdav

type Config struct {
	NzbPath    string
	ServerPort string
}

type Option func(*Config)

func defaultConfig() *Config {
	return &Config{
		ServerPort: "8080",
	}
}

func WithNzbPath(nzbPath string) Option {
	return func(c *Config) {
		c.NzbPath = nzbPath
	}
}

func WithServerPort(serverPort string) Option {
	return func(c *Config) {
		c.ServerPort = serverPort
	}
}
