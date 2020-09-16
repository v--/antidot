package dotfile

import (
	"io/ioutil"
	"strings"
)

type Dotfile struct {
	name  string
	isDir bool
}

func isDotfile(filename string) bool {
	return filename != "." && strings.HasPrefix(filename, ".")
}

func Detect(dir string) ([]Dotfile, error) {
	// TODO: better handle file errors
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	found := make([]Dotfile, 0, len(files))
	for _, fileInfo := range files {
		filename := fileInfo.Name()
		if isDotfile(filename) {
			dotfile := Dotfile{
				name:  filename,
				isDir: fileInfo.IsDir(),
			}
			found = append(found, dotfile)
		}
	}

	return found, nil
}
