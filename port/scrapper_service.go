package port

type ScrapperService interface {
	ScrapHackList() ([]ROMHack, error)
}

type ROMHack struct {
	Title       string
	DownloadURL string
	Authors     string
	Rating      string
	Downloads   string
	Type        string
}
