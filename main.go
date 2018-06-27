package main

import (
	"fmt"
	"os"

	"github.com/davyj0nes/ec2-stats/aws/ebs"
)

func main() {
	// Initialise ebs Client
	ebs := ebs.EBS{}
	ebs.InitClient("eu-west-1")

	// Get EBS Volume Information
	err := ebs.Volumes()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(ebs.EBSVolumes)

}
