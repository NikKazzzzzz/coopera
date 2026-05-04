package keyboards

import "github.com/NikKazzzzzz/coopera-bot/pkg/botlib/content"

type Keyboard interface {
	content.Content
}

type InlineKeyboard interface {
	Keyboard
}

type ReplyKeyboard interface {
	Keyboard
}
