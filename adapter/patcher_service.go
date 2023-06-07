package adapter

import (
	"fmt"
	"os/exec"
)

type PatcherService struct {
	patcherCmdPath string
}

func NewPatcherService(patcherCmdPath string) *PatcherService {
	return &PatcherService{patcherCmdPath: patcherCmdPath}
}

func (s *PatcherService) Patch(
	romPath, patchPath, outputPath string,
) error {
	cmd := exec.Command(s.patcherCmdPath, "--apply", patchPath, romPath, outputPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s: %w", output, err)
	}
	return nil
}
