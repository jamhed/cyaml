package cmd

import (
	"cyaml/split"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var splitCmd = &cobra.Command{
	Use:   "split cyamlFile [path]",
	Short: "Split cyaml file into files",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := split.Exec(args[0], args[1])
		if err != nil {
			log.Errorf("Can't split cyaml, error:%s", err)
			return
		}
	},
}

func init() {
}
