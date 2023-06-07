package port

import "smwlauncher/model"

type ScrapperService interface {
	ScrapHackList() ([]model.Hack, error)
}
