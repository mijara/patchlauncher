package adapter

import (
	"github.com/google/uuid"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type DownloaderService struct {
	downloadBase string
}

func NewDownloaderService(
	downloadBase string,
) *DownloaderService {
	return &DownloaderService{
		downloadBase: downloadBase,
	}
}

func (s *DownloaderService) Download(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	targetPath := filepath.Join(s.downloadBase, uuid.New().String()+".zip")
	file, err := os.Create(targetPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	if _, err := io.Copy(file, response.Body); err != nil {
		return "", err
	}

	return file.Name(), nil
}
