package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func GetResponse(client gpt3.Client, ctx context.Context, messages []gpt3.ChatCompletionRequestMessage) {
	err := client.ChatCompletionStream(ctx, gpt3.ChatCompletionRequest{
		Model:            "gpt-3.5-turbo",
		Messages:         messages,
		Temperature:      0.5,
		MaxTokens:        150,
		TopP:             1,
		N:                1,
		FrequencyPenalty: 0.0,
		PresencePenalty:  0.6,
		Stop:             []string{" Human:", " AI:"},
	}, func(resp *gpt3.ChatCompletionStreamResponse) {
		answer := resp.Choices[0].Delta.Content
		messages = append(messages, gpt3.ChatCompletionRequestMessage{
			Role:    "user",
			Content: answer})
		fmt.Print(answer)
	})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(13)
	}
	fmt.Printf("\n")
}

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	apiKey := viper.GetString("OPENAI_API_KEY")
	messages := []gpt3.ChatCompletionRequestMessage{{
		Role:    "system",
		Content: "You are a helpful assistant."}}

	if apiKey == "" {
		log.Fatalln("Missing API key. Please check OPEN_AI_KEY inside .env file")
	}
	ctx := context.Background()
	client := gpt3.NewClient(apiKey)
	rootCmd := &cobra.Command{
		Use:   "chatgpt",
		Short: "Chat with ChatGPT in console",
		Run: func(cmd *cobra.Command, args []string) {
			scanner := bufio.NewScanner(os.Stdin)
			quit := false

			for !quit {
				fmt.Print("Say something ('quit' to end):")
				if !scanner.Scan() {
					break
				}
				question := scanner.Text()
				switch question {
				case "quit":
					quit = true
				default:
					messages = append(messages, gpt3.ChatCompletionRequestMessage{
						Role:    "user",
						Content: question})
					GetResponse(client, ctx, messages)
				}
			}
		},
	}
	rootCmd.Execute()
}
