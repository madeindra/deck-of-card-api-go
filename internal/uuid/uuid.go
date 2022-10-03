package uuid

import "github.com/google/uuid"

type UUIDGenerate interface {
	NewString() string
	NewStringSlice(count int) []string
}

type UUIDGenerator struct {
}

func New() UUIDGenerate {
	return &UUIDGenerator{}
}

func (ui *UUIDGenerator) NewString() string {
	return uuid.NewString()
}

func (ui *UUIDGenerator) NewStringSlice(count int) []string {
	uuids := []string{}

	for i := 0; i <= count; i++ {
		uuids = append(uuids, uuid.NewString())
	}

	return uuids
}
