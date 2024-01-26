package main

import (
	"os"

	"github.com/spf13/cobra"
)

type globalCmd struct {
	*cobra.Command
}

func main() {
	gc := &globalCmd{
		Command: &cobra.Command{
			Use:   "key-assign",
			Short: "for generate / decode key with expired time",
		},
	}

	gc.RunE = func(cmd *cobra.Command, args []string) error {
		return gc.Help()
	}

	generateCmd := &cmdGenerate{}
	gc.AddCommand(generateCmd.Command())

	decodeCmd := &cmdDecode{}
	gc.AddCommand(decodeCmd.Command())

	if err := gc.Execute(); err != nil {
		os.Exit(1)
	}
}
