package port

type DownloaderService interface {
	Download(string) (string, error)
}
