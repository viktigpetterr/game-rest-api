package service

type UuidMock struct{}

func (_ UuidMock) New() string {
	return "uuid"
}
