package cmd

import (
	"cyaml/combine"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var combineCmd = &cobra.Command{
	Use:   "combine path",
	Short: "Combine yaml files",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		maybeStrip, _ := cmd.Flags().GetString("strip")
		cyaml, err := combine.Exec(args[0], maybeStrip)
		if err != nil {
			log.Errorf("Can't combine yaml, error:%s", err)
			return
		}
		re, err := yaml.Marshal(cyaml)
		if err != nil {
			log.Errorf("Can't marshal combined yaml, error:%s", err)
			return
		}
		fmt.Printf("%s", re)
	},
}

func init() {
	combineCmd.Flags().StringP("strip", "s", "", "Strip prefix from file names")
}
