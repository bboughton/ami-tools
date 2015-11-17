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
	return Image{
		Id:        *img.ImageId,
		CreatedAt: time.Now(),
	}
}
