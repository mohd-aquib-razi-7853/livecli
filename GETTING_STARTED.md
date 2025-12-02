# 🚀 Getting Started with LiveCLI

Welcome! You now have a complete AI-powered CLI application written in Go.

## ✅ What You Have

Your project includes:

### Core Application Files

- **`main.go`** - Main entry point
- **`cmd/root.go`** - CLI framework setup
- **`cmd/exec.go`** - System command execution
- **`cmd/chat.go`** - AI chat functionality
- **`cmd/ask.go`** - Quick AI questions
- **`cmd/interactive.go`** - Interactive mode (most powerful!)

### Build & Configuration

- **`go.mod`** - Go dependencies
- **`Makefile`** - Build automation
- **`setup.sh`** - Setup script
- **`test.sh`** - Test suite
- **`.env.example`** - Environment template
- **`.gitignore`** - Git ignore rules

### Documentation

- **`README.md`** - Complete documentation (8KB)
- **`QUICKSTART.md`** - 5-minute quick start
- **`DEMO.md`** - Detailed examples (6KB)
- **`EXAMPLES.md`** - Real-world use cases (10KB)
- **`PROJECT_SUMMARY.md`** - Technical overview

## 🎯 Quick Start (2 Minutes)

### 1. The binary is already built! Just set your API key:

```bash
export OPENAI_API_KEY="your-api-key-here"
```

Get your API key from: https://platform.openai.com/api-keys

### 2. Test it immediately:

```bash
# Execute a command
./livecli exec "ls -la"

# Ask AI a question (requires API key)
./livecli ask "How do I check disk space?"

# Start interactive mode (requires API key)
./livecli interactive
```

### 3. Try interactive mode:

```bash
./livecli interactive
```

Then type:

```
> /exec date
> @ask What is the current directory?
> /exec pwd
> Can you explain what pwd does?
> /exit
```

## 📖 Usage Examples

### Execute System Commands

```bash
./livecli exec "git status"
./livecli exec "ps aux | grep python"
./livecli exec --dir /tmp "ls"
```

### Ask AI Questions

```bash
./livecli ask "How do I find large files in Linux?"
./livecli ask "Explain Docker containers"
./livecli ask "Best practices for Go error handling"
```

### Start AI Chat Session

```bash
./livecli chat

# Then chat:
You> How do I use grep?
AI> [explains grep]
You> Give me 3 examples
AI> [provides examples]
You> /exit
```

### Interactive Mode (Recommended!)

```bash
./livecli interactive

# Mix commands and chat:
> /exec df -h                              # Execute command
> @ask How can I free up disk space?       # Quick question
> Can you explain the df output above?     # Chat with context
> /exec du -sh /* | sort -rh | head -5    # Execute another
> /clear                                   # Clear chat history
> /exit                                    # Exit
```

## 🔧 Available Commands

| Command                      | Description            |
| ---------------------------- | ---------------------- |
| `./livecli`                  | Show welcome screen    |
| `./livecli exec "<cmd>"`     | Execute system command |
| `./livecli chat`             | Start AI chat session  |
| `./livecli ask "<question>"` | Quick AI question      |
| `./livecli interactive`      | Combined mode (best!)  |
| `./livecli --help`           | Show help              |

## 🎨 Features

✅ **System Command Execution**

- Real-time output streaming
- Color-coded output
- Cross-platform support (Linux, macOS, Windows)

✅ **AI Integration**

- Chat with GPT-3.5-turbo or GPT-4
- Conversation history
- Customizable prompts

✅ **Interactive Mode**

- `/exec <cmd>` - Execute commands
- `@ask <question>` - Quick questions
- `<message>` - Chat with AI
- `/clear` - Clear history
- `/exit` - Exit

✅ **User Experience**

- Beautiful color output
- Readline support with history
- Clear error messages

## 🛠️ Advanced Usage

### Use Different AI Models

```bash
./livecli chat --model gpt-4
```

### Adjust AI Creativity

```bash
./livecli chat --temperature 0.9  # More creative
./livecli chat --temperature 0.3  # More focused
```

### Longer Responses

```bash
./livecli chat --max-tokens 2000
```

### Custom System Prompt

```bash
./livecli chat --system "You are a Linux expert"
```

## 📦 Installation (Optional)

Install globally to use from anywhere:

```bash
sudo mv livecli /usr/local/bin/

# Now use from anywhere:
cd ~
livecli exec "pwd"
```

## 🔄 Rebuilding

If you make changes to the code:

```bash
# Option 1: Go build
go build -o livecli main.go

# Option 2: Make
make build

# Option 3: Build for all platforms
make build-all
```

## 📚 Learn More

- **Full Documentation**: `cat README.md`
- **Quick Start**: `cat QUICKSTART.md`
- **Examples**: `cat EXAMPLES.md`
- **Demo Walkthrough**: `cat DEMO.md`

## 🧪 Testing

Run the test suite:

```bash
./test.sh
```

## 💡 Pro Tips

1. **Create aliases**:

   ```bash
   echo "alias li='./livecli interactive'" >> ~/.bashrc
   echo "alias lx='./livecli exec'" >> ~/.bashrc
   source ~/.bashrc
   ```

2. **Make API key permanent**:

   ```bash
   echo "export OPENAI_API_KEY='your-key'" >> ~/.bashrc
   source ~/.bashrc
   ```

3. **Use in scripts**:
   ```bash
   #!/bin/bash
   result=$(./livecli ask "How do I X?")
   echo "$result"
   ```

## 🎯 Real-World Workflows

### Workflow 1: Debugging

```bash
./livecli interactive
> /exec tail -50 /var/log/syslog
> I see these errors. What do they mean?
> @ask How do I fix this error?
> /exec sudo systemctl status myservice
```

### Workflow 2: Learning

```bash
./livecli interactive
> I want to learn about Docker
> @ask What is Docker?
> /exec docker --version
> How do I list containers?
> /exec docker ps
```

### Workflow 3: Development

```bash
./livecli interactive
> /exec git status
> I have uncommitted changes. Should I stash or commit?
> @ask What's a good commit message for adding authentication?
> /exec git add .
> /exec git commit -m "feat: add user authentication"
```

## 🐛 Troubleshooting

### "API key not set"

```bash
export OPENAI_API_KEY="sk-your-key-here"
```

### "Command not found"

```bash
# Use ./livecli (with ./) or install globally
sudo cp livecli /usr/local/bin/
```

### Build errors

```bash
go mod tidy
go clean
go build -o livecli main.go
```

## 🎉 You're All Set!

Start using LiveCLI now:

```bash
./livecli interactive
```

Have fun! 🚀

---

**Made with ❤️ and Go**

For more help: `./livecli --help` or read the full `README.md`
