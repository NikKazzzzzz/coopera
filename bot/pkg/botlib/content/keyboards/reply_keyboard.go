package keyboards

import (
	"github.com/NikKazzzzzz/coopera-bot/pkg/botlib/content"
	"github.com/NikKazzzzzz/coopera-bot/pkg/botlib/content/keyboards/buttons"
	"github.com/NikKazzzzzz/coopera-bot/pkg/repr"
	"github.com/NikKazzzzzz/coopera-bot/pkg/repr/json"
)

type replyKeyboard struct {
	origin  content.Content
	buttons buttons.ButtonMatrix[buttons.ReplyButton]
}

func (r replyKeyboard) Structure() repr.Structure {
	return repr.Must(r.origin.Structure().Extend(
		repr.EmptyPath(),
		json.Object(json.Fields{
			"reply_markup": json.Object(json.Fields{
				"keyboard": r.buttons.Structure(),
			}),
		}),
	))
}

func (r replyKeyboard) Method() string {
	return r.origin.Method()
}

func Reply(content content.Content, buttons buttons.ButtonMatrix[buttons.ReplyButton]) ReplyKeyboard {
	return replyKeyboard{
		origin:  content,
		buttons: buttons,
	}
}
