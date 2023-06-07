package adapter

import "os"

type FileSystemService struct {
}

func NewFileSystemService() *FileSystemService {
	return &FileSystemService{}
}

func (s *FileSystemService) Delete(path string) error {
	return os.Remove(path)
}
