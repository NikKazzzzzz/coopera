package base

import (
	"context"
	"fmt"

	"github.com/NikKazzzzzz/coopera-bot/pkg/botlib/content"
	"github.com/NikKazzzzzz/coopera-bot/pkg/botlib/core"
	"github.com/NikKazzzzzz/coopera-bot/pkg/botlib/tg"
	"github.com/NikKazzzzzz/coopera-bot/pkg/botlib/updates/attrs"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type sendContentAction struct {
	bot  tg.Bot
	view content.View
}

func (s sendContentAction) Perform(ctx context.Context, update telegram.Update) error {
	chatID, exists := attrs.ChatID(update).Value()
	if !exists {
		return fmt.Errorf("getting chat ID from update: chat ID not found")
	}
	content, err := s.view.Render(ctx, update)
	if err != nil {
		return fmt.Errorf("rendering content for update: %w", err)
	}
	return s.bot.Chat(chatID).Send(ctx, content)
}

func SendContent(bot tg.Bot, view content.View) core.Action {
	return sendContentAction{
		bot:  bot,
		view: view,
	}
}
