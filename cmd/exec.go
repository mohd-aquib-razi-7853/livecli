package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	execShell string
	execDir   string
)

var execCmd = &cobra.Command{
	Use:   "exec [command]",
	Short: "Execute system commands",
	Long: `Execute system commands with real-time output.
	
Examples:
  livecli exec "ls -la"
  livecli exec "git status"
  livecli exec --shell bash "echo Hello"`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		command := strings.Join(args, " ")
		executeCommand(command)
	},
}

func init() {
	rootCmd.AddCommand(execCmd)
	
	defaultShell := "sh"
	if runtime.GOOS == "windows" {
		defaultShell = "cmd"
	}
	
	execCmd.Flags().StringVarP(&execShell, "shell", "s", defaultShell, "Shell to use for execution")
	execCmd.Flags().StringVarP(&execDir, "dir", "d", ".", "Working directory for command execution")
}

func executeCommand(command string) {
	green := color.New(color.FgGreen, color.Bold)
	red := color.New(color.FgRed, color.Bold)
	cyan := color.New(color.FgCyan)
	
	green.Printf("\n▶ Executing: %s\n", command)
	cyan.Println("─────────────────────────────────────────────────────────────")
	
	var cmd *exec.Cmd
	
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", command)
	} else {
		cmd = exec.Command(execShell, "-c", command)
	}
	
	// Set working directory
	cmd.Dir = execDir
	
	// Create pipes for stdout and stderr
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		red.Printf("Error creating stdout pipe: %v\n", err)
		return
	}
	
	stderr, err := cmd.StderrPipe()
	if err != nil {
		red.Printf("Error creating stderr pipe: %v\n", err)
		return
	}
	
	// Start the command
	if err := cmd.Start(); err != nil {
		red.Printf("Error starting command: %v\n", err)
		return
	}
	
	// Read stdout in real-time
	go streamOutput(stdout, false)
	
	// Read stderr in real-time
	go streamOutput(stderr, true)
	
	// Wait for command to complete
	if err := cmd.Wait(); err != nil {
		red.Printf("\n✗ Command failed with error: %v\n", err)
		cyan.Println("─────────────────────────────────────────────────────────────")
		return
	}
	
	green.Println("\n✓ Command completed successfully")
	cyan.Println("─────────────────────────────────────────────────────────────\n")
}

func streamOutput(reader io.Reader, isError bool) {
	scanner := bufio.NewScanner(reader)
	red := color.New(color.FgRed)
	
	for scanner.Scan() {
		line := scanner.Text()
		if isError {
			red.Println(line)
		} else {
			fmt.Println(line)
		}
	}
	
	if err := scanner.Err(); err != nil && err != io.EOF {
		fmt.Fprintf(os.Stderr, "Error reading output: %v\n", err)
	}
}
