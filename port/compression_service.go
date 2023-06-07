package port

type CompressionService interface {
	Extract(path, which string) (string, error)
	GetFilesMatchingSuffix(filepath, suffix string) ([]string, error)
}
