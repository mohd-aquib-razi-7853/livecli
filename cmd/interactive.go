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

var interactiveCmd = &cobra.Command{
	Use:   "interactive",
	Short: "Start interactive mode (combines exec and chat)",
	Long: `Start interactive mode where you can execute commands or chat with AI.
	
Commands starting with '/' are executed as system commands.
Commands starting with '@' are sent to AI as questions.
Everything else is treated as a chat message.

Special commands:
  /exec <command>  - Execute a system command
  @ask <question>  - Ask AI a question
  /clear          - Clear chat history
  /exit or /quit  - Exit interactive mode`,
	Run: func(cmd *cobra.Command, args []string) {
		startInteractiveMode()
	},
}

func init() {
	rootCmd.AddCommand(interactiveCmd)
}

func startInteractiveMode() {
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
	magenta := color.New(color.FgMagenta, color.Bold)
	
	cyan.Println("\n╔═══════════════════════════════════════════════════════════╗")
	cyan.Println("║         🎮 Interactive Mode - LiveCLI                     ║")
	cyan.Println("╚═══════════════════════════════════════════════════════════╝")
	
	fmt.Println("\nMode Guide:")
	yellow.Println("  /exec <command>  → Execute system command")
	yellow.Println("  @ask <question>  → Ask AI a quick question")
	yellow.Println("  <message>        → Chat with AI")
	yellow.Println("  /clear           → Clear chat history")
	yellow.Println("  /exit            → Exit interactive mode")
	fmt.Println()
	
	// Setup readline
	rl, err := readline.New("> ")
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
		
		input := strings.TrimSpace(line)
		
		if input == "" {
			continue
		}
		
		// Handle exit
		if input == "/exit" || input == "/quit" {
			green.Println("\n👋 Exiting interactive mode. Goodbye!")
			break
		}
		
		// Handle clear
		if input == "/clear" {
			chatSession = geminiModel.StartChat()
			green.Println("✓ Chat history cleared")
			continue
		}
		
		// Handle command execution
		if strings.HasPrefix(input, "/exec ") {
			command := strings.TrimPrefix(input, "/exec ")
			executeCommand(command)
			continue
		}
		
		// Handle quick question
		if strings.HasPrefix(input, "@ask ") {
			question := strings.TrimPrefix(input, "@ask ")
			askQuestion(question)
			continue
		}
		
		// Handle AI chat
		magenta.Printf("\nYou: %s\n", input)
		
		fmt.Print("AI> ")
		response, err := getGeminiResponse(ctx, chatSession, input)
		if err != nil {
			color.Red("Error: %v\n", err)
			continue
		}
		
		fmt.Println(response)
		fmt.Println()
	}
}
