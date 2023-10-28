package cache

import "errors"

type Cache interface {
	get(key string) (string, error)
	set(key, value string)
}

type SimpleCache struct {
	storage map[string]string
}

func (sc SimpleCache) get(key string) (value string, err error) {
	val, ok := sc.storage[key]
	if !ok {
		err = errors.New("value is absent")
	}
	return val, err
}

func (sc SimpleCache) set(key string, value string) {
	sc.storage[key] = value
}

func GetSimpleCacheInstance() Cache {
	return SimpleCache{
		storage: make(map[string]string, 0),
	}
}
