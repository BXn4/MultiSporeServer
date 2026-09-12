package managers

import (
	"fmt"

	"multispore/internal/client"
	"multispore/internal/interfaces"

	"github.com/charmbracelet/log"
)

func (gm *GameManager) AddClient(item interfaces.ManagedItem) {
	gm.clientMutex.Lock()
	defer gm.clientMutex.Unlock()

	c := item.(*client.Client)

	clientID := gm.NextClientID()
	c.SetClientID(clientID)

	gm.clients[clientID] = c

	log.Info("Assigned a ID to client: ", clientID)
}

func (gm *GameManager) NextClientID() int {
	for i := 1; ; i++ {
		if _, exists := gm.clients[i]; !exists {
			return i
		}
	}
}

func (gm *GameManager) GetClient(id int) (interfaces.ManagedItem, error) {
	gm.clientMutex.Lock()
	defer gm.clientMutex.Unlock()

	for _, c := range gm.clients {
		/* if c.Player == nil {
		continue
		} */
		//if c.Player.GetID() == id {
		return c, nil
		//}
	}

	return nil, fmt.Errorf("Client with ID %v not found", id)
}

func (gm *GameManager) DisconnectClient(id int) {
	gm.clientMutex.Lock()
	defer gm.clientMutex.Unlock()

	_, ok := gm.clients[id]
	if !ok {
		log.Info("[Disconnect] Client ID not found:", id)
		return
	}

	delete(gm.clients, id)

	log.Info("[Disconnect] Disconnecting client id:", id)
	log.Info("[Disconnect] Client removed, remaining clients:", len(gm.clients))
}
