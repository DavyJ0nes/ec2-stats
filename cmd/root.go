package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ec2-stats",
	Short: "Tool to gain insight into your AWS EC2 resources.",
	Long: `ec2-stats allows you to generate statistics about your EC2 resources.
	This includes EC2 Instances, EBS Volumes, Load Balancers, Access Keys etc.
	It's built as a CLI tool to allow for quick data retrieval.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
