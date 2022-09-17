package service

import uuid "github.com/satori/go.uuid"

type Uuid struct{}

func (_ Uuid) New() string {
	return uuid.NewV4().String()
}
