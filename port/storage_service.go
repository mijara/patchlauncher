package port

import (
	"github.com/mijara/patchlauncher/model"
	"time"
)

type StorageService interface {
	StoreHacks(hacks []model.Hack) error
	LoadHacks() ([]model.Hack, time.Time, error)
}
