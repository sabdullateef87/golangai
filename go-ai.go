package main

import (
	"context"
	"fmt"
	"github.com/openai/openai-go"
	dotenv "github.com/joho/godotenv" // import a module with alias

)

func main() {

	dotenv.Load()
	client := openai.NewClient()
	resp, err := client.Chat.Completions.New(context.Background(), openai.ChatCompletionNewParams{
		Model: "gpt-4o-mini",
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage("How many coffees should a man drink per day?"),
		},
	})

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
}