package interactor

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"net/url"
	"smwlauncher/port"
	"strings"
)

type DownloadCompressedHack struct {
	logger            port.Logger
	downloaderService port.DownloaderService
}

func NewDownloadCompressedHack(
	logger port.Logger, downloaderService port.DownloaderService,
) *DownloadCompressedHack {
	return &DownloadCompressedHack{logger: logger, downloaderService: downloaderService}
}

type DownloadCompressedHackInput struct {
	URL string
}

func (it *DownloadCompressedHack) Execute(input DownloadCompressedHackInput) (string, error) {
	extension, err := it.getFileExtensionFromUrl(input.URL)
	if err != nil {
		return "", err
	}

	// TODO: remove this restriction from here.
	if extension != "zip" {
		return "", fmt.Errorf("extension not supported: %s", extension)
	}

	path, err := it.downloaderService.Download(input.URL)
	if err != nil {
		return "", err
	}

	it.logger.Info("compressed hack downloaded",
		zap.Any("input", input),
	)

	return path, nil
}

func (it *DownloadCompressedHack) getFileExtensionFromUrl(rawUrl string) (string, error) {
	u, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}
	pos := strings.LastIndex(u.Path, ".")
	if pos == -1 {
		return "", errors.New("couldn't find a period to indicate a file extension")
	}
	return u.Path[pos+1 : len(u.Path)], nil
}
