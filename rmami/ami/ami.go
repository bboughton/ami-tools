package ami

import (
	"fmt"

	"github.com/bboughton/ami-tools/rmami/log"
	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/ec2"
)

type Service struct {
	ec2 *ec2.EC2

	// When set to true we won't delete any resource
	dry bool

	log log.Logger
}

// NewService returns a new AMI Service
func NewService(dry bool, logger log.Logger) *Service {
	auth, err := aws.EnvAuth()
	if err != nil {
		return nil
	}

	return &Service{
		ec2: ec2.New(auth, aws.USWest2),
		dry: dry,
		log: logger,
	}
}

// Remove will detach the ami and delete all corrisponding snapshots
func (srv *Service) Remove(ami string) {
	srv.log.Debug("removing ami " + ami)
	if !srv.dry {
		_, err := srv.ec2.DeregisterImage(ami)
		if err != nil {
			return
		}
	}

	srv.deleteSnapshots(srv.listSnapshots(ami))
}

func (srv *Service) deleteSnapshots(snapIds []string) {
	srv.log.Debug(fmt.Sprint("removing snapshots ", snapIds))
	if !srv.dry {
		srv.ec2.DeleteSnapshots(snapIds)
	}
}

func (srv *Service) listSnapshots(ami string) []string {
	var snapIds []string

	filters := ec2.NewFilter()
	filters.Add("description", fmt.Sprintf("*%s*", ami))
	resp, err := srv.ec2.Snapshots(nil, filters)
	if err != nil {
		return snapIds
	}

	snaps := resp.Snapshots

	for _, snap := range snaps {
		snapIds = append(snapIds, snap.Id)
	}

	return snapIds
}
