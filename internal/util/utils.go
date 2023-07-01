package util

import (
	"errors"
	"os"
	"path/filepath"
)

const MAX_DIR_DEPTH = 100

func GetSomethingPath(base string, something string, returnRelPath bool) (string, error) {
	baseAbsPath, err := filepath.Abs(base)
	if err != nil {
		return "", err
	}

	current := baseAbsPath
	for i := 0; i < MAX_DIR_DEPTH; i++ {
		if FileExists(filepath.Join(current, something)) {
			if returnRelPath {
				rel, err := filepath.Rel(baseAbsPath, current)
				if err != nil {
					return "", err
				}

				return filepath.Join(rel, something), nil
			}

			return filepath.Join(current, something), nil
		}

		if current == "/" {
			return "", errors.New("reached the root directory")
		}

		current = filepath.Dir(current)
	}

	return "", errors.New("reached the depth-limit")
}

func LoadFileAsString(path string) (string, error) {
	if !IsFile(path) {
		return "", errors.New("this is not file")
	}

	contents, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(contents), nil
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func FileNotExists(path string) bool {
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}

func IsFile(path string) bool {
	fileInfo, err := os.Stat(path)
	return err == nil && !fileInfo.IsDir()
}
