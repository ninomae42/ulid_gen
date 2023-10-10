package id

import (
	"fmt"

	"github.com/oklog/ulid/v2"
)

type (
	ID string
)

func NewID(name string) ID {
	id := ulid.Make()
	if name != "" {
		return ID(fmt.Sprintf("%s+%s", name, id))
	}
	return ID(id.String())
}

func (i ID) String() string {
	return string(i)
}
