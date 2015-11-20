package ami

import (
	"errors"
	"fmt"
	"sort"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/bboughton/ami-tools/log"
)

type Service struct {
	ec2 *ec2.EC2

	// When set to true we won't delete any resource
	dry bool

	log log.Logger
}

// NewService returns a new AMI Service
func NewService(dry bool, logger log.Logger) *Service {
	sess := session.New(&aws.Config{Region: aws.String("us-west-2")})

	return &Service{
		ec2: ec2.New(sess),
		dry: dry,
		log: logger,
	}
}

// Remove will detach the ami and delete all corrisponding snapshots
func (srv *Service) Remove(ami string) error {
	srv.log.Debug("removing ami " + ami)
	_, err := srv.ec2.DeregisterImage(&ec2.DeregisterImageInput{
		ImageId: aws.String(ami),
		DryRun:  aws.Bool(srv.dry),
	})
	if err != nil {
		return err
	}

	err = srv.deleteSnapshots(srv.listSnapshots(ami))
	if err != nil {
		return err
	}
	return nil
}

func (srv *Service) deleteSnapshots(snapIds []string) error {
	var isError bool
	for _, snapId := range snapIds {
		srv.log.Debug(fmt.Sprint("removing snapshot ", snapId))
		_, err := srv.ec2.DeleteSnapshot(&ec2.DeleteSnapshotInput{
			SnapshotId: aws.String(snapId),
			DryRun:     aws.Bool(srv.dry),
		})
		if err != nil {
			srv.log.Debug(fmt.Sprint("failed removing snapshot ", snapId))
			isError = true
		}

	}
	if isError {
		return errors.New("failed to remove snapshots")
	}
	return nil
}

func (srv *Service) listSnapshots(ami string) []string {
	var snapIds []string

	resp, err := srv.ec2.DescribeSnapshots(&ec2.DescribeSnapshotsInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("description"),
				Values: []*string{
					aws.String(fmt.Sprintf("*%s*", ami)),
				},
			},
		},
		OwnerIds: []*string{
			aws.String("self"),
		},
	})
	if err != nil {
		return snapIds
	}

	snaps := resp.Snapshots

	for _, snap := range snaps {
		snapIds = append(snapIds, *snap.SnapshotId)
	}

	return snapIds
}

func (srv *Service) Find(filter FindFilter) Images {
	resp, err := srv.ec2.DescribeImages(&ec2.DescribeImagesInput{
		Filters:  filter.ec2filter(),
		ImageIds: filter.imageIds(),
		Owners: []*string{
			aws.String("self"),
		},
	})
	if err != nil {
		srv.log.Debug(err.Error())
		return nil
	}

	images := Images{}
	for _, img := range resp.Images {
		images.Add(newImageFromEc2Image(img))
	}

	if filter.Latest {
		srv.log.Debug("filtering out all but latest")
		sort.Sort(byCreatedAt(images))
		images = images[:1]
	}

	return images
}

type FindFilter struct {
	CreatedBy string
	Latest    bool
	Ids       []string
}

func (f FindFilter) ec2filter() []*ec2.Filter {
	var ec2filter []*ec2.Filter

	if f.CreatedBy != "" {
		ec2filter = append(ec2filter, &ec2.Filter{
			Name: aws.String("tag:Created By"),
			Values: []*string{
				aws.String(f.CreatedBy),
			},
		})
	}

	return ec2filter
}

func (f FindFilter) imageIds() []*string {
	if len(f.Ids) == 0 {
		return nil
	}
	return aws.StringSlice(f.Ids)
}
