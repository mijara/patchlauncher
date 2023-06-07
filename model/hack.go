package model

import "time"

type Hack struct {
	URL         string // Used as primary key.
	Title       string
	DownloadURL string
	Authors     string
	Rating      string
	Downloads   string
	Type        string
	UploadedAt  time.Time
}
