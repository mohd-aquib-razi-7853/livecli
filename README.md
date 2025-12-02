# LiveCLI 🚀

A powerful AI-powered command-line interface that combines system command execution with intelligent AI chat assistance.

## Features ✨

- **🤖 AI-Powered Setup**: Intelligent installation assistant - just describe what you want to install!
- **🚀 Git Automation**: One command to add, commit, and push your changes (NEW!)
- **Command Execution**: Execute system commands with real-time output streaming
- **AI Chat**: Interactive chat sessions with AI assistant
- **Quick Questions**: Ask one-off questions without starting a full chat session
- **Interactive Mode**: Unified interface combining command execution and AI chat
- **Cross-Platform**: Works on Linux, macOS, and Windows
- **Colorized Output**: Beautiful terminal output with color-coded messages
- **Command History**: Full readline support with history

## Installation 📦

### Prerequisites

- Go 1.21 or higher
- OpenAI API key

### Build from Source

```bash
# Clone the repository
git clone https://github.com/mrazi/livecli.git
cd livecli

# Download dependencies
go mod download

# Build the binary
go build -o livecli

# Optional: Install globally
sudo mv livecli /usr/local/bin/
```

## Configuration ⚙️

### Set your OpenAI API Key

You can set your API key in three ways:

1. **Environment Variable** (recommended):

```bash
export OPENAI_API_KEY="your-api-key-here"
```

2. **Command-line Flag**:

```bash
livecli chat --api-key="your-api-key-here"
```

3. **.env File**:

```bash
cp .env.example .env
# Edit .env and add your API key
```

## Usage 🎯

### Display Help

```bash
livecli
livecli --help
```

### Execute System Commands

```bash
# Execute a simple command
livecli exec "ls -la"

# Execute with custom shell
livecli exec --shell bash "echo Hello World"

# Execute in specific directory
livecli exec --dir /tmp "pwd"
```

### AI-Powered Setup Assistant 🤖 (NEW!)

Let AI figure out and execute installation commands for you!

```bash
# Install software with AI assistance
livecli setup "rust into my system"
livecli setup "vscode editor"
livecli setup "docker and docker-compose"
livecli setup "nodejs version 18"

# Auto-confirm all steps
livecli setup "python3 and pip" --yes

# Preview commands without executing
livecli setup "postgresql database" --dry-run
```

**How it works**:

1. Describe what you want to install
2. AI generates step-by-step installation commands
3. Each command is explained and shown to you
4. You confirm before execution
5. Commands execute one by one

**Example**:

```bash
$ livecli setup "rust into my system"

╔═══════════════════════════════════════════════════════════╗
║           🤖 AI Setup Assistant                           ║
╚═══════════════════════════════════════════════════════════╝

📋 Task: rust into my system

⏳ Analyzing your request and generating setup plan...

📝 Setup Plan:
─────────────────────────────────────────────────────────────

1. Install Rust using rustup
   Command: curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y

2. Source cargo environment
   Command: source $HOME/.cargo/env

3. Verify Rust installation
   Command: rustc --version

❓ Do you want to proceed with this setup plan? (yes/no): yes

🚀 Starting setup process...
[Executes each command with confirmation]
```

See `SETUP_GUIDE.md` for more examples!

### Git Workflow Automation 🚀 (NEW!)

Streamline your git workflow with a single command!

```bash
# Stage, commit, and push in one go
livecli git "feat: add new login page"

# Auto-confirm (great for scripts)
livecli git "fix: typo in readme" --yes
```

**What it does**:

1. `git add .` (Stages all changes)
2. `git commit -m "message"` (Commits with your message)
3. `git push` (Pushes to current branch)

**Safety**: It shows you the plan and asks for confirmation before running!

### AI Chat Session

```bash
# Start interactive chat
livecli chat

# Use custom model
livecli chat --model gpt-4

# Adjust temperature and max tokens
livecli chat --temperature 0.9 --max-tokens 2000
```

**Chat Commands**:

- Type your message and press Enter
- `/clear` - Clear conversation history
- `/exit` or `/quit` - Exit chat session

### Quick Questions

```bash
# Ask a single question
livecli ask "How do I list all running processes?"

# Get command help
livecli ask "Explain what grep command does"

# Programming questions
livecli ask "How to reverse a string in Go?"
```

### Interactive Mode

The most powerful mode - combines everything!

```bash
livecli interactive
```

**Interactive Mode Commands**:

- `/exec <command>` - Execute a system command
- `@ask <question>` - Ask AI a quick question
- `<message>` - Chat with AI (maintains conversation history)
- `/clear` - Clear chat history
- `/exit` - Exit interactive mode

**Examples**:

```
> /exec ls -la
> @ask How do I find large files?
> Can you help me debug this error?
```

## Examples 📚

### Example 1: Command Execution

```bash
$ livecli exec "git status"

▶ Executing: git status
─────────────────────────────────────────────────────────────
On branch main
Your branch is up to date with 'origin/main'.

nothing to commit, working tree clean

✓ Command completed successfully
─────────────────────────────────────────────────────────────
```

### Example 2: AI Chat

```bash
$ livecli chat

╔═══════════════════════════════════════════════════════════╗
║           💬 AI Chat Session Started                      ║
╚═══════════════════════════════════════════════════════════╝

Commands: /clear (clear history), /exit or Ctrl+C (quit)
Model: gpt-3.5-turbo

You> How do I find all files larger than 100MB?

AI> You can use the `find` command with the `-size` option:

find /path/to/search -type f -size +100M

This will:
- Search in `/path/to/search`
- Look for files only (`-type f`)
- Find files larger than 100MB (`-size +100M`)

You can also add `-exec ls -lh {} \;` to see file sizes in human-readable format.
```

### Example 3: Interactive Mode

```bash
$ livecli interactive

╔═══════════════════════════════════════════════════════════╗
║         🎮 Interactive Mode - LiveCLI                     ║
╚═══════════════════════════════════════════════════════════╝

Mode Guide:
  /exec <command>  → Execute system command
  @ask <question>  → Ask AI a quick question
  <message>        → Chat with AI
  /clear           → Clear chat history
  /exit            → Exit interactive mode

> /exec df -h

▶ Executing: df -h
─────────────────────────────────────────────────────────────
Filesystem      Size  Used Avail Use% Mounted on
/dev/sda1       50G   30G   20G  60% /
...

> @ask How can I free up disk space?

💡 Answer:
Here are some ways to free up disk space:
1. Remove old logs: sudo find /var/log -type f -name "*.log" -delete
2. Clean package manager cache: sudo apt clean (Ubuntu/Debian)
3. Remove unused Docker images: docker system prune -a
...

> Thanks! Can you show me more advanced cleanup techniques?

AI> Of course! Here are more advanced techniques:
...
```

## Command Reference 📖

### Global Flags

- `--api-key`: OpenAI API key
- `--model, -m`: AI model to use (default: gpt-3.5-turbo)

### exec Command

```bash
livecli exec [flags] <command>
```

**Flags**:

- `--shell, -s`: Shell to use (default: sh on Unix, cmd on Windows)
- `--dir, -d`: Working directory (default: current directory)

### setup Command

```bash
livecli setup [flags] <task description>
```

**Flags**:

- `--yes, -y`: Auto-confirm all commands
- `--dry-run`: Show commands without executing

### git Command

```bash
livecli git [flags] <message>
```

**Flags**:

- `--yes, -y`: Auto-confirm all actions

### chat Command

```bash
livecli chat [flags]
```

**Flags**:

- `--system, -s`: System prompt for AI
- `--max-tokens, -t`: Maximum tokens in response (default: 1000)
- `--temperature, -T`: Temperature for AI responses (default: 0.7)

### ask Command

```bash
livecli ask [question]
```

### interactive Command

```bash
livecli interactive [flags]
```

Uses the same flags as the chat command.

## Development 🛠️

### Project Structure

```
livecli/
├── main.go              # Entry point
├── cmd/                 # Command implementations
│   ├── root.go         # Root command and CLI setup
│   ├── exec.go         # Command execution
│   ├── setup.go        # AI-powered setup assistant
│   ├── git.go          # Git workflow automation (NEW!)
│   ├── chat.go         # AI chat session
│   ├── ask.go          # Quick questions
│   └── interactive.go  # Interactive mode
├── go.mod              # Go dependencies
├── go.sum              # Dependency checksums
├── README.md           # This file
└── .env.example        # Example environment config
```

### Building

```bash
# Build for current platform
go build -o livecli

# Build for specific platform
GOOS=linux GOARCH=amd64 go build -o livecli-linux
GOOS=darwin GOARCH=amd64 go build -o livecli-macos
GOOS=windows GOARCH=amd64 go build -o livecli.exe
```

### Testing

```bash
# Run tests
go test ./...

# Run tests with coverage
go test -cover ./...
```

## Technologies Used 🔧

- **[Cobra](https://github.com/spf13/cobra)**: CLI framework
- **[go-openai](https://github.com/sashabaranov/go-openai)**: OpenAI API client
- **[color](https://github.com/fatih/color)**: Colorized terminal output
- **[readline](https://github.com/chzyer/readline)**: Interactive input with history

## Contributing 🤝

Contributions are welcome! Please feel free to submit a Pull Request.

## License 📄

MIT License - feel free to use this project for any purpose.

## Support 💬

If you have questions or need help:

1. Check the [Examples](#examples-) section
2. Open an issue on GitHub
3. Read the command help: `livecli <command> --help`

## Roadmap 🗺️

- [ ] Add command suggestions based on AI analysis
- [ ] Save conversation history to file
- [ ] Support for multiple AI providers (Anthropic, etc.)
- [ ] Plugin system for custom commands
- [ ] Auto-completion for commands
- [ ] Web UI for chat interface

---

Made with ❤️ by the LiveCLI team
