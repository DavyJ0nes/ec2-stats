# EC2 Stats

> This is still in very early development so isn't super useful at the moment.

[![Go Report Card](https://goreportcard.com/badge/github.com/DavyJ0nes/ec2-stats)](https://goreportcard.com/report/github.com/DavyJ0nes/ec2-stats)

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

[MIT](./LICENSE)
