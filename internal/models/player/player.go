package player

import "multispore/internal/models/creature"

type Player struct {
	ID        int               `gorm:"column:id;primaryKey;autoIncrement;type:int"`
	DNAPoints int               `gorm:"column:dna_points;type:int;default:0"`
	Creature  creature.Creature `gorm:"column:creature;type:text"`
	Stage     int/*will be stages type*/ `gorm:"column:stage;type:int;default:0"`
}

func (player *Player) TableName() string {
	return "player"
}
