package port

type FileSystemService interface {
	Delete(path string) error
}
