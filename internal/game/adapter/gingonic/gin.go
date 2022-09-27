package gingonic

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/viktigpetterr/game-rest-api/internal/game/adapter/gingonic/user"
	"github.com/viktigpetterr/game-rest-api/internal/game/adapter/mysql"
	"log"
	"os"
)

func Run() {
	r, err := Init()
	if err != nil {
		log.Fatalln(err)
	}
	err = r.Run(":" + os.Getenv("PORT"))
	if err != nil {
		log.Fatalln(err)
	}
}

func Init() (*gin.Engine, error) {
	err := SetUpEnv()
	if err != nil {
		return nil, err
	}
	r := Router()
	_ = r.SetTrustedProxies([]string{os.Getenv("GIN_TRUSTED_PROXY")})
	err = mysql.Connection{}.Migrate()

	return r, err
}

func Router() *gin.Engine {
	r := gin.Default()
	user.Router(r)

	return r
}

func SetUpEnv() error {
	mode := gin.Mode()
	var err error
	if mode == "debug" {
		err = godotenv.Load(".env.local")
	} else {
		err = godotenv.Load(".env")
	}

	return err
}
