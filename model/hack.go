package model

type Hack struct {
	URL         string // Used as primary key.
	Title       string
	DownloadURL string
	Authors     string
	Rating      string
	Downloads   string
	Type        string
}
