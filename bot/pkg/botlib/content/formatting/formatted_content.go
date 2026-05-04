package formatting

import (
	"github.com/NikKazzzzzz/coopera-bot/pkg/botlib/content"
	"github.com/NikKazzzzzz/coopera-bot/pkg/repr"
	"github.com/NikKazzzzzz/coopera-bot/pkg/repr/json"
)

type ParseMode string

const (
	ParseModeHTML       ParseMode = "HTML"
	ParseModeMarkdown   ParseMode = "Markdown"
	ParseModeMarkdownV2 ParseMode = "MarkdownV2"
)

type formattedContent struct {
	origin    content.Content
	parseMode ParseMode
}

func (f formattedContent) Structure() repr.Structure {
	return repr.Must(f.origin.Structure().Extend(
		repr.EmptyPath(),
		json.Object(json.Fields{
			"parse_mode": json.Str(string(f.parseMode)),
		}),
	))
}

func (f formattedContent) Method() string {
	return f.origin.Method()
}

func Formatted(content content.Content, parseMode ParseMode) content.Content {
	return formattedContent{
		origin:    content,
		parseMode: parseMode,
	}
}
