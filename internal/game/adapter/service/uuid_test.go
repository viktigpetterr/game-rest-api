package service

import (
	"testing"
)

func TestNew(t *testing.T) {
	s := Uuid{}
	uuid := s.New()
	if len(uuid) == 0 {
		t.Error("failed to assert that uuid was not empty")
	}
}
