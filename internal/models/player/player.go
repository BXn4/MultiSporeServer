package player

import (
	"multispore/internal/models/creature"
	"multispore/internal/types"
)

type Player struct {
	ID        int               `gorm:"column:id;primaryKey;autoIncrement;type:int"`
	Username  string            `gorm:"column:username;type:string;not null"`
	Password  string            `gorm:"column:password;type:string;not null"`
	DNAPoints int               `gorm:"column:dna_points;type:int;default:0"`
	Creature  creature.Creature `gorm:"column:creature;type:text"`
	Stage     types.Stages      `gorm:"column:stage;type:int;default:0"`
}

func (player *Player) TableName() string {
	return "player"
}
