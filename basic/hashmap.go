package basic

import (
	"context"
	"sync"
	"time"
)

type MyChan struct {
	sync.Once
	ch chan struct{}
}

func NewMyChan() *MyChan {
	return &MyChan{
		ch: make(chan struct{}),
	}
}
func (m *MyChan) Close() {
	m.Do(func() {
		close(m.ch)
	})
}

type MyHashMap struct {
	sync.Mutex
	mp map[int]int
	ch map[int]*MyChan
}

func NewHashMap() *MyHashMap {
	return &MyHashMap{
		mp: make(map[int]int),
	}
}

func (m *MyHashMap) Put(key int, value int) {
	m.Lock()
	defer m.Unlock()
	m.mp[key] = value
	ch, ok := m.ch[key]
	if !ok {
		return
	}
	ch.Close()
}

func (m *MyHashMap) Get(key int, maxWaitDuration time.Duration) (int, error) {
	m.Lock()
	value, ok := m.mp[key]
	if ok {
		m.Unlock()
		return value, nil
	}
	ch, ok := m.ch[key]
	if !ok {
		ch = NewMyChan()
		m.ch[key] = ch
	}
	tCtx, cancel := context.WithTimeout(context.Background(), maxWaitDuration)
	defer cancel()
	m.Unlock()

	select {
	case <-tCtx.Done():
		return -1, tCtx.Err()
	case <-ch.ch:
	}

	m.Lock()
	value, ok = m.mp[key]
	m.Unlock()
	return value, nil
}
