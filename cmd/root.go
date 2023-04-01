package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hello",
	Short: "A brief description of your application",
	Long:  `A longer description of your application`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
