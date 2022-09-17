package mysql

import (
	"github.com/viktigpetterr/game-rest-api/internal/game/adapter/mysql/model"
)

func Migrate() error {
	db, err := Connection()
	if err != nil {
		return err
	}

	return db.AutoMigrate(model.User{}, model.GameState{})
}
