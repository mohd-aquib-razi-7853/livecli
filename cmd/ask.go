package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/google/generative-ai-go/genai"
	"github.com/spf13/cobra"
	"google.golang.org/api/option"
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
		color.Red("Error: Gemini API key not set. Use --api-key flag or set GEMINI_API_KEY environment variable.")
		return
	}
	
	cyan := color.New(color.FgCyan, color.Bold)
	green := color.New(color.FgGreen, color.Bold)
	
	cyan.Printf("\n❓ Question: %s\n\n", question)
	
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		color.Red("Error creating Gemini client: %v\n", err)
		return
	}
	defer client.Close()

	geminiModel := client.GenerativeModel(model)
	geminiModel.SetTemperature(float32(temperature))
	geminiModel.SetMaxOutputTokens(int32(maxTokens))
	geminiModel.SystemInstruction = &genai.Content{
		Parts: []genai.Part{genai.Text("You are a helpful AI assistant specialized in programming, system administration, and command-line tools. Provide concise and accurate answers.")},
	}

	resp, err := geminiModel.GenerateContent(ctx, genai.Text(question))
	if err != nil {
		color.Red("Error: %v\n", err)
		return
	}
	
	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		color.Red("Error: No response from AI\n")
		return
	}
	
	green.Println("💡 Answer:")
	
	// Extract text from the response
	for _, part := range resp.Candidates[0].Content.Parts {
		if txt, ok := part.(genai.Text); ok {
			fmt.Println(string(txt))
		}
	}
	fmt.Println()
}
