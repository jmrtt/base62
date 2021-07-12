package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "base62",
	Short: "base62 is a command line utility to encode/decode base62 data",
	Long: `base62 encode/decode data using Cobra and base62 implementation
inspired by the Java library https://github.com/seruco/base62`,
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			fmt.Println("Error: ", err)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
