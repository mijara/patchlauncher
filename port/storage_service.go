package port

import (
	"smwlauncher/model"
	"time"
)

type StorageService interface {
	StoreHacks(hacks []model.Hack) error
	LoadHacks() ([]model.Hack, time.Time, error)
}
