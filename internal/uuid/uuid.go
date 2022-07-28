package uuid

import "github.com/google/uuid"

type UUID interface {
	NewString() string
	NewStringSlice(count int) []string
}

type UUIDImpl struct {
}

func New() UUID {
	return &UUIDImpl{}
}

func (ui *UUIDImpl) NewString() string {
	return uuid.NewString()
}

func (ui *UUIDImpl) NewStringSlice(count int) []string {
	uuids := []string{}

	for i := 0; i <= count; i++ {
		uuids = append(uuids, uuid.NewString())
	}

	return uuids
}
