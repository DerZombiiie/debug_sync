package sync

import (
	s "sync"

	"log"
	"path"
	"runtime"
)

var mutIdc int

type RWMutex struct {
	mu s.RWMutex

	id  int
	gid int // identifier of this mutex

	*log.Logger
	ionce s.Once
}

func (m *RWMutex) init() {
	if m.Logger == nil {
		m.Logger = log.Default()
	}

	m.ionce.Do(func() {
		m.gid = mutIdc
		mutIdc++
	})
}

func (m *RWMutex) RLock() {
	m.init()

	_, p, line, ok := runtime.Caller(1)
	if !ok {
		panic("cant get caller!")
	}

	_, name := path.Split(p)

	m.Printf("%s:%d -> %2d.RLock(%d) -> request\n", name, line, m.gid, m.id)
	m.mu.RLock()
	m.Printf("%s:%d -> %2d.RLock(%d) -> aquired\n", name, line, m.gid, m.id)

	m.id++
}

func (m *RWMutex) RUnlock() {
	m.init()

	_, p, line, ok := runtime.Caller(1)
	if !ok {
		panic("cant get caller!")
	}

	_, name := path.Split(p)

	m.Printf("%s:%d -> %2d.RUnlock(%d)\n", name, line, m.gid, m.id)
	m.mu.RUnlock()

	m.id++
}

func (m *RWMutex) Lock() {
	m.init()

	_, p, line, ok := runtime.Caller(1)
	if !ok {
		panic("cant get caller!")
	}

	_, name := path.Split(p)

	m.Printf("%s:%d -> %2d.Lock(%d) -> request\n", name, line, m.gid, m.id)
	m.mu.Lock()
	m.Printf("%s:%d -> %2d.Lock(%d) -> aquired\n", name, line, m.gid, m.id)
}

func (m *RWMutex) Unlock() {
	m.init()

	_, p, line, ok := runtime.Caller(1)
	if !ok {
		panic("cant get caller!")
	}

	_, name := path.Split(p)

	m.Printf("%s:%d -> %2d.Unlock(%d)\n", name, line, m.gid, m.id)
	m.mu.Unlock()
}
