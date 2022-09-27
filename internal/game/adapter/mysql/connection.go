package mysql

import (
	"github.com/viktigpetterr/game-rest-api/internal/game/adapter/mysql/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var config = mysql.Config{
	DSN:                       "",    // data source name
	DefaultStringSize:         256,   // default size for string fields
	DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
	DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
	DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
	SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
}

type IConnection interface {
	Open() (*gorm.DB, error)
}

type Connection struct{}

func (c Connection) Open() (*gorm.DB, error) {
	dsn := os.Getenv("MYSQL_DSN")
	config.DSN = dsn
	db, err := gorm.Open(
		mysql.New(config),
		&gorm.Config{},
	)
	return db, err
}

func (c Connection) Migrate() error {
	db, err := c.Open()
	if err != nil {
		return err
	}

	return db.AutoMigrate(model.User{}, model.GameState{})
}
