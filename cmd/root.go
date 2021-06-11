package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "api2go",
	Short: "Generate golang models from .yaml specs",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("root command\n")
		os.Exit(1)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
