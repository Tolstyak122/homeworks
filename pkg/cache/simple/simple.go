package simple

import (
	"errors"

	"github.com/Tolstyak122/homeworks/pkg/cache/contracts"
)

type simpleCache struct {
	storage map[string]string
}

func (sc simpleCache) Get(key string) (value string, err error) {
	val, ok := sc.storage[key]
	if !ok {
		err = errors.New("value is absent")
	}
	return val, err
}

func (sc simpleCache) Set(key string, value string) {
	sc.storage[key] = value
}

func GetCacheInstance() contracts.Cache {
	return simpleCache{
		storage: make(map[string]string, 0),
	}
}
