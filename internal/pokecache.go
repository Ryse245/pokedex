package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheMap map[string]cacheEntry
	interval time.Duration
	mux      *sync.Mutex
	pokeBox  map[string]Pokemon
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration, mux *sync.Mutex) Cache {
	ret := Cache{map[string]cacheEntry{}, interval, mux, map[string]Pokemon{}}
	ticker := time.NewTicker(interval)
	go func() {
		for range ticker.C {
			ret.reapLoop()
		}
	}()
	return ret
}

func (c Cache) Add(key string, value []byte) {
	c.mux.Lock()
	c.cacheMap[key] = cacheEntry{time.Now(), value}
	c.mux.Unlock()
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	data, found := c.cacheMap[key]
	if !found {
		return nil, false
	}
	return data.val, true
}

func (c Cache) AddPokemon(key string, val Pokemon) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.pokeBox[key] = val
}

func (c Cache) GetPokemon(key string) (Pokemon, bool) {
	data, exist := c.pokeBox[key]
	if exist {
		return data, true
	} else {
		return Pokemon{}, false
	}
}

func (c Cache) reapLoop() {
	c.mux.Lock()
	defer c.mux.Unlock()
	for key, val := range c.cacheMap {
		if time.Since(val.createdAt) > c.interval {
			delete(c.cacheMap, key)
		}
	}
}
