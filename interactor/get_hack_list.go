package interactor

import "smwlauncher/port"

type GetHackList struct {
	logger          port.Logger
	scrapperService port.ScrapperService
}

func NewGetHackList(
	logger port.Logger,
	scrapperService port.ScrapperService,
) *GetHackList {
	return &GetHackList{
		logger:          logger,
		scrapperService: scrapperService,
	}
}

type GetHackListOutput struct {
	Hacks []port.ROMHack
}

func (it *GetHackList) Execute() (GetHackListOutput, error) {
	hacks, err := it.scrapperService.ScrapHackList()
	if err != nil {
		return GetHackListOutput{}, err
	}

	it.logger.Info("hack list gotten")

	return GetHackListOutput{
		Hacks: hacks,
	}, nil
}
