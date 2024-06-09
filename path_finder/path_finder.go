package path_finder

import (
	"log"
	"os"
	"path/filepath"
)

func GetPathLastElement(path string) (rootPath string, lastElement string) {
	return filepath.Split(path)
}

func CombinePath(path string, element string) string {
	return filepath.Join(path, element)
}

func GetWorkingDirectoryPath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}
