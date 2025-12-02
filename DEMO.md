# LiveCLI Demo Guide 🎬

This guide will walk you through all the features of LiveCLI with practical examples.

## Prerequisites

Make sure you have:

1. Built the application: `go build -o livecli main.go`
2. Set your OpenAI API key: `export OPENAI_API_KEY="your-key-here"`

## Demo 1: Welcome Screen

```bash
./livecli
```

This will display the welcome screen with all available commands.

## Demo 2: Execute System Commands

### Simple Command

```bash
./livecli exec "echo Hello, LiveCLI!"
```

### List Files with Details

```bash
./livecli exec "ls -lah"
```

### Check System Info

```bash
./livecli exec "uname -a"
```

### Git Status (if in a git repo)

```bash
./livecli exec "git status"
```

### Execute in Specific Directory

```bash
./livecli exec --dir /tmp "pwd && ls"
```

### Find Large Files

```bash
./livecli exec "find . -type f -size +1M -exec ls -lh {} \; | head -10"
```

## Demo 3: AI Quick Questions

### System Administration

```bash
./livecli ask "How do I find all running processes?"
```

### Programming Help

```bash
./livecli ask "Explain the difference between sync and async in programming"
```

### Command Line Help

```bash
./livecli ask "How do I use grep to search for text in files?"
```

### Debugging Help

```bash
./livecli ask "What does 'permission denied' error mean and how to fix it?"
```

### Go Programming

```bash
./livecli ask "How do I handle errors in Go?"
```

## Demo 4: Interactive Chat Session

Start a chat session:

```bash
./livecli chat
```

Then try these conversations:

**Example 1: Learn About Docker**

```
You> What is Docker?
AI> [explains Docker]

You> How do I create a Dockerfile?
AI> [explains Dockerfile creation]

You> Can you give me an example for a Go application?
AI> [provides example]
```

**Example 2: Debugging Help**

```
You> I'm getting a segmentation fault in my C program. How do I debug it?
AI> [provides debugging strategies]

You> What tools can I use?
AI> [suggests gdb, valgrind, etc.]
```

**Commands in Chat**:

- `/clear` - Clear conversation history
- `/exit` - Exit chat session

## Demo 5: Interactive Mode (Most Powerful!)

Start interactive mode:

```bash
./livecli interactive
```

Now you can mix commands and chat:

**Workflow Example: System Cleanup**

```
> @ask How do I check disk usage?

💡 Answer:
Use the `df` command for disk usage...

> /exec df -h

▶ Executing: df -h
[shows disk usage]

> @ask How can I find large files?

💡 Answer:
Use: find / -type f -size +100M...

> /exec find . -type f -size +10M -exec ls -lh {} \;

[shows large files]

> Can you explain what the find command does in detail?

AI> [detailed explanation of find command]

> /clear

✓ Chat history cleared
```

**Workflow Example: Git Operations**

```
> /exec git status

[shows git status]

> @ask How do I write a good commit message?

[AI provides guidelines]

> Thanks! Let me check the diff

AI> [responds]

> /exec git diff

[shows changes]
```

## Demo 6: Advanced Features

### Use Different AI Models

```bash
./livecli chat --model gpt-4
```

### Adjust AI Creativity

```bash
# More creative (higher temperature)
./livecli chat --temperature 1.5

# More focused (lower temperature)
./livecli chat --temperature 0.3
```

### Longer Responses

```bash
./livecli chat --max-tokens 2000
```

### Custom System Prompt

```bash
./livecli chat --system "You are a Linux system administration expert. Provide detailed technical answers."
```

## Demo 7: Real-World Workflows

### Workflow 1: Setup a New Go Project

```bash
./livecli interactive
```

```
> /exec mkdir myproject && cd myproject

> @ask How do I initialize a Go module?

> /exec go mod init github.com/user/myproject

> @ask What's a good project structure for a Go CLI app?

> Can you help me create a main.go file for a simple CLI?
```

### Workflow 2: Debug Production Issues

```bash
./livecli interactive
```

```
> /exec tail -n 50 /var/log/syslog

> @ask How do I analyze these logs for errors?

> /exec grep -i error /var/log/syslog | tail -20

> What could cause these errors?

> @ask How do I check if a service is running?

> /exec systemctl status myservice
```

### Workflow 3: Learn While Doing

```bash
./livecli interactive
```

```
> I need to learn about Docker. Can you start with the basics?

AI> [explains Docker basics]

> /exec docker --version

> @ask How do I list running containers?

> /exec docker ps

> @ask How do I view logs for a container?

> Can you explain Docker networking?
```

## Demo 8: Combining with Other Tools

### Use with Pipes (in exec)

```bash
./livecli exec "ps aux | grep python"
```

### Use with Environment Variables

```bash
export MY_VAR="test"
./livecli exec "echo \$MY_VAR"
```

### Chain Commands

```bash
./livecli exec "cd /tmp && touch testfile && ls -la testfile"
```

## Tips and Tricks

1. **Save Common Commands**: Create aliases in your `.bashrc` or `.zshrc`:

   ```bash
   alias lai='./livecli interactive'
   alias lask='./livecli ask'
   alias lexec='./livecli exec'
   ```

2. **Use in Scripts**: You can call LiveCLI from shell scripts:

   ```bash
   #!/bin/bash
   result=$(./livecli ask "Best practices for bash scripting")
   echo "$result"
   ```

3. **Quick Debugging**:

   ```bash
   ./livecli ask "Why is my command failing: $(./livecli exec 'your-command' 2>&1)"
   ```

4. **Learning Mode**: Use interactive mode to learn as you work
   ```
   > @ask How do I do X?
   > /exec [try the command]
   > Can you explain what just happened?
   ```

## Troubleshooting

### API Key Issues

```bash
# Test if API key is set
echo $OPENAI_API_KEY

# Set temporarily
export OPENAI_API_KEY="your-key"

# Or use flag
./livecli chat --api-key="your-key"
```

### Build Issues

```bash
# Clean and rebuild
go clean
go mod tidy
go build -o livecli main.go
```

### Permission Issues

```bash
# Make executable
chmod +x livecli

# For global install
sudo cp livecli /usr/local/bin/
```

## Next Steps

1. Install globally: `sudo mv livecli /usr/local/bin/`
2. Set up shell aliases
3. Explore the `--help` for each command
4. Read the full README.md
5. Customize the system prompts for your use case

Enjoy using LiveCLI! 🚀
