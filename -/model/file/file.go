package file

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func GetFile(ext, pathS string) ([]string, []string) {
	var files []string
	var html []string
	filepath.Walk(pathS, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			r, err := regexp.MatchString(ext, f.Name())
			if err == nil && r {
				html = append(html, ReadFile(path))
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

func ConcatFile(pathS string) (string, string, string) {
	var html, css, js string
	filepath.Walk(pathS, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			r, err := regexp.MatchString(".html", f.Name())
			if err == nil && r {
				html = ReadFile(path)
			}
			r, err = regexp.MatchString(".css", f.Name())
			if err == nil && r {
				css += ReadFile(path)
			}
			r, err = regexp.MatchString(".js", f.Name())
			if err == nil && r {
				js += ReadFile(path)
			}

		}
		return nil
	})
	return html, css, js
}

func ReadFile(pathS string) string {
	var files string
	file, _ := os.Open(pathS)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		files += fmt.Sprintln(scanner.Text())
	}
	return files
}

func WriteFile(file, path string) {
	ioutil.WriteFile(path, []byte(file), 0644)
}
