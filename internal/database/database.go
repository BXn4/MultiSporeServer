package database

import (
	"fmt"

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
		"PRAGMA journal_mode=WAL",    // Allows concurrent reads during writes
		"PRAGMA busy_timeout=5000",   // Prevents immediate "locked" errors
		"PRAGMA synchronous=NORMAL",  // Good for WAL mode (faster than FULL)
		"PRAGMA cache_size=-128000",  // 128MB cache (negative = KB)
		"PRAGMA foreign_keys=ON",     // Good practice for data integrity
		"PRAGMA temp_store=memory",   // Faster temporary operations
		"PRAGMA mmap_size=268435456", // 256MB memory-mapped I/O (faster reads)
	}
	for _, pragma := range pragmas {
		if _, err := sqlDB.Exec(pragma); err != nil {
			return nil, fmt.Errorf("failed to set pragma: %w", err)
		}
	}

	/*err = db.AutoMigrate(&player.Player{}, &cafe.Cafe{}, &coops.Coop{})
	if err != nil {
		return nil, err
		}*/

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
