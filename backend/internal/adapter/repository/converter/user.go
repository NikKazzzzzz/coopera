package converter

import (
	"github.com/NikKazzzzzz/coopera-backend/internal/adapter/repository/model/user_model"
	"github.com/NikKazzzzzz/coopera-backend/internal/entity"
)

func FromEntityToModel(euser entity.UserEntity) user_model.User {
	return user_model.User{
		TelegramID: *euser.TelegramID,
		Username:   *euser.Username,
		PhotoURL:   euser.PhotoURL,
	}
}

func FromModelToEntity(muser user_model.User) entity.UserEntity {
	return entity.UserEntity{
		ID:         &muser.ID,
		TelegramID: &muser.TelegramID,
		Username:   &muser.Username,
		PhotoURL:   muser.PhotoURL,
		CreatedAt:  &muser.CreatedAt,
	}
}

func FromModelToEntityWithTeams(m user_model.UserWithTeams) entity.UserEntity {
	teams := make([]entity.TeamWithRole, len(m.Teams))
	for i, t := range m.Teams {
		teams[i] = entity.TeamWithRole{
			TeamID:    t.ID,
			TeamName:  t.Name,
			Role:      entity.Role(t.Role),
			TeamEmoji: t.Emoji,
			TeamColor: t.Color,
		}
	}

	return entity.UserEntity{
		ID:                 &m.ID,
		TelegramID:         &m.TelegramID,
		Username:           &m.Username,
		PhotoURL:           m.PhotoURL,
		CreatedAt:          &m.CreatedAt,
		Teams:              teams,
		Wallpaper:          m.Wallpaper,
		WallpaperCustomURL: m.WallpaperCustomURL,
		Theme:              m.Theme,
	}
}
