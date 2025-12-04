package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/chzyer/readline"
	"github.com/fatih/color"
	openai "github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

var (
	systemPrompt string
	maxTokens    int
	temperature  float64
)

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Start an interactive AI chat session",
	Long: `Start an interactive chat session with AI assistant.
	
Type your messages and get AI responses. Type 'exit' or 'quit' to end the session.
Use '/clear' to clear conversation history.`,
	Run: func(cmd *cobra.Command, args []string) {
		startChatSession()
	},
}

func init() {
	rootCmd.AddCommand(chatCmd)
	
	chatCmd.Flags().StringVarP(&systemPrompt, "system", "s", "You are a helpful AI assistant specialized in programming, system administration, and command-line tools.", "System prompt for the AI")
	chatCmd.Flags().IntVarP(&maxTokens, "max-tokens", "t", 1000, "Maximum tokens in response")
	chatCmd.Flags().Float64VarP(&temperature, "temperature", "T", 0.7, "Temperature for AI responses (0.0-2.0)")
}

func startChatSession() {
	if apiKey == "" {
		color.Red("Error: OpenAI API key not set. Use --api-key flag or set OPENAI_API_KEY environment variable.")
		return
	}
	
	ctx := context.Background()
	client := openai.NewClient(apiKey)

	// Maintain conversation history
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: systemPrompt,
		},
	}
	
	cyan := color.New(color.FgCyan, color.Bold)
	green := color.New(color.FgGreen, color.Bold)
	yellow := color.New(color.FgYellow)
	
	cyan.Println("\n╔═══════════════════════════════════════════════════════════╗")
	cyan.Println("║           💬 AI Chat Session Started                      ║")
	cyan.Println("╚═══════════════════════════════════════════════════════════╝")
	yellow.Println("\nCommands: /clear (clear history), /exit or Ctrl+C (quit)")
	fmt.Printf("Model: %s\n\n", model)
	
	// Setup readline for better input handling
	rl, err := readline.New("You> ")
	if err != nil {
		color.Red("Error initializing readline: %v", err)
		return
	}
	defer rl.Close()
	
	for {
		line, err := rl.Readline()
		if err != nil {
			break
		}
		
		userInput := strings.TrimSpace(line)
		
		if userInput == "" {
			continue
		}
		
		// Handle special commands
		if userInput == "/exit" || userInput == "/quit" {
			green.Println("\n👋 Goodbye!")
			break
		}
		
		if userInput == "/clear" {
			messages = []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: systemPrompt,
				},
			}
			green.Println("✓ Conversation history cleared")
			continue
		}
		
		// Add user message to history
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: userInput,
		})
		
		// Get AI response
		fmt.Print("\nAI> ")
		response, err := getOpenAIResponse(ctx, client, messages)
		if err != nil {
			color.Red("Error: %v\n", err)
			// Remove the last user message if there was an error
			messages = messages[:len(messages)-1]
			continue
		}
		
		// Add assistant response to history
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: response,
		})
		
		fmt.Println(response)
		fmt.Println()
	}
}

func getOpenAIResponse(ctx context.Context, client *openai.Client, messages []openai.ChatCompletionMessage) (string, error) {
	resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:       model,
			Messages:    messages,
			Temperature: float32(temperature),
			MaxTokens:   maxTokens,
		},
	)
	if err != nil {
		return "", fmt.Errorf("chat error: %w", err)
	}
	
	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no response from AI")
	}
	
	return resp.Choices[0].Message.Content, nil
}
