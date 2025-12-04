package safemap

import (
	"sync"
)

type SafeMap struct {
	data map[int]int
	mx   sync.Mutex
}

func (sm *SafeMap) SetValue(key, value int) {
	sm.mx.Lock()
	defer sm.mx.Unlock()

	sm.data[key] = value
}

func (sm *SafeMap) GetValue(key int) (value int, ok bool) {
	sm.mx.Lock()
	defer sm.mx.Unlock()

	value, ok = sm.data[key]
	return
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		data: make(map[int]int),
	}
}
