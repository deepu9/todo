package utils

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
)

func CheckPathExists(filePath string) (fs.FileInfo, bool) {
	fileInfo, err := os.Stat(filePath)

	return fileInfo, !errors.Is(err, os.ErrNotExist)
}

func getCwd() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return dir
}

func GetDataPath() string {
	dataDir := filepath.Join(getCwd(), DataDirName)
	if _, err := os.Stat(dataDir); errors.Is(err, os.ErrNotExist) {
		if err := os.Mkdir(dataDir, os.ModePerm); err != nil {
			panic(err)
		}
	}

	return dataDir
}
