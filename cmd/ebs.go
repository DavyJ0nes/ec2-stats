package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/davyj0nes/ec2-stats/aws/ebs"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// ebsCmd outputs information about EBS Volumes
var ebsCmd = &cobra.Command{
	Use:   "ebs",
	Short: "Generates stats about AWS EBS Volumes",
	Long: `Example
	ec2-stats ebs`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := ebsCommand(); err != nil {
			fmt.Printf("Error Running ebsCommand:\n %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(ebsCmd)
	// copyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func ebsCommand() error {
	// Initialise ebs Client
	ebs := ebs.New("eu-west-1")

	// Get EBS Volume Information
	err := ebs.Volumes()
	if err != nil {
		return errors.Wrap(err, "Error getting Volume info:")
	}

	output(ebs.TextOutput())
	return nil
}

func output(data []string) {
	padding := 3
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', tabwriter.AlignRight|tabwriter.Debug)

	for _, line := range data {
		fmt.Fprintln(w, line)
	}

	w.Flush()
}
