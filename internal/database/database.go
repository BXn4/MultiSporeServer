package database

import (
	"fmt"
	"multispore/internal/models/creature"
	"multispore/internal/models/player"
	"multispore/internal/types"

	"github.com/charmbracelet/log"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type DBConfig struct {
	Database string
}

type Database struct {
	conn *gorm.DB
}

func ConnectToDB(config *DBConfig) (*Database, error) {
	db, err := gorm.Open(sqlite.Open("database/"+config.Database+".db"), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(1)
	sqlDB.SetMaxIdleConns(1)

	pragmas := []string{
		"PRAGMA journal_mode=WAL",
		"PRAGMA busy_timeout=5000",
		"PRAGMA synchronous=NORMAL",
		"PRAGMA cache_size=-128000",
		"PRAGMA foreign_keys=ON",
		"PRAGMA temp_store=memory",
		"PRAGMA mmap_size=268435456",
	}
	for _, pragma := range pragmas {
		if _, err := sqlDB.Exec(pragma); err != nil {
			return nil, fmt.Errorf("failed to set pragma: %w", err)
		}
	}

	err = db.AutoMigrate(&player.Player{}, &creature.Creature{})
	if err != nil {
		return nil, err
	}

	database := &Database{conn: db}

	return database, nil
}

func (db *Database) Close() error {
	sqlDB, err := db.conn.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (db *Database) CreateAccount(name, password string, creature creature.Creature) (*player.Player, error) {
	hashedPasswd, err := HashPassword(password)
	if err != nil {
		return nil, err
	}

	player := &player.Player{
		Username: name,
		Password: hashedPasswd,
		Stage:    types.CELL,
	}

	err = db.conn.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(player).Error; err != nil {
			return fmt.Errorf("Cant create player: %w", err)
		}

		creature.PlayerID = player.GetID()
		if err := tx.Create(&creature).Error; err != nil {
			return fmt.Errorf("Cant create creature: %w", err)
		}

		if err := tx.Model(player).Update("creature_id", creature.ID).Error; err != nil {
			return fmt.Errorf("Cant set player creature_id: %w", err)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	log.Info("Registered!")

	return player, nil
}
