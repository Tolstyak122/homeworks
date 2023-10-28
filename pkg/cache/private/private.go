package private

import (
	"errors"

	"github.com/Tolstyak122/homeworks/pkg/cache/contracts"
)

var storage []map[string]string = []map[string]string{}

type privateCache struct {
	indx int
}

func (sc privateCache) Get(key string) (value string, err error) {
	val, ok := storage[sc.indx][key]
	if !ok {
		err = errors.New("value is absent")
	}
	return val, err
}

func (sc privateCache) Set(key string, value string) {
	storage[sc.indx][key] = value
}

func GetCacheInstance() contracts.Cache {
	indx := len(storage)
	storage = append(storage, map[string]string{})
	return privateCache{
		indx: indx,
	}
}
