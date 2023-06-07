package interactor

import (
	"smwlauncher/model"
	"smwlauncher/port"
)

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
	Hacks []model.Hack
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
