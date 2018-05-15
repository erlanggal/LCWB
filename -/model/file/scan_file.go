package file

import (
	//"fmt"
	function "github.com/kevinrizkhy/LCWB/-/model/utility"
	"os"
	"path/filepath"
	"regexp"
)

func GetFile(ext, pathS string) ([]string, []string) {
	var files []string
	var html []string
	filepath.Walk(pathS, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			r, err := regexp.MatchString(ext, f.Name())
			if err == nil && r {
				html = append(html, function.ConcatFile(path))
			}
		} else {
			if f.Name() != "css" && f.Name() != "js" {
				files = append(files, f.Name())
			}
		}
		return nil
	})
	return files, html
}
