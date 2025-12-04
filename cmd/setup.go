package cmd

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/fatih/color"
	openai "github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

type SetupStep struct {
	Command     string `json:"command"`
	Description string `json:"description"`
	Optional    bool   `json:"optional"`
}

type SetupPlan struct {
	Steps []SetupStep `json:"steps"`
}

var (
	autoConfirm bool
	dryRun      bool
)

var setupCmd = &cobra.Command{
	Use:   "setup [task description]",
	Short: "AI-powered setup and installation assistant",
	Long: `Use AI to generate and execute installation/setup commands.
	
The AI will analyze your request, generate necessary commands, explain them,
and ask for confirmation before executing each step.

Examples:
  livecli setup "rust into my system"
  livecli setup "docker and docker-compose"
  livecli setup "nodejs version 18"
  livecli setup "vscode editor"`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		executeSetup(task)
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
	
	setupCmd.Flags().BoolVarP(&autoConfirm, "yes", "y", false, "Auto-confirm all commands")
	setupCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Show commands without executing")
}

func executeSetup(task string) {
	if apiKey == "" {
		color.Red("Error: OpenAI API key not set. Use --api-key flag or set OPENAI_API_KEY environment variable.")
		return
	}
	
	cyan := color.New(color.FgCyan, color.Bold)
	green := color.New(color.FgGreen, color.Bold)
	yellow := color.New(color.FgYellow, color.Bold)
	magenta := color.New(color.FgMagenta, color.Bold)
	
	cyan.Println("\n╔═══════════════════════════════════════════════════════════╗")
	cyan.Println("║           🤖 AI Setup Assistant                           ║")
	cyan.Println("╚═══════════════════════════════════════════════════════════╝")
	
	fmt.Printf("\n📋 Task: %s\n", task)
	yellow.Println("\n⏳ Analyzing your request and generating setup plan...")
	
	// Get setup plan from AI
	plan, err := generateSetupPlan(task)
	if err != nil {
		color.Red("\n❌ Error generating setup plan: %v", err)
		return
	}
	
	if len(plan.Steps) == 0 {
		color.Yellow("\n⚠️  No setup steps were generated. The task might already be complete or unclear.")
		return
	}
	
	// Display the plan
	cyan.Println("\n📝 Setup Plan:")
	cyan.Println("─────────────────────────────────────────────────────────────")
	
	for i, step := range plan.Steps {
		if step.Optional {
			yellow.Printf("\n%d. [OPTIONAL] %s\n", i+1, step.Description)
		} else {
			fmt.Printf("\n%d. %s\n", i+1, step.Description)
		}
		magenta.Printf("   Command: %s\n", step.Command)
	}
	
	cyan.Println("\n─────────────────────────────────────────────────────────────")
	
	// Ask for overall confirmation
	if !autoConfirm && !dryRun {
		fmt.Print("\n❓ Do you want to proceed with this setup plan? (yes/no): ")
		reader := bufio.NewReader(os.Stdin)
		response, _ := reader.ReadString('\n')
		response = strings.TrimSpace(strings.ToLower(response))
		
		if response != "yes" && response != "y" {
			yellow.Println("\n❌ Setup cancelled by user.")
			return
		}
	}
	
	if dryRun {
		green.Println("\n✓ Dry run complete. No commands were executed.")
		return
	}
	
	// Execute each step
	green.Println("\n\n🚀 Starting setup process...")
	
	for i, step := range plan.Steps {
		cyan.Printf("\n\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
		cyan.Printf("Step %d/%d\n", i+1, len(plan.Steps))
		cyan.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		
		fmt.Printf("\n📌 %s\n", step.Description)
		magenta.Printf("💻 Command: %s\n", step.Command)
		
		// Ask for confirmation for each step
		if !autoConfirm {
			if step.Optional {
				fmt.Print("\n❓ Execute this optional step? (yes/no/skip): ")
			} else {
				fmt.Print("\n❓ Execute this command? (yes/no/skip all): ")
			}
			
			reader := bufio.NewReader(os.Stdin)
			response, _ := reader.ReadString('\n')
			response = strings.TrimSpace(strings.ToLower(response))
			
			if response == "no" || response == "n" {
				yellow.Println("⏭️  Skipped by user")
				continue
			}
			
			if response == "skip all" || response == "skip" {
				yellow.Println("\n⏭️  Skipping remaining steps")
				break
			}
		}
		
		// Execute the command
		green.Printf("\n▶ Executing...\n")
		cyan.Println("─────────────────────────────────────────────────────────────")
		
		executeCommand(step.Command)
		
		green.Printf("\n✓ Step %d/%d completed\n", i+1, len(plan.Steps))
	}
	
	// Final summary
	cyan.Println("\n\n╔═══════════════════════════════════════════════════════════╗")
	green.Println("║           ✅ Setup Complete!                              ║")
	cyan.Println("╚═══════════════════════════════════════════════════════════╝")
	
	fmt.Printf("\n✓ Executed %d steps successfully\n", len(plan.Steps))
	green.Println("\n💡 Tip: Verify the installation with relevant commands (e.g., version checks)")
	fmt.Println()
}

func generateSetupPlan(task string) (SetupPlan, error) {
	ctx := context.Background()
	client := openai.NewClient(apiKey)
	
	// Detect OS
	osInfo := detectOS()
	
	systemPrompt := fmt.Sprintf(`You are an expert system administrator and DevOps engineer. Generate a precise, safe setup plan for the user's request.

Operating System: %s
Task: %s

IMPORTANT RULES:
1. Generate ONLY the necessary commands for THIS specific OS
2. Use the system's package manager (apt, dnf, yum, brew, etc.)
3. Each command should be safe and commonly used
4. Include verification commands when helpful
5. Mark optional steps (like adding to PATH if it's automatic)
6. Keep commands simple and atomic (one logical action per command)
7. Include sudo only when absolutely necessary
8. For URLs/downloads, use official sources only

Respond with ONLY a valid JSON object in this EXACT format (no markdown, no explanation):
{
  "steps": [
    {
      "command": "the exact command to run",
      "description": "brief description of what this does",
      "optional": false
    }
  ]
}

Example for "install docker":
{
  "steps": [
    {"command": "sudo apt update", "description": "Update package index", "optional": false},
    {"command": "sudo apt install -y docker.io", "description": "Install Docker", "optional": false},
    {"command": "sudo systemctl start docker", "description": "Start Docker service", "optional": false},
    {"command": "sudo systemctl enable docker", "description": "Enable Docker on boot", "optional": false},
    {"command": "sudo usermod -aG docker $USER", "description": "Add user to docker group", "optional": true},
    {"command": "docker --version", "description": "Verify Docker installation", "optional": false}
  ]
}`, osInfo, task)

	resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: systemPrompt,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: fmt.Sprintf("Generate setup commands for: %s", task),
				},
			},
			Temperature: 0.3, // Lower temperature for more consistent output
			MaxTokens:   2000,
		},
	)
	if err != nil {
		return SetupPlan{}, fmt.Errorf("AI request failed: %w", err)
	}
	
	if len(resp.Choices) == 0 {
		return SetupPlan{}, fmt.Errorf("no response from AI")
	}
	
	content := resp.Choices[0].Message.Content
	
	// Clean up the response (remove markdown code blocks if present)
	content = strings.TrimSpace(content)
	content = strings.TrimPrefix(content, "```json")
	content = strings.TrimPrefix(content, "```")
	content = strings.TrimSuffix(content, "```")
	content = strings.TrimSpace(content)
	
	// Parse JSON response
	var plan SetupPlan
	if err := json.Unmarshal([]byte(content), &plan); err != nil {
		// If JSON parsing fails, show what we got
		fmt.Printf("Debug - AI Response:\n%s\n", content)
		return SetupPlan{}, fmt.Errorf("failed to parse AI response as JSON: %w", err)
	}
	
	return plan, nil
}

func detectOS() string {
	switch runtime.GOOS {
	case "windows":
		return "Windows"
	case "darwin":
		return "macOS"
	case "linux":
		// Try to read /etc/os-release for more specific Linux distro info
		data, err := os.ReadFile("/etc/os-release")
		if err == nil {
			content := string(data)
			if strings.Contains(content, "Ubuntu") {
				return "Ubuntu Linux (apt)"
			}
			if strings.Contains(content, "Debian") {
				return "Debian Linux (apt)"
			}
			if strings.Contains(content, "Fedora") {
				return "Fedora Linux (dnf)"
			}
			if strings.Contains(content, "CentOS") || strings.Contains(content, "Red Hat") {
				return "CentOS/RHEL (yum/dnf)"
			}
			if strings.Contains(content, "Arch") {
				return "Arch Linux (pacman)"
			}
		}
		return "Linux (generic)"
	default:
		return fmt.Sprintf("%s (generic)", runtime.GOOS)
	}
}
