package services

import (
	"easyrent-server/internal/utils"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type FileService struct{}

// NewFileService creates a new instance of FileService.
func NewFileService() *FileService {
	return &FileService{}
}
// SaveAvatar saves the provided avatar file to the storage and returns the URL of the saved avatar.
func (s *FileService) SaveAvatar(
	file *multipart.FileHeader,
) (string,error) {
	ext := filepath.Ext(file.Filename)
	filename := uuid.New().String() + ext
	path := filepath.Join(
		"storage",
		"avatars",
		filename,
	)

	err := os.MkdirAll(
		filepath.Dir(path),
		0755,
	)
	if err != nil {
		return "",err
	}

	err = utils.SaveFile(
		file,
		path,
	)
	if err != nil {
		return "",err
	}

	return "/storage/avatars/"+filename,nil
}