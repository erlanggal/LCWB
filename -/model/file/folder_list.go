package file

import (
	"os"
	"path/filepath"
	"strings"
)

func FolderList(pathS string) []string {
	var list []string
	parent_folder := strings.Split(pathS, "/")[1]
	filepath.Walk(pathS, func(path string, f os.FileInfo, _ error) error {
		if f.IsDir() {
			if f.Name() != "css" && f.Name() != "js" && f.Name() != parent_folder {
				list = append(list, f.Name())
			}
		}
		return nil
	})
	return list
}
