package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/shaikzhafir/chad-gpt/clientService"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "myapp",
		Short: "MyApp is a CLI application",
		Long:  `MyApp is a simple CLI application built using Cobra.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("welcome to chad cli gpt")
			prompt := promptui.Select{
				Label: "Select Day",
				Items: []string{"New chat", "Previous chat", "Exit"},
			}

			_, result, err := prompt.Run()
			switch result {
			case "New chat":
				fmt.Println("New chat selected")
				StartChat(context.Background())
			case "Previous chat":
				fmt.Println("Previous chat selected")
				prompt2 := promptui.Select{
					Label: "Select Day",
					Items: []string{"chat 1", "chat 2", "chat 3"},
				}
				_, result, err := prompt2.Run()
				switch result {
				case "chat 1":
					fmt.Println("New chat selected")
					StartChat(context.Background())
				case "chat 2":
					fmt.Println("Previous chat selected")
				case "chat 3":
					fmt.Println("Exit")
					return
				case "Exit":
					fmt.Println("Exit")
					return
				}
				if err != nil {
					fmt.Printf("Prompt failed %v\n", err)
					return
				}

			case "Exit":
				fmt.Println("Exit")
				return
			}
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

			fmt.Printf("You choose %q\n", result)
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}

	/*	c := clientService.NewClientService()
		var promptChan = make(chan string)

		scanner := bufio.NewScanner(os.Stdin)
		go func() {
			for scanner.Scan() {
				promptChan <- scanner.Text()
			}
		}()

		for {
			select {
			case prompt := <-promptChan:
				// append prompt to chat completion request
				// and send it to the stream
				c.SendPromptToStream(ctx, prompt)
			}
		}*/
}

func StartChat(ctx context.Context) {
	c := clientService.NewClientService()
	var promptChan = make(chan string)
	reader := bufio.NewReader(os.Stdin)

	go func() {
		for {
			text, _ := reader.ReadString('\n')
			promptChan <- text
		}
	}()

	for {
		select {
		case prompt := <-promptChan:
			if prompt == "\n" { // skip empty lines
				continue
			}
			c.SendPromptToStream(ctx, prompt)
		}
	}
}
