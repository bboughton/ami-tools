package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/ec2"
)

var (
	Name    string
	Version string
	Build   string
	DEBUG   = false
)

func main() {
	var (
		long bool
		usr  string
		err  error
	)

	flag.BoolVar(&long, "long", false, "show long format")
	flag.Parse()

	if os.Getenv("DEBUG") == "1" {
		DEBUG = true
	}

	usr, err = getUser()
	if err != nil {
		info(err.Error())
		os.Exit(1)
	}

	for _, image := range getImages(usr) {
		fmt.Println(format(image, long))
	}
}

// getUser will return the value for the USER env var
//
// Using the env var so that it can be overwritten by the user. Using os/user
// will return the username however it can't be overwritten by setting USER
func getUser() (string, error) {
	var err error
	user := os.Getenv("USER")
	debug(fmt.Sprint("username: ", user))
	if len(user) == 0 {
		err = errors.New("please set USER environment variable")
	}
	return user, err
}

func getImages(user string) []ec2.Image {
	auth, err := aws.EnvAuth()
	if err != nil {
		debug(err.Error())
		return nil
	}
	client := ec2.New(auth, aws.USWest2)

	filters := ec2.NewFilter()
	filters.Add("tag:Created By", user)

	resp, err := client.ImagesByOwners(nil, []string{"self"}, filters)
	if err != nil {
		debug(err.Error())
		return nil
	}

	return resp.Images
}

func debug(msg string) {
	if DEBUG {
		log.Printf("DEBUG %s\n", msg)
	}
}

func info(msg string) {
	log.Printf(msg)
}

func format(img ec2.Image, long bool) string {
	var str string
	if long {
		str = fmt.Sprintf("%s %s", img.Id, img.Name)
	} else {
		str = img.Id
	}
	return str
}
