package managers

import (
	"multispore/internal/client"
	"multispore/internal/database"
	"sync"
)

type GameManager struct {
	db *database.Database

	roomMutex sync.Mutex
	rooms     map[int]*LoadedRoom

	clientMutex sync.Mutex
	clients     map[int]*client.Client
}

func NewGameManager() (*GameManager, error) {
	gm := &GameManager{
		rooms:   make(map[int]*LoadedRoom, 0),
		clients: make(map[int]*client.Client, 0),
	}

	return gm, nil
}

func (gm *GameManager) GetClients() map[int]*client.Client {
	gm.clientMutex.Lock()
	defer gm.clientMutex.Unlock()

	return gm.clients
}
