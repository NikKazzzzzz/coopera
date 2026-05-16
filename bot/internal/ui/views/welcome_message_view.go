package views

import (
	"context"
	"fmt"
	"net/url"

	"github.com/NikKazzzzzz/coopera-bot/pkg/botlib/content"
	"github.com/NikKazzzzzz/coopera-bot/pkg/botlib/tg"
	"github.com/NikKazzzzzz/coopera-bot/pkg/botlib/content/keyboards"
	"github.com/NikKazzzzzz/coopera-bot/pkg/botlib/content/keyboards/buttons"
	"github.com/NikKazzzzzz/coopera-bot/pkg/botlib/updates/attrs"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type welcomeMessageView struct {
	frontendURL string
	bot         tg.Bot
}

func (w welcomeMessageView) Render(ctx context.Context, update telegram.Update) (content.Content, error) {
	chatID, _ := attrs.ChatID(update).Value()
	username, _ := attrs.Username(update).Value()

	authURL := fmt.Sprintf("%s/auth-callback?tg_id=%d&username=%s", w.frontendURL, chatID, url.QueryEscape(username))

	if photoURL, err := w.bot.UserPhotoURL(ctx, chatID); err == nil && photoURL != "" {
		authURL += "&photo_url=" + url.QueryEscape(photoURL)
	}

	return keyboards.Inline(
		content.Text("Добро пожаловать в Coopera!\n\nНажмите кнопку ниже, чтобы войти на сайт."),
		buttons.Matrix(
			buttons.Row(buttons.URLButton("Войти на сайт", authURL)),
		),
	), nil
}

func WelcomeMessage(frontendURL string, bot tg.Bot) content.View {
	return welcomeMessageView{frontendURL: frontendURL, bot: bot}
}
