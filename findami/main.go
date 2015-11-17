package main

import (
	"flag"
	"fmt"

	"github.com/bboughton/ami-tools/ami"
	"github.com/bboughton/ami-tools/log"
)

var (
	Name    string
	Version string
	Build   string
)

func main() {
	var (
		usr    string
		latest bool
	)

	flag.StringVar(&usr, "created-by", "", "filter for images by the user that created them")
	flag.BoolVar(&latest, "latest", false, "only show latest")
	flag.Parse()

	logger := log.NewLogger(false)
	client := ami.NewService(false, logger)
	filter := ami.FindFilter{
		CreatedBy: usr,
		Latest:    latest,
	}
	for _, image := range client.Find(filter) {
		fmt.Println(image.Id)
	}
}
