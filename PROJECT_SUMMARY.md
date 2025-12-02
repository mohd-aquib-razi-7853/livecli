# LiveCLI - Project Summary

## 📋 Overview

**LiveCLI** is a complete AI-powered command-line interface application built in Go that combines:

- System command execution with real-time output streaming
- AI chat functionality using OpenAI's GPT models
- Interactive mode that seamlessly blends both capabilities

## 🏗️ Architecture

### Project Structure

```
livecli/
├── main.go                 # Application entry point
├── cmd/                    # Command implementations
│   ├── root.go            # Root command & CLI setup (Cobra)
│   ├── exec.go            # System command execution
│   ├── chat.go            # AI chat session
│   ├── ask.go             # Quick AI questions
│   └── interactive.go     # Combined interactive mode
├── go.mod                 # Go module dependencies
├── go.sum                 # Dependency checksums
├── Makefile              # Build automation
├── setup.sh              # Setup script
├── README.md             # Complete documentation
├── QUICKSTART.md         # 5-minute getting started guide
├── DEMO.md               # Detailed usage examples
└── .env.example          # Environment configuration template
```

### Technology Stack

| Component       | Technology                                             | Purpose                        |
| --------------- | ------------------------------------------------------ | ------------------------------ |
| CLI Framework   | [Cobra](https://github.com/spf13/cobra)                | Command structure & parsing    |
| AI Integration  | [go-openai](https://github.com/sashabaranov/go-openai) | OpenAI API client              |
| Terminal Colors | [color](https://github.com/fatih/color)                | Colorized output               |
| Input Handling  | [readline](https://github.com/chzyer/readline)         | Interactive input with history |
| Language        | Go 1.21+                                               | Core implementation            |

## 🎯 Features

### 1. Command Execution (`livecli exec`)

- Real-time output streaming (stdout & stderr)
- Cross-platform shell support (sh/bash on Unix, cmd on Windows)
- Custom working directory support
- Color-coded output (errors in red, success in green)

### 2. AI Chat (`livecli chat`)

- Interactive chat sessions with conversation history
- Customizable system prompts
- Adjustable temperature and max tokens
- Commands: `/clear`, `/exit`

### 3. Quick Questions (`livecli ask`)

- One-off AI queries without full chat session
- Fast responses for simple questions
- Perfect for command-line tips and quick help

### 4. Interactive Mode (`livecli interactive`)

- **Most powerful feature**: Combines all capabilities
- Prefix-based command routing:
  - `/exec <command>` → Execute system command
  - `@ask <question>` → Quick AI question
  - `<message>` → AI chat (maintains history)
  - `/clear` → Clear chat history
  - `/exit` → Exit mode

## 🔧 Configuration

### Environment Variables

```bash
OPENAI_API_KEY    # Your OpenAI API key (required)
OPENAI_MODEL      # Optional: gpt-3.5-turbo (default), gpt-4, etc.
```

### Command-Line Flags

```bash
--api-key         # Override environment API key
--model, -m       # AI model selection
--temperature     # AI creativity (0.0-2.0)
--max-tokens      # Response length limit
--shell, -s       # Shell for command execution
--dir, -d         # Working directory
```

## 🚀 Usage Examples

### Basic Usage

```bash
# Welcome screen
./livecli

# Execute command
./livecli exec "ls -la"

# Quick question
./livecli ask "How do I find large files?"

# Chat session
./livecli chat

# Interactive mode
./livecli interactive
```

### Advanced Usage

```bash
# Different model
./livecli chat --model gpt-4

# Higher creativity
./livecli chat --temperature 1.5

# Custom shell
./livecli exec --shell bash "echo $SHELL"

# Specific directory
./livecli exec --dir /tmp "pwd"
```

### Interactive Workflow Example

```
> /exec git status
> @ask How do I write a good commit message?
> Can you review my changes based on the output above?
> /exec git diff
> Thanks! Let me commit now
> /exit
```

## 🎨 Design Highlights

### User Experience

- **Color-coded output**: Different colors for different message types
- **Real-time streaming**: See command output as it happens
- **Conversation history**: AI maintains context in chat mode
- **Clear visual separators**: Easy to distinguish between sections
- **Helpful prompts**: Clear indication of what mode you're in

### Code Quality

- **Modular architecture**: Each command in separate file
- **Error handling**: Comprehensive error messages
- **Cross-platform**: Works on Linux, macOS, and Windows
- **Clean abstractions**: Reusable functions for AI and exec operations

## 📦 Build & Distribution

### Local Build

```bash
go build -o livecli main.go
```

### Multi-Platform Build

```bash
make build-all  # Builds for Linux, macOS (amd64/arm64), Windows
```

### Installation

```bash
# Local use
./livecli

# Global installation
sudo mv livecli /usr/local/bin/

# Now use from anywhere
livecli
```

## 🔐 Security Considerations

1. **API Key Protection**:

   - Never commit `.env` file
   - Use environment variables
   - `.gitignore` includes `.env`

2. **Command Execution**:

   - Commands run with user's permissions
   - No privilege escalation
   - Careful with directory traversal

3. **AI Interactions**:
   - System prompts can be customized
   - Conversation history kept in memory only
   - No persistent storage of sensitive data

## 📚 Documentation

### User Documentation

- **README.md**: Complete documentation with all features
- **QUICKSTART.md**: Get started in 5 minutes
- **DEMO.md**: Detailed examples and workflows

### Developer Documentation

- Code is well-commented
- Each function has clear purpose
- Modular design for easy extension

## 🔮 Future Enhancements

Potential features to add:

- [ ] Command suggestions based on AI analysis
- [ ] Save conversation history to file
- [ ] Support for multiple AI providers (Anthropic, Cohere, etc.)
- [ ] Plugin system for custom commands
- [ ] Auto-completion for commands
- [ ] Web UI for chat interface
- [ ] Command history with search
- [ ] Configuration file support (YAML/JSON)
- [ ] Streaming responses for better UX
- [ ] Command templates and macros

## 🧪 Testing

### Manual Testing Checklist

- [x] Build completes successfully
- [x] `livecli` shows welcome screen
- [x] `livecli exec` executes commands
- [x] `livecli exec` shows real-time output
- [ ] `livecli chat` works with valid API key
- [ ] `livecli ask` returns AI responses
- [ ] `livecli interactive` handles all command types
- [x] Cross-platform shell detection works
- [x] Error messages are clear and helpful

### Adding Tests

Create test files:

```go
// cmd/exec_test.go
package cmd

import "testing"

func TestExecuteCommand(t *testing.T) {
    // Test command execution
}
```

Run tests:

```bash
go test ./...
make test
```

## 🤝 Contributing

To contribute:

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## 📝 License

MIT License - See LICENSE file for details

## 🙏 Acknowledgments

- **Cobra**: For excellent CLI framework
- **go-openai**: For OpenAI API client
- **color**: For terminal colors
- **readline**: For interactive input

## 📞 Support

- GitHub Issues: Report bugs or request features
- Command help: `livecli <command> --help`
- Documentation: README.md, QUICKSTART.md, DEMO.md

---

**Built with ❤️ using Go**

Version: 1.0.0  
Last Updated: December 2025
