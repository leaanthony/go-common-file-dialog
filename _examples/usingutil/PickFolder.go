package main

import (
	"github.com/leaanthony/go-common-file-dialog/cfd"
	"github.com/leaanthony/go-common-file-dialog/cfdutil"
	"log"
)

func main() {
	result, err := cfdutil.ShowPickFolderDialog(cfd.DialogConfig{
		Title:  "Pick Folder",
		Role:   "PickFolderExample",
		Folder: "C:\\",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Chosen folder: %s\n", result)
}
