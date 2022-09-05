package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "api2go",
	Short: "Generate golang models from .yaml specs",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("root command\n")
		os.Exit(1)
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
	_:
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
