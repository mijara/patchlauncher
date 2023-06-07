package port

import "github.com/mijara/patchlauncher/model"

type ScrapperService interface {
	ScrapHackList() ([]model.Hack, error)
}
