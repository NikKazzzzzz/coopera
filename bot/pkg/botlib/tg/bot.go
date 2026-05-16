package tg

import (
	"context"
	"fmt"

	"github.com/NikKazzzzzz/coopera-bot/pkg/botlib/tg/transport"
	"github.com/tidwall/gjson"
)

type bot struct {
	dataSource transport.Client
	token      string
}

func (b bot) Chat(id int64) Chat {
	return NewChat(id, b.dataSource)
}

func (b bot) UserPhotoURL(ctx context.Context, userID int64) (string, error) {
	payload := fmt.Sprintf(`{"user_id":%d,"limit":1}`, userID)
	data, err := b.dataSource.Execute(ctx, "getUserProfilePhotos", []byte(payload))
	if err != nil {
		return "", err
	}
	fileID := gjson.GetBytes(data, "result.photos.0.0.file_id").String()
	if fileID == "" {
		return "", nil
	}
	payload2 := fmt.Sprintf(`{"file_id":"%s"}`, fileID)
	data2, err := b.dataSource.Execute(ctx, "getFile", []byte(payload2))
	if err != nil {
		return "", err
	}
	filePath := gjson.GetBytes(data2, "result.file_path").String()
	if filePath == "" {
		return "", nil
	}
	return fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", b.token, filePath), nil
}

func NewBot(token string, dataSource transport.Client) Bot {
	return bot{
		dataSource: dataSource,
		token:      token,
	}
}
