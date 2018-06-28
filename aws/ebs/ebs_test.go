package ebs

import (
	"reflect"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

type mockEC2 struct {
	ec2iface.EC2API
}

func createMockVolume(name, team, state string) *ec2.Volume {
	var volumeTags []*ec2.Tag
	nameTag := &ec2.Tag{
		Key:   aws.String("Name"),
		Value: aws.String(name),
	}

	teamTag := &ec2.Tag{
		Key:   aws.String("Team"),
		Value: aws.String(team),
	}

	envTag := &ec2.Tag{
		Key:   aws.String("Environment"),
		Value: aws.String("Development"),
	}

	volumeTags = append(volumeTags, nameTag, teamTag, envTag)
	createdTime := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)

	mockVolume := &ec2.Volume{
		VolumeId:   aws.String("vol-0123456789abcdefa"),
		VolumeType: aws.String("gp2"),
		CreateTime: &createdTime,
		Encrypted:  aws.Bool(true),
		State:      aws.String(state),
		Tags:       volumeTags,
	}

	return mockVolume

}

// DescribeVolumes is an AWS API function that has been mocked for testing
func (m mockEC2) DescribeVolumes(*ec2.DescribeVolumesInput) (*ec2.DescribeVolumesOutput, error) {
	var mockVolumes []*ec2.Volume
	attachedVolume := createMockVolume("attachedVolume", "Nalgene", "in-use")
	detachedVolume := createMockVolume("detachedVolume", "Peli", "available")

	mockVolumes = append(mockVolumes, attachedVolume, detachedVolume)

	mockOutput := &ec2.DescribeVolumesOutput{
		Volumes: mockVolumes,
	}

	return mockOutput, nil
}

// Ensure that the Volumes function works as expected
func TestVolumes(t *testing.T) {
	testEBS := EBS{
		client: mockEC2{},
	}

	err := testEBS.Volumes()
	if err != nil {
		t.Error("Unexpected Error Occured", err)
	}

	if len(testEBS.EBSVolumes) != 2 {
		t.Errorf("Error. want: %v, got: %v", 2, len(testEBS.EBSVolumes))
	}
}

func TestDetailedTextOutput(t *testing.T) {
	testEBS := EBS{
		client: mockEC2{},
	}

	err := testEBS.Volumes()
	if err != nil {
		t.Error("Unexpected Error Occured", err)
	}

	want := []string{
		"ID\tType\tState\tCreated On\t",
		"vol-0123456789abcdefa\tgp2\tin-use\t17-11-2009 20:34:58\t",
		"vol-0123456789abcdefa\tgp2\tavailable\t17-11-2009 20:34:58\t",
	}

	got := testEBS.DetailedTextOutput()

	if reflect.DeepEqual(want, got) != true {
		t.Errorf("Error. want: %v, got: %v", want, got)
	}

}

// func TestExampleFunction(t *testing.T) {
// 	type args struct {
// 		s string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want int
// 	}{
// 		{
// 			"test1",
// 			args{"test1"},
// 			5,
// 		},
// 		{
// 			"test2",
// 			args{""},
// 			0,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := exampleFunction(tt.args.s); got != tt.want {
// 				t.Errorf("exampleFunction() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
