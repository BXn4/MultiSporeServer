package player

import (
	"multispore/internal/models/creature"
	"multispore/internal/types"
	"strings"
	"sync"
)

type Player struct {
	ID         int          `gorm:"column:id;primaryKey;autoIncrement;type:int"`
	Username   string       `gorm:"column:username;type:string;not null"`
	Password   string       `gorm:"column:password;type:string;not null"`
	DNAPoints  int          `gorm:"column:dna_points;type:int;default:0"`
	CreatureID int          `gorm:"column:creature_id;type:int;not null"`
	Stage      types.Stages `gorm:"column:stage;type:int;default:0"`

	creature creature.Creature `gorm:"-"`
	mutex    sync.Mutex        `gorm:"-"`
}

func (player *Player) TableName() string {
	return "player"
}

func (p *Player) String() string {
	params := []string{}
	return strings.Join(params, "+")
}

// ** SETTERS ** // ** SETTERS ** // ** SETTERS ** //
func (p *Player) SetID(v int) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.ID = v
}

func (p *Player) SetPassword(v string) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.Password = v
}

func (p *Player) SetUsername(v string) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.Username = v
}

func (p *Player) SetCreature(v creature.Creature) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.creature = v
	p.CreatureID = v.ID
}

func (p *Player) SetStage(v types.Stages) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.Stage = v
}

func (p *Player) SetDNA(v int) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.DNAPoints = v
}

func (p *Player) AddDNA(v int) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.DNAPoints += v
}

// ** GETTERS ** // ** GETTERS ** // ** GETTERS ** //
func (p *Player) GetID() int {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	return p.ID
}

func (p *Player) GetStage() types.Stages {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	return p.Stage
}

func (p *Player) GetCreature() creature.Creature {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	return p.creature
}

func (p *Player) GetCreatureID() int {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	return p.CreatureID
}
