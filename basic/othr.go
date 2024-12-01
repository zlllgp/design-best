package basic

import "sync"

const (
	_defaultPort = 8000
)

type SMap struct {
	sync.Mutex

	data map[string]string
}

func newSMap() *SMap {
	return &SMap{
		data: make(map[string]string),
	}
}

func (m *SMap) Get(k string) string {
	m.Lock()
	defer m.Unlock()

	return m.data[k]
}
