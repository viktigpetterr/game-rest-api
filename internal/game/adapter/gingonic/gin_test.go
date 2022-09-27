package gingonic

import (
	"github.com/gin-gonic/gin"
	"io/fs"
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	cwd, _ := os.Getwd()
	err := os.Chdir("../../../..")
	defer func(dir string) {
		err = os.Chdir(dir)
		if err != nil {
			t.Error("failed to reset working directory")
		}
	}(cwd)
	if err != nil {
		t.Error("failed to change directory")
	}
	r, err := Init()
	if err != nil {
		t.Error("failed to init application")
	}
	if r == nil {
		t.Error("failed to assert that engine was initialized")
	}
}

func TestRouter(t *testing.T) {
	r := Router()
	if r == nil {
		t.Error("failed to initiate router")
	}
}

func TestSetUpEnv(t *testing.T) {
	gin.SetMode("debug")
	err := SetUpEnv().(*fs.PathError)
	if err.Path != ".env.local" {
		t.Error("failed to assert that debug was set up")
	}
	gin.SetMode("release")
	err = SetUpEnv().(*fs.PathError)
	if err.Path != ".env" {
		t.Error("failed to assert that release env was set up")
	}
}
