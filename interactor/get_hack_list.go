package interactor

import (
	"go.uber.org/zap"
	"smwlauncher/model"
	"smwlauncher/port"
	"time"
)

type GetHackList struct {
	logger          port.Logger
	scrapperService port.ScrapperService
	storageService  port.StorageService
}

func NewGetHackList(
	logger port.Logger,
	scrapperService port.ScrapperService,
	storageService port.StorageService,
) *GetHackList {
	return &GetHackList{
		logger:          logger,
		scrapperService: scrapperService,
		storageService:  storageService,
	}
}

type GetHackListInput struct {
	RefreshCache bool
}

type GetHackListOutput struct {
	Hacks []model.Hack
}

func (it *GetHackList) Execute(input GetHackListInput) (*GetHackListOutput, error) {
	if !input.RefreshCache {
		hacks, updatedAt, err := it.storageService.LoadHacks()
		if err != nil {
			return nil, err
		}

		it.logger.Info("hack list gotten",
			zap.Any("input", input),
			zap.Bool("cache", true),
		)

		if time.Now().Sub(updatedAt) < time.Minute*15 {
			return &GetHackListOutput{
				Hacks: hacks,
			}, nil
		}
	}

	hacks, err := it.scrapperService.ScrapHackList()
	if err != nil {
		return nil, err
	}

	if err := it.storageService.StoreHacks(hacks); err != nil {
		return nil, err
	}

	it.logger.Info("hack list gotten",
		zap.Any("input", input),
		zap.Bool("cache", false),
	)

	return &GetHackListOutput{
		Hacks: hacks,
	}, nil
}
