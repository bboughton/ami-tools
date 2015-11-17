package main

import (
	"flag"
	"os"

	"github.com/bboughton/ami-tools/ami"
	"github.com/bboughton/ami-tools/log"
)

var (
	Name     string
	Version  string
	Revision string
	DEBUG    = false
)

func main() {
	DEBUG = getDebug()
	logger := log.NewLogger(DEBUG)

	var dry bool

	fs := flag.NewFlagSet(Name, flag.ExitOnError)
	fs.BoolVar(&dry, "dry", false, "do a dry run")
	fs.Parse(os.Args[1:])

	amis := fs.Args()

	client := ami.NewService(dry, logger)
	var errorOccured bool
	for _, ami := range amis {
		err := client.Remove(ami)
		if err != nil {
			logger.Info(err.Error())
		}
	}

	if errorOccured {
		os.Exit(1)
	}
}

func getDebug() bool {
	return os.Getenv("DEBUG") == "1"
}
