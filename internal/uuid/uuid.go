package uuid

import "github.com/google/uuid"

type UUIDGenerate interface {
	NewString() string
}
type UUIDGenerator struct {
	UUID UUIDGenerate
}

func (g *UUIDGenerator) NewStringSlice(count int) []string {
	uuids := []string{}

	for i := 0; i <= count; i++ {
		uuids = append(uuids, g.UUID.NewString())
	}

	return uuids
}

type GoogleUUID struct {
}

func (g *GoogleUUID) NewString() string {
	return uuid.NewString()
}
