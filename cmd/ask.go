package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/fatih/color"
	openai "github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

var askCmd = &cobra.Command{
	Use:   "ask [question]",
	Short: "Ask a quick question to AI",
	Long: `Ask a single question to the AI and get an immediate response.
	
Examples:
  livecli ask "How do I list all running processes?"
  livecli ask "Explain what 'grep' command does"`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		question := strings.Join(args, " ")
		askQuestion(question)
	},
}

func init() {
	rootCmd.AddCommand(askCmd)
}

func askQuestion(question string) {
	if apiKey == "" {
		color.Red("Error: OpenAI API key not set. Use --api-key flag or set OPENAI_API_KEY environment variable.")
		return
	}
	
	cyan := color.New(color.FgCyan, color.Bold)
	green := color.New(color.FgGreen, color.Bold)
	
	cyan.Printf("\n❓ Question: %s\n\n", question)
	
	ctx := context.Background()
	client := openai.NewClient(apiKey)

	resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "You are a helpful AI assistant specialized in programming, system administration, and command-line tools. Provide concise and accurate answers.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: question,
				},
			},
			Temperature: float32(temperature),
			MaxTokens:   maxTokens,
		},
	)
	if err != nil {
		color.Red("Error: %v\n", err)
		return
	}
	
	if len(resp.Choices) == 0 {
		color.Red("Error: No response from AI\n")
		return
	}
	
	green.Println("💡 Answer:")
	fmt.Println(resp.Choices[0].Message.Content)
	fmt.Println()
}
