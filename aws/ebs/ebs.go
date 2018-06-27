package ebs

import (
	"fmt"
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

// New creates a pointer to EBS instance with an initialised client
func New(region string) *EBS {
	ebs := EBS{}
	ebs.initClient(region)

	return &ebs
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

// TextOutput returns the slice of data about all the EBS Volumes found.
// output is formatted so it ca be used by tabwriter for better CLI formatting.
func (ebs *EBS) TextOutput() []string {
	var output []string

	// add headings to output
	output = append(output, "ID\tType\tState\tCreated On\t")

	for _, v := range ebs.EBSVolumes {
		// setting format to DD-MM-YYYY HH:MM:SS
		timeString := v.Created.Format("02-01-2006 15:04:05")

		output = append(output,
			fmt.Sprintf("%s\t%s\t%s\t%s\t", v.ID, v.Type, v.State, timeString))
	}

	return output
}

// initClient configures the API client for EC2
func (ebs *EBS) initClient(region string) {
	cfg := &aws.Config{
		Region: aws.String(region),
	}
	sess := session.New(cfg)
	ebs.client = ec2.New(sess)
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
