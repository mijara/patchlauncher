package adapter

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"smwlauncher/model"
	"time"
)

type StorageService struct {
}

func NewStorageService() *StorageService {
	return &StorageService{}
}

func (s *StorageService) StoreHacks(hacks []model.Hack) error {
	storedHacks, _, err := s.LoadHacks()
	if err != nil {
		return err
	}

	storedHacksMap := make(map[string]model.Hack)
	for _, storedHack := range storedHacks {
		storedHacksMap[storedHack.URL] = storedHack
	}

	for _, hack := range hacks {
		storedHacksMap[hack.URL] = hack
	}

	finalHacks := make([]model.Hack, 0)
	for _, storedHack := range storedHacksMap {
		finalHacks = append(finalHacks, storedHack)
	}

	fp, err := os.OpenFile("storage/hacks.json", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0660)
	if err != nil {
		return err
	}
	defer fp.Close()

	return json.NewEncoder(fp).Encode(&HackStorage{
		UpdatedAt: time.Now(),
		Hacks:     finalHacks,
	})
}

func (s *StorageService) LoadHacks() ([]model.Hack, time.Time, error) {
	fp, err := os.OpenFile("storage/hacks.json", os.O_RDWR|os.O_CREATE, 0660)
	if err != nil {
		return nil, time.Time{}, err
	}
	defer fp.Close()

	hackStorage := &HackStorage{}
	if err := json.NewDecoder(fp).Decode(hackStorage); err != nil {
		if errors.Is(err, io.EOF) {
			return nil, time.Time{}, nil
		}
		return nil, time.Time{}, err
	}

	return hackStorage.Hacks, hackStorage.UpdatedAt, nil
}

type HackStorage struct {
	UpdatedAt time.Time
	Hacks     []model.Hack
}
