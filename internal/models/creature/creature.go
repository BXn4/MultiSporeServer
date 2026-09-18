package creature

type Creature struct {
	ID        int         `gorm:"column:id;type:int;not null"`
	PlayerID  int         `gorm:"column:player_id;type:int;not null"`
	Name      string      `gorm:"column:name;type:string;not null"`
	Rigblocks []Rigblocks `gorm:"column:rigblocks;type:string;not null"`
	Health    float32     `gorm:"column:health;type:int;not null"`
	MaxHealth float32     `gorm:"column:max_health;type:int;not null"`
	Hunger    float32     `gorm:"column:hunger;type:int;not null"`
}
