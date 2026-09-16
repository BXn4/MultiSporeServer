package managers

import (
	"multispore/internal/types"
	"multispore/internal/types/response"
	"sync"

	"github.com/charmbracelet/log"
)

type LoadedRoom struct {
	occupants map[int](chan<- response.Response)
	mu        sync.Mutex
	gm        *GameManager
	running   bool
	stage     types.Stages
}

func NewLoadedRoom(gm *GameManager, sg types.Stages) *LoadedRoom {
	return &LoadedRoom{
		gm:        gm,
		occupants: make(map[int](chan<- response.Response), 0),
		running:   false,
		stage:     sg,
	}
}

func (lr *LoadedRoom) Send(id int, args ...string) {
	lr.mu.Lock()
	defer lr.mu.Unlock()

	lr.send(id, args...)
}

func (lr *LoadedRoom) Broadcast(args ...string) {
	lr.mu.Lock()
	defer lr.mu.Unlock()

	lr.broadcast(args...)
}

func (lr *LoadedRoom) Announce(playerID int, args ...string) {
	lr.mu.Lock()
	defer lr.mu.Unlock()

	lr.announce(playerID, args...)
}

func (lr *LoadedRoom) IsEmpty() bool {
	lr.mu.Lock()
	defer lr.mu.Unlock()
	return len(lr.occupants) == 0
}

func (lr *LoadedRoom) InRoom(id int) bool {
	lr.mu.Lock()
	defer lr.mu.Unlock()
	_, ok := lr.occupants[id]
	return ok
}

func (lr *LoadedRoom) IsRunning() bool {
	lr.mu.Lock()
	defer lr.mu.Unlock()
	return lr.running
}

func (lr *LoadedRoom) SetRunning(b bool) {
	lr.mu.Lock()
	defer lr.mu.Unlock()

	lr.running = b
}

func (lr *LoadedRoom) Join(playerID int, channel chan<- response.Response) {
	lr.mu.Lock()
	defer lr.mu.Unlock()
	// TODO
	lr.occupants[playerID] = channel
}

func (lr *LoadedRoom) Leave(id int) {
	lr.mu.Lock()
	defer lr.mu.Unlock()
}

func (lr *LoadedRoom) send(id int, args ...string) {
	resp := response.NewExtensionResponse(args...)
	log.Logf(log.Level(-3), "to %v: %s", id, resp.Wrap())
	lr.occupants[id] <- resp
}

func (lr *LoadedRoom) broadcast(args ...string) {
	resp := response.NewExtensionResponse(args...)
	log.Logf(log.Level(-1), "%s", resp.Wrap())
	log.Debugf("Broadcasted for: %d clients", len(lr.occupants))
	for _, channel := range lr.occupants {
		channel <- resp
	}
}

func (lr *LoadedRoom) announce(playerID int, args ...string) {
	resp := response.NewExtensionResponse(args...)
	log.Logf(log.Level(-2), "%s", resp.Wrap())
	for id, channel := range lr.occupants {
		if id == playerID {
			continue
		}
		channel <- resp
	}
}
