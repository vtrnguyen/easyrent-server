package services

import (
	"easyrent-server/internal/utils"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

type FileService struct{}

// NewFileService creates a new instance of FileService.
func NewFileService() *FileService {
	return &FileService{}
}

func (s *FileService) saveFileToFolder(
	file *multipart.FileHeader,
	folder string,
) (string, error) {
	ext := filepath.Ext(file.Filename)
	filename := uuid.New().String() + ext
	path := filepath.Join(
		"storage",
		folder,
		filename,
	)

	err := os.MkdirAll(
		filepath.Dir(path),
		0755,
	)
	if err != nil {
		return "", err
	}

	err = utils.SaveFile(
		file,
		path,
	)
	if err != nil {
		return "", err
	}

	return "/" + filepath.ToSlash(path), nil
}

// SaveAvatar saves the provided avatar file to the storage and returns the URL of the saved avatar.
func (s *FileService) SaveAvatar(
	file *multipart.FileHeader,
) (string, error) {
	return s.saveFileToFolder(file, "avatars")
}

// SavePropertyImage saves a property image file and returns its URL.
func (s *FileService) SavePropertyImage(
	file *multipart.FileHeader,
) (string, error) {
	return s.saveFileToFolder(file, filepath.Join("properties", "images"))
}

// SavePropertyVideo saves a property video file and returns its URL.
func (s *FileService) SavePropertyVideo(
	file *multipart.FileHeader,
) (string, error) {
	return s.saveFileToFolder(file, filepath.Join("properties", "videos"))
}

// DeleteByURL removes a stored file using its public URL.
func (s *FileService) DeleteByURL(url string) error {
	if url == "" {
		return nil
	}

	path := filepath.FromSlash(strings.TrimPrefix(url, "/"))
	err := os.Remove(path)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	return nil
}
