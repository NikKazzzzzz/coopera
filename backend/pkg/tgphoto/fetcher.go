package tgphoto

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Fetcher struct {
	token  string
	client *http.Client
}

func New(token string) *Fetcher {
	return &Fetcher{
		token:  token,
		client: &http.Client{Timeout: 5 * time.Second},
	}
}

func (f *Fetcher) GetPhotoURL(userID int64) (string, error) {
	if f.token == "" {
		return "", nil
	}

	photosURL := fmt.Sprintf("https://api.telegram.org/bot%s/getUserProfilePhotos?user_id=%d&limit=1", f.token, userID)
	resp, err := f.client.Get(photosURL)
	if err != nil {
		return "", fmt.Errorf("getUserProfilePhotos: %w", err)
	}
	defer resp.Body.Close()

	var photosResult struct {
		OK     bool `json:"ok"`
		Result struct {
			Photos [][]struct {
				FileID string `json:"file_id"`
			} `json:"photos"`
		} `json:"result"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&photosResult); err != nil {
		return "", fmt.Errorf("decode getUserProfilePhotos: %w", err)
	}
	if !photosResult.OK || len(photosResult.Result.Photos) == 0 || len(photosResult.Result.Photos[0]) == 0 {
		return "", nil
	}

	// берём наибольший размер (последний в массиве)
	photos := photosResult.Result.Photos[0]
	fileID := photos[len(photos)-1].FileID

	fileURL := fmt.Sprintf("https://api.telegram.org/bot%s/getFile?file_id=%s", f.token, fileID)
	fileResp, err := f.client.Get(fileURL)
	if err != nil {
		return "", fmt.Errorf("getFile: %w", err)
	}
	defer fileResp.Body.Close()

	var fileResult struct {
		OK     bool `json:"ok"`
		Result struct {
			FilePath string `json:"file_path"`
		} `json:"result"`
	}
	if err := json.NewDecoder(fileResp.Body).Decode(&fileResult); err != nil {
		return "", fmt.Errorf("decode getFile: %w", err)
	}
	if !fileResult.OK || fileResult.Result.FilePath == "" {
		return "", nil
	}

	return fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", f.token, fileResult.Result.FilePath), nil
}
