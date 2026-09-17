package database

import (
	"errors"
	"fmt"
	"multispore/internal/models/player"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (db *Database) Authenticate(name string, pass string) (*player.Player, int, error) {
	var p player.Player
	err := db.conn.Where("username = ?", name).First(&p).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 1, fmt.Errorf("Player \"%v\" not found!", name)
		}
		return nil, 1, fmt.Errorf("DB Error: %v", err)
	}

	if !VerifyPassword(p.Password, pass) {
		return nil, 1, errors.New("Access Denied!")
	}

	/*if p.GetIsBanned() {
	return nil, 2, errors.New("You are banned!")
	}*/

	return &p, 0, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func VerifyPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
