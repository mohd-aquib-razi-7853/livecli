# 🚀 Quick Start Guide - LiveCLI

Get up and running with LiveCLI in 5 minutes!

## Step 1: Build (30 seconds)

```bash
# Download dependencies and build
go mod download
go build -o livecli main.go

# Or simply use make
make build
```

## Step 2: Set API Key (1 minute)

Get your OpenAI API key from: https://platform.openai.com/api-keys

Then set it:

```bash
export OPENAI_API_KEY="sk-your-key-here"

# Add to your shell config to make it permanent
echo 'export OPENAI_API_KEY="sk-your-key-here"' >> ~/.bashrc
source ~/.bashrc
```

## Step 3: Test It! (30 seconds)

```bash
# Show welcome screen
./livecli

# Execute a command
./livecli exec "ls -la"

# Ask AI a question
./livecli ask "How do I find large files in Linux?"
```

## Step 4: Try Interactive Mode (2 minutes)

```bash
./livecli interactive
```

Then try:

```
> /exec pwd
> @ask How do I check disk space?
> /exec df -h
> Can you explain the output?
> /exit
```

## Step 5: Install Globally (Optional, 1 minute)

```bash
sudo mv livecli /usr/local/bin/
# Now you can use it from anywhere:
livecli
```

## Common Commands

| Command               | Description        | Example                         |
| --------------------- | ------------------ | ------------------------------- |
| `livecli`             | Show help          | `livecli`                       |
| `livecli exec`        | Run system command | `livecli exec "git status"`     |
| `livecli ask`         | Quick AI question  | `livecli ask "What is Docker?"` |
| `livecli chat`        | Start AI chat      | `livecli chat`                  |
| `livecli interactive` | Combined mode      | `livecli interactive`           |

## Pro Tips

1. **Create aliases** for faster access:

   ```bash
   alias li='livecli interactive'
   alias la='livecli ask'
   alias lx='livecli exec'
   ```

2. **Use different models**:

   ```bash
   livecli chat --model gpt-4
   ```

3. **Adjust AI creativity**:
   ```bash
   livecli chat --temperature 0.9  # More creative
   livecli chat --temperature 0.3  # More focused
   ```

## Troubleshooting

**Issue**: "API key not set"

```bash
echo $OPENAI_API_KEY  # Check if set
export OPENAI_API_KEY="your-key"  # Set it
```

**Issue**: "Command not found"

```bash
./livecli  # Run from current directory
# OR install globally first
```

**Issue**: Build errors

```bash
go mod tidy
go clean
go build -o livecli main.go
```

## What's Next?

- Read **README.md** for complete documentation
- Check **DEMO.md** for detailed examples and workflows
- Run `./livecli <command> --help` for command-specific help

## Support

- Check available commands: `./livecli --help`
- Command help: `./livecli chat --help`
- Issues: Create an issue on GitHub

---

**That's it! You're ready to use LiveCLI! 🎉**

Start with: `./livecli interactive` for the best experience.
