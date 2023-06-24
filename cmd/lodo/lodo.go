package main

import (
	"fmt"
	"time"

	"github.com/dariuszkorolczukcom/lodo/cmd/view"
	"github.com/dariuszkorolczukcom/lodo/internal/file"
	"github.com/dariuszkorolczukcom/lodo/internal/walker"
)

var fileList []file.File

func main() {
	// Get a greeting message and print it.
	view.Create(run)
}

func run(fileName string) []file.File {
	fmt.Println("run: " + fileName)
	fileList = walker.Walk(fileName)
	for _, file := range fileList {
		file.Read()
		file.IdentifyAll()
		fmt.Printf("File Name: %s\nModified: %s\nSize: %.2fkB\nPesels: %d\nPost Codes: %d\nNames: %d\nPhones: %d\nEmails: %d\n\n",
			file.Name,
			file.Modified.Format(time.Stamp),
			float64(file.Size/1024),
			file.PeselNo,
			file.PostCodeNo,
			file.NameNo,
			file.PhoneNo,
			file.EmailNo)
	}
	return fileList
}
