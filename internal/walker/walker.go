package walker

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/dariuszkorolczukcom/lodo/internal/file"
)

var result []file.File

func Walk(path string) []file.File {
	fmt.Println("Walk " + strings.Split(path, "file://")[1])
	location, err := os.Getwd()
	if err != nil {
		log.Fatalf(err.Error())
	}
	relPath, err := filepath.Rel(location, strings.Split(path, "file://")[1])

	fmt.Println("location " + location)
	fmt.Println("relPath " + relPath)
	if err != nil {
		log.Fatal(err)
	}
	iterate(relPath)
	return result
}

func iterate(path string) {
	fmt.Println("iterate" + path)

	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalf(err.Error())
		}
		var file file.File
		file.Modified = info.ModTime()
		file.Size = info.Size()
		file.Name = path
		splitPath := strings.Split(path, "/")
		if !info.IsDir() {
			result = append(result, file)
		} else if info.Name() != splitPath[len(splitPath)-1] {
			iterate(filepath.Join(path, info.Name()))
		}
		return nil
	})
}
