package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"

	"github.com/bboughton/ami-tools/ami"
	"github.com/bboughton/ami-tools/rmami/log"
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
		if valid(ami) {
			client.Remove(ami)
		} else {
			errorOccured = true
			logger.Info(fmt.Sprintf("%s: invalid ami id\n", ami))
		}
	}

	if errorOccured {
		os.Exit(1)
	}
}

func getDebug() bool {
	return os.Getenv("DEBUG") == "1"
}

func valid(ami string) bool {
	validId := regexp.MustCompile(`^ami-[a-zA-Z0-9]+$`)
	return validId.MatchString(ami)
}
