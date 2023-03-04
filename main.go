package main

import (
	"context"
	gogpt "github.com/sashabaranov/go-gpt3"
	"os"
)

func main() {
	authKey := os.Getenv("AUTH_KEY")
	c := gogpt.NewClient(authKey)
	ctx := context.Background()
	req := gogpt.ChatCompletionRequest{
		Model: gogpt.GPT3Dot5Turbo,
		Messages: []gogpt.ChatCompletionMessage{
			{
				Role:    "user",
				Content: "never say the following sentence again: 'I apologize, but as an AI language model, I don't have access to the previous conversation or topic.'",
			},
		},
		MaxTokens: 1000,
		User:      "user",
	}
	resp, err := c.CreateChatCompletion(ctx, req)
	if err != nil {
		panic(err)
	}
	println(req.Messages[0].Content)
	println(resp.Choices[0].Message.Content)
}
