package cache

type Cache interface {
	Get(key string) any
	Set(key string, value int)
}

type cacheLRU struct {
	l []map[string]any
	c int
}

func NewCacheLRU(capacity int) Cache {
	list := make([]map[string]any, 0, capacity)

	return &cacheLRU{
		l: list,
		c: capacity,
	}
}

func (lru *cacheLRU) Get(key string) any {

	for _, v := range lru.l {
		_, ok := v[key]
		if !ok {
			continue
		}
		return v[key]
	}

	return -1
}

func (lru *cacheLRU) Set(key string, value int) {

	if len(lru.l)+1 > lru.c {
		lru.l = lru.l[1:]
	}
	lru.l = append(lru.l, map[string]any{key: value})
}
