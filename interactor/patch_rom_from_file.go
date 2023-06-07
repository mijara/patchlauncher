package interactor

import (
	"github.com/mijara/patchlauncher/port"
	"go.uber.org/zap"
)

type PatchRomFromFile struct {
	patcherService port.PatcherService
	logger         port.Logger
}

func NewPatchRomFromFile(
	patcherService port.PatcherService,
	logger port.Logger,
) *PatchRomFromFile {
	return &PatchRomFromFile{
		patcherService: patcherService,
		logger:         logger,
	}
}

type PatchRomFromFileInput struct {
	ROMPath    string
	PatchPath  string
	OutputPath string
}

func (it *PatchRomFromFile) Execute(input PatchRomFromFileInput) error {
	if err := it.patcherService.Patch(
		input.ROMPath,
		input.PatchPath,
		input.OutputPath,
	); err != nil {
		return err
	}

	it.logger.Info("rom patched from file",
		zap.Any("input", input),
	)

	return nil
}
