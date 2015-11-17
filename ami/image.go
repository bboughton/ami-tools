package ami

import (
	"time"

	"github.com/aws/aws-sdk-go/service/ec2"
)

type Image struct {
	Id        string
	CreatedAt time.Time
}

type Images []Image

func (imgs *Images) Add(img Image) {
	*imgs = append(*imgs, img)
}

func newImageFromEc2Image(img *ec2.Image) Image {
	createdAt, err := time.Parse(time.RFC3339, *img.CreationDate)
	if err != nil {
		createdAt = time.Now()
	}

	return Image{
		Id:        *img.ImageId,
		CreatedAt: createdAt,
	}
}

// byCreatedAt should be used to sort Images collection by CreatedAt
type byCreatedAt []Image

func (img byCreatedAt) Len() int {
	return len(img)
}

func (img byCreatedAt) Swap(i, j int) {
	img[i], img[j] = img[j], img[i]
}

func (img byCreatedAt) Less(i, j int) bool {
	return img[i].CreatedAt.After(img[j].CreatedAt)
}
