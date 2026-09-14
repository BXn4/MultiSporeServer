package managers

import (
	"multispore/internal/types/response"
	"sync"

	"github.com/charmbracelet/log"
)

type LoadedLocation struct {
	occupants map[int](chan<- response.Response)
	mu        sync.Mutex
	gm        *GameManager
	running   bool
}

func NewLoadedLocation(gm *GameManager) *LoadedLocation {
	return &LoadedLocation{
		gm:        gm,
		occupants: make(map[int](chan<- response.Response), 0),
		running:   false,
	}
}

func (lc *LoadedLocation) Send(id int, args ...string) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	lc.send(id, args...)
}

func (lc *LoadedLocation) Broadcast(args ...string) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	lc.broadcast(args...)
}

func (lc *LoadedLocation) Announce(playerID int, args ...string) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	lc.announce(playerID, args...)
}

func (lc *LoadedLocation) IsEmpty() bool {
	lc.mu.Lock()
	defer lc.mu.Unlock()
	return len(lc.occupants) == 0
}

func (lc *LoadedLocation) AtLocation(id int) bool {
	lc.mu.Lock()
	defer lc.mu.Unlock()
	_, ok := lc.occupants[id]
	return ok
}

func (lc *LoadedLocation) IsRunning() bool {
	lc.mu.Lock()
	defer lc.mu.Unlock()
	return lc.running
}

func (lc *LoadedLocation) SetRunning(b bool) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	lc.running = b
}

func (lc *LoadedLocation) Join(playerID int, channel chan<- response.Response) {
	lc.mu.Lock()
	defer lc.mu.Unlock()
	// TODO
	lc.occupants[playerID] = channel
}

func (lc *LoadedLocation) Leave(id int) {
	lc.mu.Lock()
	defer lc.mu.Unlock()
}

func (lc *LoadedLocation) send(id int, args ...string) {
	resp := response.NewExtensionResponse(args...)
	log.Logf(log.Level(-3), "to %v: %s", id, resp.Wrap())
	lc.occupants[id] <- resp
}

func (lc *LoadedLocation) broadcast(args ...string) {
	resp := response.NewExtensionResponse(args...)
	log.Logf(log.Level(-1), "%s", resp.Wrap())
	log.Debugf("Broadcasted for: %d clients", len(lc.occupants))
	for _, channel := range lc.occupants {
		channel <- resp
	}
}

func (lc *LoadedLocation) announce(playerID int, args ...string) {
	resp := response.NewExtensionResponse(args...)
	log.Logf(log.Level(-2), "%s", resp.Wrap())
	for id, channel := range lc.occupants {
		if id == playerID {
			continue
		}
		channel <- resp
	}
}
