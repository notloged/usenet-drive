package filereader

//go:generate mockgen -source=./cache.go -destination=./cache_mock.go -package=filereader Cache
import (
	"github.com/dgraph-io/ristretto"
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
	engine      *ristretto.Cache
	segmentSize int
}

const OneMB = 1024 * 1024

func NewCache(segmentSize, maxCacheSize int, debug bool) (Cache, error) {
	numOfCounters := int64(((maxCacheSize * OneMB) / segmentSize) * 10)
	// Download cache
	engine, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: numOfCounters,              // number of keys to track frequency
		MaxCost:     int64(maxCacheSize) * 1e+6, // maximum cost of cache to bytes
		BufferItems: 64,                         // number of keys per Get buffer.
		Metrics:     debug,
	})
	if err != nil {
		return nil, err
	}

	return &cache{
		segmentSize: segmentSize,
		engine:      engine,
	}, nil
}

func (c *cache) Get(key string) []byte {
	value, found := c.engine.Get(key)
	if !found {
		return nil
	}

	return value.([]byte)
}

func (c *cache) Set(key string, entry []byte) {
	r := make([]byte, len(entry))
	copy(r, entry)

	c.engine.Set(key, r, int64(c.segmentSize))
	c.engine.Wait()
}

func (c *cache) Delete(key string) {
	c.engine.Del(key)
}

func (c *cache) Reset() {
	c.engine.Clear()
}

func (c *cache) Len() int64 {
	return c.engine.MaxCost()
}

func (c *cache) Close() {
	c.engine.Close()
}

func (c *cache) Has(key string) bool {
	_, found := c.engine.Get(key)
	return found
}
