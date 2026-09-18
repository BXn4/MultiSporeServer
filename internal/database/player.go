package database

import (
	"fmt"
	"multispore/internal/models/player"

	"gorm.io/gorm"
)

func (db *Database) GetPlayerByName(name string) (*player.Player, error) {
	var p player.Player
	err := db.conn.Where("username = ?", name).First(&p).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("NAME: %v NOT FOUND", name)
		}
		return nil, fmt.Errorf("SQL ERR: %v", err)
	}

	return &p, nil

}
