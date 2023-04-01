package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

func init() {
	rootCmd.AddCommand(helloCmd)
}

var helloCmd = &cobra.Command{
	Use:   "chat",
	Short: "Print a hello message and prompt the user to select from three different options",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello! Please select an option:")
		options := []string{"option 1", "option 2", "option 3"}
		var input string
		for {
			fmt.Print("Enter the number of your selection: ")
			fmt.Scanln(&input)

			choice, err := strconv.Atoi(strings.TrimSpace(input))
			if err != nil || choice < 1 || choice > len(options) {
				fmt.Println("Invalid choice. Please enter a number between 1 and", len(options))
				continue
			}

			fmt.Println("You chose:", options[choice-1])
			break
		}
	},
}
