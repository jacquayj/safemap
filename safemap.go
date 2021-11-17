package safemap

import (
	"sync"
)

type Map[K comparable, V any] struct {
	mapData map[K]V
	mutex   sync.Mutex
}

func NewMap[K comparable, V any](startMap map[K]V) *Map[K, V] {
	sm := new(Map[K, V])
	sm.mapData = startMap
	return sm
}

func (sm *Map[K, V]) Delete(key K) bool {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()
	_, ok := sm.mapData[key]
	if ok {
		delete(sm.mapData, key)
	}
	return ok
}

func (sm *Map[K, V]) Get(key K) (V, bool) {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()
	val, ok := sm.mapData[key]
	return val, ok
}

func (sm *Map[K, V]) Set(key K, val V) {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()
	sm.mapData[key] = val
}
