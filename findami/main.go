package main

import (
	"flag"
	"fmt"

	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/ec2"
)

var (
	Name    string
	Version string
	Build   string
)

func main() {
	var (
		usr string
	)

	flag.StringVar(&usr, "created-by", "", "filter for images by the user that created them")
	flag.Parse()

	for _, image := range getImages(usr) {
		fmt.Println(image.Id)
	}
}

func getImages(user string) []ec2.Image {
	auth, err := aws.EnvAuth()
	if err != nil {
		return nil
	}
	client := ec2.New(auth, aws.USWest2)

	filters := ec2.NewFilter()
	if user != "" {
		filters.Add("tag:Created By", user)
	}

	resp, err := client.ImagesByOwners(nil, []string{"self"}, filters)
	if err != nil {
		return nil
	}

	return resp.Images
}
