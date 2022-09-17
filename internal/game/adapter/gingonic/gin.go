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
	SetUpEnv()
	r := Router()
	_ = r.SetTrustedProxies([]string{os.Getenv("GIN_TRUSTED_PROXY")})
	err := mysql.Migrate()
	if err != nil {
		log.Fatalln(err)
	}
	err = r.Run(":" + os.Getenv("PORT"))
	if err != nil {
		log.Fatalln(err)
	}
}

func Router() *gin.Engine {
	r := gin.Default()
	user.Router(r)

	return r
}

func SetUpEnv() {
	mode := gin.Mode()
	if mode == "debug" {
		_ = godotenv.Load(".env.local")
	} else {
		_ = godotenv.Load(".env")
	}
}
