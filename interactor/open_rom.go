package interactor

import "smwlauncher/port"

type OpenROM struct {
	openerService port.OpenerService
	logger        port.Logger
}

func NewOpenROM(openerService port.OpenerService, logger port.Logger) *OpenROM {
	return &OpenROM{openerService: openerService, logger: logger}
}

type OpenROMInput struct {
	ROMPath string
}

func (it *OpenROM) Execute(input OpenROMInput) error {
	if err := it.openerService.Open(input.ROMPath); err != nil {
		return err
	}

	it.logger.Debug("rom opened")

	return nil
}
