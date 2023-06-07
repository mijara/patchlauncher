package interactor

import (
	"github.com/mijara/patchlauncher/model"
	"github.com/mijara/patchlauncher/port"
	"go.uber.org/zap"
	"sort"
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
	hacks := make([]model.Hack, 0)
	cached := false

	if !input.RefreshCache {
		cachedHacks, updatedAt, err := it.storageService.LoadHacks()
		if err != nil {
			return nil, err
		}

		if time.Now().Sub(updatedAt) < time.Minute*15 {
			hacks = cachedHacks
			cached = true
		}
	}

	if !cached {
		scrappedHacks, err := it.scrapperService.ScrapHackList()
		if err != nil {
			return nil, err
		}

		hacks = scrappedHacks
	}

	if err := it.storageService.StoreHacks(hacks); err != nil {
		return nil, err
	}

	// Sort hacks.
	sort.Slice(hacks, func(i, j int) bool {
		return hacks[i].UploadedAt.After(hacks[j].UploadedAt)
	})

	it.logger.Info("hack list gotten",
		zap.Any("input", input),
		zap.Bool("cache", cached),
	)

	return &GetHackListOutput{
		Hacks: hacks,
	}, nil
}
