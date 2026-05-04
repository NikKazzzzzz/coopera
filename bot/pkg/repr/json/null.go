package json

import (
	"github.com/NikKazzzzzz/coopera-bot/pkg/repr"
)

type null struct{}

func (n null) Marshal() ([]byte, error) {
	return []byte("null"), nil
}

func Null() repr.Primitive {
	return null{}
}
