package interactor

import (
	"github.com/mijara/patchlauncher/port"
	"go.uber.org/zap"
)

type DeleteFiles struct {
	logger            port.Logger
	fileSystemService port.FileSystemService
}

func NewDeleteFiles(
	logger port.Logger, fileSystemService port.FileSystemService,
) *DeleteFiles {
	return &DeleteFiles{
		logger:            logger,
		fileSystemService: fileSystemService,
	}
}

type DeleteFilesInput struct {
	Files []string
}

func (it *DeleteFiles) Execute(input DeleteFilesInput) error {
	for _, file := range input.Files {
		if err := it.fileSystemService.Delete(file); err != nil {
			return err
		}
	}

	it.logger.Info("files deleted", zap.Any("input", input))

	return nil
}
