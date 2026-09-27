package database

import (
	"fmt"
	"multispore/internal/models/creature"
	"multispore/internal/models/player"

	"gorm.io/gorm"
)

func (db *Database) GetPlayerByName(name string) (*player.Player, error) {
	var p player.Player
	var creature creature.Creature
	err := db.conn.Where("username = ?", name).First(&p).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("NAME: %v NOT FOUND", name)
		}
		return nil, fmt.Errorf("SQL ERR: %v", err)
	}

	if err := db.conn.First(&creature, p.CreatureID).Error; err != nil {
		return nil, fmt.Errorf("Cant load creature %d for player %d: %w", p.CreatureID, p.ID, err)
	}
	p.SetCreature(creature)

	return &p, nil

}
