package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cyaml",
	Short: "A yaml file combiner",
}

func init() {
	rootCmd.AddCommand(combineCmd)
	rootCmd.AddCommand(splitCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
