package ebs

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/pkg/errors"
)

// Volume contains fields pertaining to an AWS EBS VColume
type Volume struct {
	Created time.Time
	ID      string
	State   string
	Type    string
	Tags    []*ec2.Tag
}

// EBS contains info relating to EBS resources within AWS EC2
type EBS struct {
	client     ec2iface.EC2API
	EBSVolumes []Volume
}

// InitClient configures the API client for EC2
func (ebs *EBS) InitClient(region string) {
	cfg := &aws.Config{
		Region: aws.String(region),
	}
	sess := session.New(cfg)
	ebs.client = ec2.New(sess)
}

// Volumes gets information about EBS Volumes
func (ebs *EBS) Volumes() error {
	input := &ec2.DescribeVolumesInput{}

	output, err := ebs.client.DescribeVolumes(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				return errors.Wrap(aerr, "Known AWS API Error")
			}
		} else {
			// If Error is not a known AWS Error then return a wrapped unknown error
			return errors.Wrap(err, "Unknown Error")
		}
	}

	ebs.EBSVolumes = filterVolumes(output.Volumes)

	return nil
}

// filterVolumes takes the data from the AWS API and pulls out only what we want
func filterVolumes(volumes []*ec2.Volume) []Volume {
	var output []Volume

	for _, v := range volumes {
		output = append(output, Volume{
			Created: *v.CreateTime,
			ID:      *v.VolumeId,
			State:   *v.State,
			Tags:    v.Tags,
			Type:    *v.VolumeType,
		})
	}

	return output
}
