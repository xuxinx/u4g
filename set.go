package u4g

import (
	"sync"
)

type Set interface {
	Add(items ...interface{})
	Delete(items ...interface{})
	Clear()
	Has(interface{}) bool
	Items() []interface{}
	Len() int
}

func NewSet() Set {
	return &setImpl{
		items: make(map[interface{}]struct{}),
	}
}

type setImpl struct {
	l     int
	items map[interface{}]struct{}

	m sync.RWMutex
}

func (si *setImpl) Add(items ...interface{}) {
	si.m.Lock()
	defer si.m.Unlock()
	for _, item := range items {
		if _, ok := si.items[item]; !ok {
			si.items[item] = struct{}{}
			si.l++
		}
	}
}

func (si *setImpl) Delete(items ...interface{}) {
	si.m.Lock()
	defer si.m.Unlock()
	for _, item := range items {
		if _, ok := si.items[item]; ok {
			delete(si.items, item)
			si.l--
		}
	}
}

func (si *setImpl) Clear() {
	si.m.Lock()
	defer si.m.Unlock()
	si.items = make(map[interface{}]struct{})
	si.l = 0
}

func (si *setImpl) Has(item interface{}) bool {
	si.m.RLock()
	defer si.m.RUnlock()
	_, ok := si.items[item]

	return ok
}

func (si *setImpl) Items() []interface{} {
	si.m.RLock()
	defer si.m.RUnlock()
	if si.l == 0 {
		return nil
	}

	items := make([]interface{}, 0, si.l)
	for item, _ := range si.items {
		items = append(items, item)
	}

	return items
}

func (si *setImpl) Len() int {
	si.m.RLock()
	defer si.m.RUnlock()
	return si.l
}
