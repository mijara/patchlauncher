package adapter

import (
	"fmt"
	"os/exec"
)

type OpenerService struct {
}

func NewOpenerService() *OpenerService {
	return &OpenerService{}
}

func (o *OpenerService) Open(file string) error {
	cmd := exec.Command("open", file)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s: %w", string(output), err)
	}

	return nil
}
