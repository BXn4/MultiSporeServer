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

func (gm *GameManager) AddRoom(id int, stage types.Stages) (*LoadedRoom, error) {
	gm.roomMutex.Lock()
	defer gm.roomMutex.Unlock()

	loc, err := gm.getRoomByStage(stage)
	if err == nil {
		return loc, nil
	}

	gm.rooms[id] = loc
	loc.stage = stage

	log.Info("Added Room")

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
