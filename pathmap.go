package ginprometheus

import (
	"fmt"
	"sync"
)

//KVStore can store key/values as strings
type KVStore interface {
	Get(s string) (string, error)
	Set(key, val string)
}

type pathMap struct {
	sync.RWMutex
	pathMap map[string]string
}

var errNotFound = fmt.Errorf("not found")

func newPathMap() *pathMap {
	return &pathMap{
		pathMap: make(map[string]string),
	}

}

func (pm *pathMap) Get(s string) (string, error) {
	pm.RLock()
	defer pm.RUnlock()
	if in, ok := pm.pathMap[s]; ok {
		return in, nil
	}

	return "", errNotFound
}

func (pm *pathMap) Set(key, val string) {
	pm.Lock()
	pm.pathMap[key] = val
	pm.Unlock()
}
