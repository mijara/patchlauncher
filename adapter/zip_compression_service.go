package adapter

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type ZipCompressionService struct {
}

func NewZipCompressionService() *ZipCompressionService {
	return &ZipCompressionService{}
}

func (s *ZipCompressionService) Extract(path, which string) (string, error) {
	if err := s.validateZip(path); err != nil {
		return "", err
	}

	archive, err := zip.OpenReader(path)
	if err != nil {
		return "", err
	}
	defer archive.Close()

	destPath := filepath.Join(filepath.Dir(path), filepath.Base(which))
	dest, err := os.Create(destPath)
	if err != nil {
		return "", err
	}
	defer dest.Close()

	for _, file := range archive.File {
		if file.Name != which {
			continue
		}

		fp, err := file.Open()
		if err != nil {
			return "", err
		}
		defer fp.Close()

		if _, err := io.Copy(dest, fp); err != nil {
			return "", err
		}
	}

	return dest.Name(), nil
}

func (s *ZipCompressionService) GetFilesMatchingSuffix(path, suffix string) ([]string, error) {
	if err := s.validateZip(path); err != nil {
		return nil, err
	}

	results := make([]string, 0)

	archive, err := zip.OpenReader(path)
	if err != nil {
		return nil, err
	}
	defer archive.Close()

	for _, file := range archive.File {
		if !strings.HasSuffix(file.Name, suffix) {
			continue
		}

		results = append(results, file.Name)
	}

	return results, nil
}

func (s *ZipCompressionService) validateZip(path string) error {
	if filepath.Ext(path) == "zip" {
		return fmt.Errorf("not a zip file: %s", path)
	}
	return nil
}
