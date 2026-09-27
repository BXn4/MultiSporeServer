package creature

type Creature struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement"`   // creatureID
	PlayerID  int       `gorm:"column:player_id;type:int;not null"`   // creatorPlayerID
	Name      string    `gorm:"column:name;type:string;not null"`     // creatureName
	Rigblocks Rigblocks `gorm:"column:rigblock;type:string;not null"` // rigblocks
	Health    float32   `gorm:"column:health;type:int;not null"`      // health for the creature
	MaxHealth float32   `gorm:"column:max_health;type:int;not null"`  // max health, will be calculated from the parts
	Hunger    float32   `gorm:"column:hunger;type:int;not null"`      // hunger for the creature
}

func (Creature) TableName() string {
	return "creatures"
}

func NewCreatureFromString(s string) *Creature {
	rigblocks := NewRigblockFromString(s)
	if rigblocks == nil {
		return nil
	}

	return &Creature{
		Rigblocks: *rigblocks,
	}
}
