package cmd

import (
	"github.com/spf13/cobra"
)

var exampleCmd = &cobra.Command{
	Use:   "example_cmd",
	Short: "This is an example command",
	Long:  `This is a longer description of this command`,
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	RootCmd.AddCommand(exampleCmd)
}
