package generator

import (
	"fmt"

	"github.com/aidarkhanov/nanoid/v2"
)

func GenerateGroupID() (id string, err error) {
	id, err = generate(3)
	id = fmt.Sprintf("g-%s", id)
	return
}

func GenerateAdminID(id string, err error) {
	id, err = generate(2)
	id = fmt.Sprintf("u-%s", id)
	return
}

func GenerateAddressID() (id string, err error) {
	id, err = generate(3)
	id = fmt.Sprintf("a-%s", id)
	return
}

func GeneratePropertyID() (id string, err error) {
	id, err = generate(7)
	id = fmt.Sprintf("p-%s", id)
	return
}

func generate(size int) (id string, err error) {
	id, err = nanoid.GenerateString(nanoid.DefaultAlphabet, size)
	return
}
