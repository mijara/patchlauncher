package interactor

import (
	"go.uber.org/zap"
	"smwlauncher/port"
)

type ExtractPatch struct {
	logger             port.Logger
	compressionService port.CompressionService
}

func NewExtractPatch(logger port.Logger, compressionService port.CompressionService) *ExtractPatch {
	return &ExtractPatch{logger: logger, compressionService: compressionService}
}

type ExtractPatchInput struct {
	CompressedHackPath string
	PatchPath          string
}

func (it *ExtractPatch) Execute(input ExtractPatchInput) (string, error) {
	path, err := it.compressionService.Extract(input.CompressedHackPath, input.PatchPath)
	if err != nil {
		return "", err
	}

	it.logger.Info("patch extracted", zap.Any("input", input))

	return path, nil
}
