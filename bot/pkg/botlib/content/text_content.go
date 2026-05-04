package content

import (
	"github.com/NikKazzzzzz/coopera-bot/pkg/repr"
	"github.com/NikKazzzzzz/coopera-bot/pkg/repr/json"
)

type textContent struct {
	text string
}

func (t textContent) Structure() repr.Structure {
	return json.Object(json.Fields{
		"text": json.Str(t.text),
	})
}

func (t textContent) Method() string {
	return "sendMessage"
}

func Text(text string) Content {
	return textContent{text: text}
}
