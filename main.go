package main

import (
	"bufio"
	"chat-backend/clientService"
	"context"
	"os"
)

func main() {
	ctx := context.Background()
	c := clientService.NewClientService()
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
	}
}
