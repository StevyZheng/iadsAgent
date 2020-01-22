package common

import (
	"io/ioutil"
	"path/filepath"
)

func ListFiles(path string) (files []string, err error) {
	files = make([]string, 0)
	fs, err := ioutil.ReadDir(path)
	for _, file := range fs {
		if !file.IsDir() {
			files = append(files, filepath.Join(path, file.Name()))
		}
	}
	return files, err
}
