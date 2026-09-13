package creature

type Creature struct {
	ID        int        `gorm:"column:id;type:int;not null"`
	PlayerID  int        `gorm:"column:player_id;type:int;not null"`
	Name      string     `gorm:"column:name;type:string;not null"`
	Rigblocks []Rigblock `gorm:"column:rigblocks;type:string;not null"`
	Health    float32    `gorm:"column:health;type:int;not null"`
	MaxHealth float32    `gorm:"column:max_healthtype:int;not null"`
}
