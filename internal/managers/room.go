package managers

import (
	"fmt"
	"multispore/internal/database"
	"multispore/internal/types"

	"github.com/charmbracelet/log"
)

func (gm *GameManager) SetDB(db *database.Database) {
	gm.db = db
}

func (gm *GameManager) SetRoom(id int, room *LoadedRoom) {
	gm.roomMutex.Lock()
	defer gm.roomMutex.Unlock()

	gm.rooms[id] = room
}

func (gm *GameManager) RemoveRoom(id int) {
	gm.roomMutex.Lock()
	defer gm.roomMutex.Unlock()
}

func (gm *GameManager) AddRoom(stage types.Stages) (*LoadedRoom, error) {
	gm.roomMutex.Lock()
	defer gm.roomMutex.Unlock()

	item, err := gm.getRoomByStage(stage)
	if err == nil {
		return item, nil
	}

	id := len(gm.rooms)
	loc := &LoadedRoom{
		stage: stage,
	}
	gm.rooms[id] = loc

	log.Infof("Added Room: [%d] stage: %d", id, stage)

	return loc, nil
}

func (gm *GameManager) getRoomByStage(stage types.Stages) (*LoadedRoom, error) {
	gm.roomMutex.Lock()
	defer gm.roomMutex.Unlock()

	for _, loc := range gm.rooms {
		if loc.stage == stage {
			return loc, nil
		}
	}

	return nil, fmt.Errorf("Room for stage %d not found", stage)
}
