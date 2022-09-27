package testdata

import (
	"github.com/joho/godotenv"
	"github.com/viktigpetterr/game-rest-api/internal/game/adapter/mysql"
	"gorm.io/gorm"
)

var ConnectionErr error

type Connection struct {
	db *gorm.DB
}

func (c *Connection) Open() (*gorm.DB, error) {
	if ConnectionErr != nil {
		return nil, ConnectionErr
	}
	if c.db != nil {
		return c.db, nil
	}
	err := godotenv.Load("../../../../.env.local")
	if err != nil {
		return nil, err
	}
	connection := mysql.Connection{}
	err = connection.Migrate()
	if err != nil {
		return nil, err
	}
	db, err := connection.Open()
	if err != nil {
		return nil, err
	}
	c.db = db
	return c.db, nil
}

func Reset() {
	ConnectionErr = nil
}
