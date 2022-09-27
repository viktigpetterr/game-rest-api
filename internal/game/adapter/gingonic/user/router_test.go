package user

import (
	"github.com/gin-gonic/gin"
	"strings"
	"testing"
)

func TestRouter(t *testing.T) {
	engine := gin.Default()
	Router(engine)

	routes := engine.Routes()
	if len(routes) != 6 {
		t.Error("failed to assert the number of routes in the user group")
	}
	handler := "github.com/viktigpetterr/game-rest-api/internal/game/adapter/gingonic/user.Controller"
	for _, route := range routes {
		if !strings.Contains(route.Path, "/user") {
			t.Error("failed to assert user routing group")
		}
		if !strings.Contains(route.Handler, handler) {
			t.Error("failed to assert that the user controller was the handler")
		}
	}
}
