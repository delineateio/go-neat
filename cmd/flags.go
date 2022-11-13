package cmd

import (
	e "github.com/delineateio/go-neat/errors"
	"github.com/spf13/cobra"
)

func addStrFlag(cmd *cobra.Command, name, desc string, value *string) {
	cmd.Flags().StringVar(value, name, "", desc)
}

func addRequired(name string) {
	err := newFeatureCmd.MarkFlagRequired(name)
	e.CheckIfError(err, "failed to initialise the '%s' flag", name)
}
