package content

import (
	"github.com/NikKazzzzzz/coopera-bot/pkg/repr"
)

type Content interface {
	Structure() repr.Structure
	Method() string
}
