package convert

import (
	"io/ioutil"
	"log"
	"path/filepath"
)

func Dirsearch(dirname string) []string {

	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {

			paths = append(paths, Dirsearch(filepath.Join(dirname, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dirname, file.Name()))
	}

	return paths
}
