package interactor

import (
	"github.com/mijara/patchlauncher/port"
	"go.uber.org/zap"
)

type GetCompressedHackPatches struct {
	logger             port.Logger
	compressionService port.CompressionService
}

func NewGetCompressedHackPatches(
	logger port.Logger,
	compressionService port.CompressionService,
) *GetCompressedHackPatches {
	return &GetCompressedHackPatches{
		logger:             logger,
		compressionService: compressionService,
	}
}

type GetCompressedHackPatchesInput struct {
	CompressedHackPath string
}

func (it *GetCompressedHackPatches) Execute(input GetCompressedHackPatchesInput) ([]string, error) {
	patches, err := it.compressionService.GetFilesMatchingSuffix(input.CompressedHackPath, ".bps")
	if err != nil {
		return nil, err
	}

	it.logger.Info("compressed hack patches gotten",
		zap.Any("input", input),
	)

	return patches, nil
}
