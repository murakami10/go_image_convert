package convert

import (
	"strings"
)

type File struct {
	Name string
	Ext  string
}

func Tofile(filenames []string) []File {

	var files []File

	for _, filename := range filenames {
		slice := strings.Split(filename, ".")
		if len(slice) == 1 {
			files = append(files, File{filename, ""})
			continue
		}
		files = append(files, File{filename, slice[len(slice)-1]})
	}

	return files
}

func Checkext(check string) bool {

	extlist := [...]string{"jpeg", "jpg", "png", "gif"}
	var ok bool
	for _, ext := range extlist {
		if check == ext {
			ok = true
		}
	}

	return ok
}
