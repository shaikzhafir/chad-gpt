package clientService

import (
	"context"
	"errors"
	"fmt"
	gogpt "github.com/sashabaranov/go-gpt3"
	"io"
	"log"
	"os"
)

type clientService struct {
	client         *gogpt.Client
	chatMessageArr []gogpt.ChatCompletionMessage
}

type ClientService interface {
	SendPromptToStream(ctx context.Context, prompt string) error
}

func NewClientService() ClientService {
	authKey := os.Getenv("AUTH_KEY")
	c := gogpt.NewClient(authKey)
	if c == nil {
		log.Fatal(errors.New("no auth key added"))
	}

	clientService := &clientService{
		client:         c,
		chatMessageArr: make([]gogpt.ChatCompletionMessage, 0),
	}
	return clientService
}

func (c *clientService) SendPromptToStream(ctx context.Context, prompt string) error {
	c.chatMessageArr = append(c.chatMessageArr, gogpt.ChatCompletionMessage{
		Content: prompt,
		Role:    "user",
	})

	req := gogpt.ChatCompletionRequest{
		Model:     gogpt.GPT3Dot5Turbo,
		MaxTokens: 1000,
		Messages:  c.chatMessageArr,
		Stream:    true,
		User:      "user",
	}

	stream, err := c.client.CreateChatCompletionStream(ctx, req)
	if err != nil {
		return err
	}
	var storeResponse string
	defer stream.Close()
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Printf("\nStream finished\n")
			c.chatMessageArr = append(c.chatMessageArr, gogpt.ChatCompletionMessage{
				Content: storeResponse,
				Role:    "assistant",
			})
			return nil
		}

		if err != nil {
			fmt.Printf("Stream error: %v\n", err)
			return err
		}
		fmt.Printf(response.Choices[0].Delta.Content)
		if response.Choices[0].Delta.Content != "\n" {
			storeResponse += response.Choices[0].Delta.Content
		}
	}
}
