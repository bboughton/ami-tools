package main

import (
	"fmt"
	"os"

	"github.com/bboughton/ami-tools/ami"
	"github.com/bboughton/ami-tools/log"
)

var (
	Name    string
	Version string
	Build   string
)

func main() {
	logger := log.NewLogger(false)
	client := ami.NewService(false, logger)
	ids := os.Args[1:]
	images := client.Find(ami.FindFilter{
		Ids: ids,
	})
	for _, img := range images {
		fmt.Println(img.Id, img.CreatedBy, img.CreatedAt)
	}
}
