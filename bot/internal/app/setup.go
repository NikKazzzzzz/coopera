package app

import (
	nethttp "net/http"

	"github.com/NikKazzzzzz/coopera-bot/internal/domain"
	"github.com/NikKazzzzzz/coopera-bot/internal/domain/http"
	"github.com/NikKazzzzzz/coopera-bot/internal/domain/transport"
	"github.com/NikKazzzzzz/coopera-bot/pkg/botlib/core"
	"github.com/NikKazzzzzz/coopera-bot/pkg/botlib/engine"
	"github.com/NikKazzzzzz/coopera-bot/pkg/botlib/forms"
	formskeyvalue "github.com/NikKazzzzzz/coopera-bot/pkg/botlib/forms/keyvalue"
	"github.com/NikKazzzzzz/coopera-bot/pkg/botlib/keyvalue"
	"github.com/NikKazzzzzz/coopera-bot/pkg/botlib/sessions"
	sessionskeyvalue "github.com/NikKazzzzzz/coopera-bot/pkg/botlib/sessions/keyvalue"
	"github.com/NikKazzzzzz/coopera-bot/pkg/botlib/tg"
	tgtransport "github.com/NikKazzzzzz/coopera-bot/pkg/botlib/tg/transport"
)

func Store() keyvalue.Store {
	return keyvalue.MemoryStore()
}

func Sessions(store keyvalue.Store) sessions.Sessions {
	return sessionskeyvalue.Sessions(store)
}

func Forms(store keyvalue.Store) forms.Forms {
	return formskeyvalue.KeyValueForms(store)
}

func Bot(token string) tg.Bot {
	return tg.NewBot(tgtransport.HTTPClient(token))
}

func Community(s string) domain.Community {
	return http.Community(transport.HTTPClient(s, &nethttp.Client{}))
}

func Engine(token string, action core.Action) engine.Engine {
	return engine.ShutdownEngine(
		engine.SingleWorkerEngine(
			token, action,
		),
	)
}
