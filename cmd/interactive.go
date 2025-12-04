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
			messages = []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: systemPrompt,
				},
			}
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
		
		// Add user message to history
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: input,
		})
		
		fmt.Print("AI> ")
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
