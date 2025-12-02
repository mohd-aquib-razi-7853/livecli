package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	gitAutoConfirm bool
)

var gitCmd = &cobra.Command{
	Use:   "git [commit message]",
	Short: "Automate git add, commit, and push",
	Long: `Automate your git workflow in one command.
	
This command performs the following actions:
1. Stages all changes (git add .)
2. Commits with your message (git commit -m "message")
3. Pushes to the current branch (git push)

You will be asked for confirmation before execution.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		message := strings.Join(args, " ")
		runGitWorkflow(message)
	},
}

func init() {
	rootCmd.AddCommand(gitCmd)
	gitCmd.Flags().BoolVarP(&gitAutoConfirm, "yes", "y", false, "Auto-confirm all git actions")
}

func runGitWorkflow(message string) {
	cyan := color.New(color.FgCyan, color.Bold)
	green := color.New(color.FgGreen, color.Bold)
	yellow := color.New(color.FgYellow)
	red := color.New(color.FgRed, color.Bold)

	cyan.Println("\n╔═══════════════════════════════════════════════════════════╗")
	cyan.Println("║              🚀 Git Workflow Automator                    ║")
	cyan.Println("╚═══════════════════════════════════════════════════════════╝")

	// Define the steps
	steps := []struct {
		desc string
		cmd  string
	}{
		{"Stage all changes", "git add ."},
		{"Commit changes", fmt.Sprintf("git commit -m \"%s\"", message)},
		{"Push to remote", "git push"},
	}

	// Display plan
	fmt.Printf("\n📋 Commit Message: %s\n", message)
	cyan.Println("\n📝 Execution Plan:")
	cyan.Println("─────────────────────────────────────────────────────────────")
	
	for i, step := range steps {
		fmt.Printf("\n%d. %s\n", i+1, step.desc)
		color.Magenta("   Command: %s", step.cmd)
	}
	cyan.Println("\n─────────────────────────────────────────────────────────────")

	// Confirmation
	if !gitAutoConfirm {
		fmt.Print("\n❓ Proceed with these git operations? (yes/no): ")
		reader := bufio.NewReader(os.Stdin)
		response, _ := reader.ReadString('\n')
		response = strings.TrimSpace(strings.ToLower(response))

		if response != "yes" && response != "y" {
			yellow.Println("\n❌ Operation cancelled.")
			return
		}
	}

	green.Println("\n🚀 Executing git workflow...")

	// Execute steps sequentially
	for i, step := range steps {
		cyan.Printf("\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
		cyan.Printf("Step %d/%d: %s\n", i+1, len(steps), step.desc)
		cyan.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

		// We use a custom execution here to ensure we stop on error
		err := runCommandWithOutput(step.cmd)
		if err != nil {
			red.Printf("\n❌ Step failed: %v\n", err)
			red.Println("Stopping workflow execution.")
			return
		}
		green.Printf("\n✓ %s completed\n", step.desc)
	}

	cyan.Println("\n╔═══════════════════════════════════════════════════════════╗")
	green.Println("║           ✅ Git Workflow Complete!                       ║")
	cyan.Println("╚═══════════════════════════════════════════════════════════╝")
	fmt.Println()
}

// runCommandWithOutput runs a command and streams output, returning error if it fails
func runCommandWithOutput(commandStr string) error {
	// Determine shell
	shell := "sh"
	if os.Getenv("SHELL") != "" {
		shell = os.Getenv("SHELL")
	}
	
	var cmd *exec.Cmd
	// Simple command splitting for basic cases, but using shell for complex ones
	cmd = exec.Command(shell, "-c", commandStr)
	
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	return cmd.Run()
}
