package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	dotenv "github.com/joho/godotenv" //import a module with alias
	"github.com/openai/openai-go"
)

func main() {

	dotenv.Load()
	client := openai.NewClient()
	ctx := context.Background()

	// Read input from command prompt
	for {
		fmt.Print("\nEnter question > ")
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(query(ctx, client, line))
	}
}

func makeRequest(questions string) openai.ChatCompletionNewParams {
	maxToken, _ := strconv.Atoi(os.Getenv("MAX_TOKEN"))
	temperature, _ := strconv.ParseFloat(os.Getenv("TEMPERATURE"), 32)
	model := os.Getenv("GPT_MODEL")

	return openai.ChatCompletionNewParams{
		Model:       model,
		MaxTokens:   openai.Int(int64(maxToken)),
		Temperature: openai.Float(temperature),
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(questions),
		},
	}
}

func query(ctx context.Context, client openai.Client, question string) string {
	resp, err := client.Chat.Completions.New(ctx, makeRequest(question))
	if err != nil {
		log.Fatal("Error => ", err)
	}
	return resp.Choices[0].Message.Content
}
