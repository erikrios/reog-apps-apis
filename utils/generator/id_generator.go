package generator

import (
	"fmt"

	"github.com/aidarkhanov/nanoid/v2"
)

type IDGenerator interface {
	GenerateGroupID() (id string, err error)
	GenerateAdminID(id string, err error)
	GeneratePropertyID() (id string, err error)
}

type nanoidIDGenerator struct {
}

func NewNanoidIDGenerator() *nanoidIDGenerator {
	return &nanoidIDGenerator{}
}

func (n *nanoidIDGenerator) GenerateGroupID() (id string, err error) {
	id, err = n.generate(3)
	id = fmt.Sprintf("g-%s", id)
	return
}

func (n *nanoidIDGenerator) GenerateAdminID() (id string, err error) {
	id, err = n.generate(2)
	id = fmt.Sprintf("a-%s", id)
	return
}

func (n *nanoidIDGenerator) GeneratePropertyID() (id string, err error) {
	id, err = n.generate(7)
	id = fmt.Sprintf("p-%s", id)
	return
}

func (n *nanoidIDGenerator) generate(size int) (id string, err error) {
	id, err = nanoid.GenerateString(nanoid.DefaultAlphabet, size)
	return
}
