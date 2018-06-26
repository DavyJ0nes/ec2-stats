# EC2 Stats

## Description

Simple cli tool to gather information about EC2 resources within an AWS region of an AWS account.

Main motivation behind building this was to work on testing and mocking the aws-sdk-go library.

## Usage

Usage instructions with code examples

```shell
# Install to $GOPATH/bin
go install

# Return stats about EBS volumes to STDOUT
ec2-stats ebs

# Return stats about EBS volumes as JSON blob
ec2-stats ebs --json
```

## License

MIT
