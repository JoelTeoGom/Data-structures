package SYNC_MAP

import (
	"data-structures/mapa"
	"sync"
)

type SyncMap struct {
	mu   sync.RWMutex
	data mapa.Map
}

func Build(size int) *SyncMap {
	return &SyncMap{
		data: *mapa.Build(size),
	}
}

func (sm *SyncMap) Set(key string, value any) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data.Set(key, value)
}

func (sm *SyncMap) Get(key string) (any, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	return sm.data.Get(key)
}
