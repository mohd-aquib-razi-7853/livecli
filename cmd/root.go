package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	apiKey string
	model  string
)

var rootCmd = &cobra.Command{
	Use:   "livecli",
	Short: "LiveCLI - AI-powered command-line interface",
	Long: `LiveCLI is an intelligent CLI tool that combines system command execution 
with AI-powered chat assistance. Execute commands, get AI help, and boost your productivity.`,
	Run: func(cmd *cobra.Command, args []string) {
		displayWelcome()
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Global flags
	rootCmd.PersistentFlags().StringVar(&apiKey, "api-key", os.Getenv("OPENAI_API_KEY"), "OpenAI API key (or set OPENAI_API_KEY env var)")
	rootCmd.PersistentFlags().StringVarP(&model, "model", "m", "gpt-4o-mini", "AI model to use")
}

func displayWelcome() {
	cyan := color.New(color.FgCyan, color.Bold)
	yellow := color.New(color.FgYellow)
	
	cyan.Println("\n╔═══════════════════════════════════════════════════════════╗")
	cyan.Println("║            🚀 Welcome to LiveCLI v1.0                     ║")
	cyan.Println("╚═══════════════════════════════════════════════════════════╝")
	
	fmt.Println("\nAvailable Commands:")
	yellow.Println("  livecli exec <command>    - Execute system commands")
	yellow.Println("  livecli setup <task>      - AI-powered setup assistant")
	yellow.Println("  livecli git <message>     - Automate git add, commit, push (NEW!)")
	yellow.Println("  livecli chat              - Start AI chat session")
	yellow.Println("  livecli interactive       - Interactive mode (exec + chat)")
	yellow.Println("  livecli ask <question>    - Quick AI question")
	
	fmt.Println("\nExamples:")
	fmt.Println("  livecli exec \"ls -la\"")
	fmt.Println("  livecli setup \"rust into my system\"")
	fmt.Println("  livecli setup \"vscode editor\"")
	fmt.Println("  livecli chat")
	fmt.Println("  livecli ask \"How do I list all running processes?\"")
	fmt.Println("  livecli interactive")
	
	fmt.Println("\nUse 'livecli <command> --help' for more information about a command.")
	fmt.Println()
}
