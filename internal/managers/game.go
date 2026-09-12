package managers

import (
	"multispore/internal/client"
	"sync"
)

type GameManager struct {
	locationMutex sync.Mutex

	clientMutex sync.Mutex
	clients     map[int]*client.Client
}

func NewGameManager() (*GameManager, error) {
	gm := &GameManager{
		clients: make(map[int]*client.Client, 0),
	}

	return gm, nil
}

func (gm *GameManager) GetClients() map[int]*client.Client {
	gm.clientMutex.Lock()
	defer gm.clientMutex.Unlock()

	return gm.clients
}
