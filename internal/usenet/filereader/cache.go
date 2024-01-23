package filereader

//go:generate mockgen -source=./cache.go -destination=./cache_mock.go -package=filereader Cache
import (
	"context"
	"log/slog"
	"time"

	"github.com/allegro/bigcache/v3"
)

type Cache interface {
	Get(key string) []byte
	Set(key string, entry []byte)
	Delete(key string)
	Reset()
	Len() int64
	Close()
	Has(key string) bool
}

type cache struct {
	engine *bigcache.BigCache
	log    *slog.Logger
}

func NewCache(segmentSize, maxCacheSize int, debug bool, log *slog.Logger) (Cache, error) {
	// Download cache
	config := bigcache.Config{
		// number of shards (must be a power of 2)
		Shards: 4,

		// time after which entry can be evicted
		LifeWindow: 5 * time.Minute,

		// Interval between removing expired entries (clean up).
		// If set to <= 0 then no action is performed.
		// Setting to < 1 second is counterproductive — bigcache has a one second resolution.
		CleanWindow: 5 * time.Minute,

		// max entry size in bytes, used only in initial memory allocation
		MaxEntrySize: segmentSize,

		// prints information about additional memory allocation
		Verbose: debug,

		// cache will not allocate more memory than this limit, value in MB
		// if value is reached then the oldest entries can be overridden for the new ones
		// 0 value means no size limit
		HardMaxCacheSize: maxCacheSize / 3,
	}

	engine, err := bigcache.New(context.Background(), config)
	if err != nil {
		return nil, err
	}

	return &cache{
		engine: engine,
		log:    log,
	}, nil
}

func (c *cache) Get(key string) []byte {
	if e, err := c.engine.Get(key); err == nil {
		return e
	}

	return nil
}

func (c *cache) Set(key string, entry []byte) {
	e := make([]byte, len(entry))
	copy(e, entry)

	err := c.engine.Set(key, e)
	if err != nil {
		c.log.Error("Error setting cache entry", "key", key, "err", err)
	}
}

func (c *cache) Delete(key string) {
	err := c.engine.Delete(key)
	if err != nil {
		c.log.Error("Error deleting cache entry", "key", key, "err", err)
	}
}

func (c *cache) Reset() {
	err := c.engine.Reset()
	if err != nil {
		c.log.Error("Error resetting cache", "err", err)
	}
}

func (c *cache) Len() int64 {
	return int64(c.engine.Len())
}

func (c *cache) Close() {
	err := c.engine.Close()
	if err != nil {
		c.log.Error("Error closing cache", "err", err)
	}
}

func (c *cache) Has(key string) bool {
	_, err := c.engine.Get(key)
	return err == nil
}
