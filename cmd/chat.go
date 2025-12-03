package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/chzyer/readline"
	"github.com/fatih/color"
	"github.com/google/generative-ai-go/genai"
	"github.com/spf13/cobra"
	"google.golang.org/api/option"
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
		color.Red("Error: Gemini API key not set. Use --api-key flag or set GEMINI_API_KEY environment variable.")
		return
	}
	
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		color.Red("Error creating Gemini client: %v", err)
		return
	}
	defer client.Close()

	geminiModel := client.GenerativeModel(model)
	geminiModel.SetTemperature(float32(temperature))
	geminiModel.SetMaxOutputTokens(int32(maxTokens))
	geminiModel.SystemInstruction = &genai.Content{
		Parts: []genai.Part{genai.Text(systemPrompt)},
	}

	chatSession := geminiModel.StartChat()
	
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
			chatSession = geminiModel.StartChat()
			green.Println("✓ Conversation history cleared")
			continue
		}
		
		// Get AI response
		fmt.Print("\nAI> ")
		response, err := getGeminiResponse(ctx, chatSession, userInput)
		if err != nil {
			color.Red("Error: %v\n", err)
			continue
		}
		
		fmt.Println(response)
		fmt.Println()
	}
}

func getGeminiResponse(ctx context.Context, chat *genai.ChatSession, message string) (string, error) {
	resp, err := chat.SendMessage(ctx, genai.Text(message))
	if err != nil {
		return "", fmt.Errorf("chat error: %w", err)
	}
	
	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("no response from AI")
	}
	
	// Extract text from the response
	var text string
	for _, part := range resp.Candidates[0].Content.Parts {
		if txt, ok := part.(genai.Text); ok {
			text += string(txt)
		}
	}
	
	return text, nil
}
