package usenetfilereader

import (
	"log/slog"

	connectionpool "github.com/javi11/usenet-drive/internal/usenet/connection-pool"
	corruptednzbsmanager "github.com/javi11/usenet-drive/internal/usenet/corrupted-nzbs-manager"
	"github.com/javi11/usenet-drive/internal/usenet/nzbloader"
)

type Config struct {
	cp        connectionpool.UsenetConnectionPool
	log       *slog.Logger
	nzbLoader *nzbloader.NzbLoader
	cNzb      corruptednzbsmanager.CorruptedNzbsManager
}

type Option func(*Config)

func defaultConfig() *Config {
	return &Config{}
}

func WithConnectionPool(cp connectionpool.UsenetConnectionPool) Option {
	return func(c *Config) {
		c.cp = cp
	}
}

func WithLogger(log *slog.Logger) Option {
	return func(c *Config) {
		c.log = log
	}
}

func WithNzbLoader(nzbLoader *nzbloader.NzbLoader) Option {
	return func(c *Config) {
		c.nzbLoader = nzbLoader
	}
}

func WithCorruptedNzbsManager(cNzb corruptednzbsmanager.CorruptedNzbsManager) Option {
	return func(c *Config) {
		c.cNzb = cNzb
	}
}
