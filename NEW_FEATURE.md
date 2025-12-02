# 🎉 NEW FEATURE: AI-Powered Setup Assistant

## What's New?

LiveCLI now includes an intelligent **setup command** that uses AI to automatically generate and execute installation commands for you!

## The Problem It Solves

**Before**: You wanted to install something but didn't know the exact commands:

```bash
# You had to Google:
"how to install rust on ubuntu"
"install vscode linux command line"
"docker installation ubuntu 22.04"

# Then copy-paste multiple commands
sudo apt update
sudo apt install ...
# Hope you got all the steps right
```

**Now**: Just tell LiveCLI what you want:

```bash
./livecli setup "rust into my system"
./livecli setup "vscode editor"
./livecli setup "docker and docker-compose"

# AI figures out all the commands
# Shows you what it will do
# Asks for confirmation
# Executes step-by-step!
```

## How It Works

### Step 1: You Describe What You Want

```bash
./livecli setup "rust into my system"
```

### Step 2: AI Generates Installation Plan

```
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

4. Verify Cargo installation
   Command: cargo --version

─────────────────────────────────────────────────────────────
```

### Step 3: You Review and Confirm

```
❓ Do you want to proceed with this setup plan? (yes/no): yes
```

### Step 4: Executes Step-by-Step

```
🚀 Starting setup process...

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Step 1/4
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

📌 Install Rust using rustup
💻 Command: curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y

❓ Execute this command? (yes/no/skip all): yes

▶ Executing...
[Shows real-time output]

✓ Step 1/4 completed

[Continues for each step...]

╔═══════════════════════════════════════════════════════════╗
║           ✅ Setup Complete!                              ║
╚═══════════════════════════════════════════════════════════╝

✓ Executed 4 steps successfully
```

## Key Features

### 🧠 Intelligent

- AI analyzes your request
- Detects your operating system
- Generates appropriate commands for your system
- Uses official repositories and sources

### 🔒 Safe

- Shows every command before executing
- Asks for confirmation at each step
- Clearly marks optional steps
- `--dry-run` mode to preview without executing

### 💡 Educational

- Explains what each command does
- You learn while installing
- See the exact commands being used

### ⚙️ Flexible

- `--yes` flag to auto-confirm all
- `--dry-run` to preview only
- Skip individual steps
- Stop at any time

## Usage Examples

### Basic Installation

```bash
# Programming languages
./livecli setup "rust into my system"
./livecli setup "python3 and pip"
./livecli setup "nodejs version 18"
./livecli setup "golang latest"

# Development tools
./livecli setup "vscode editor"
./livecli setup "git with latest version"
./livecli setup "docker and docker-compose"

# Databases
./livecli setup "postgresql database"
./livecli setup "mysql server"
./livecli setup "redis"

# Web servers
./livecli setup "nginx web server"
```

### Advanced Usage

```bash
# Auto-confirm (for scripts)
./livecli setup "docker" --yes

# Preview only
./livecli setup "kubernetes kubectl" --dry-run

# Use GPT-4 for more complex setups
./livecli setup "complete web dev environment" --model gpt-4

# Multiple tools at once
./livecli setup "nodejs, npm, yarn, and typescript"
```

## Command Options

```bash
livecli setup <task description> [flags]
```

### Flags

- `--yes, -y` - Auto-confirm all steps
- `--dry-run` - Show commands without executing
- `--model, -m` - AI model to use (default: gpt-3.5-turbo)
- `--api-key` - OpenAI API key

### Response Options

When asked to confirm a step:

- `yes` or `y` - Execute this step
- `no` or `n` - Skip this step
- `skip all` - Skip all remaining steps

## Real-World Examples

### Example 1: Fresh Ubuntu Setup

```bash
./livecli setup "essential development tools: git, curl, wget, build-essential"
```

### Example 2: Web Development Environment

```bash
./livecli setup "nodejs 18, npm, and vscode"
```

### Example 3: Docker for DevOps

```bash
./livecli setup "docker, docker-compose, and kubectl"
```

### Example 4: Data Science with Python

```bash
./livecli setup "python3, pip, jupyter notebook, pandas, and numpy"
```

## Safety Features

1. **OS Detection**: Commands tailored to your specific OS
2. **Confirmation Required**: Must approve the overall plan
3. **Step-by-Step Approval**: Confirm each command individually
4. **Optional Steps Marked**: Non-critical steps clearly indicated
5. **Dry Run Mode**: Preview without any execution
6. **Official Sources**: Uses only official repos and downloads
7. **Stop Anytime**: Can skip remaining steps at any point

## Tips for Best Results

### Be Specific

```bash
# ❌ Too vague
./livecli setup "editor"

# ✅ Specific
./livecli setup "vscode editor with C++ extensions"
```

### Mention Versions

```bash
./livecli setup "nodejs version 18"
./livecli setup "python 3.11"
```

### Use Dry Run First

```bash
# Preview first
./livecli setup "complex software" --dry-run

# Then execute if it looks good
./livecli setup "complex software"
```

## Comparison

### Traditional Way

```bash
# 1. Search on Google/StackOverflow
# 2. Find multiple commands
# 3. Copy-paste each one
# 4. Hope you didn't miss any steps
# 5. Debug errors without context

wget -qO- https://packages.microsoft.com/keys/microsoft.asc | gpg --dearmor > packages.microsoft.gpg
sudo install -D -o root -g root -m 644 packages.microsoft.gpg /etc/apt/keyrings/packages.microsoft.gpg
echo "deb [arch=amd64,arm64,armhf signed-by=/etc/apt/keyrings/packages.microsoft.gpg] https://packages.microsoft.com/repos/code stable main" | sudo tee /etc/apt/sources.list.d/vscode.list
sudo apt update
sudo apt install -y code
rm -f packages.microsoft.gpg
code --version
```

### LiveCLI Way

```bash
./livecli setup "vscode editor"

# ✓ AI figures out all commands
# ✓ Explains each step
# ✓ You confirm before execution
# ✓ Runs safely step-by-step
# ✓ Verifies installation
```

## Technical Details

### How It Works

1. **User Input**: You describe what you want
2. **OS Detection**: Automatically detects your OS and package manager
3. **AI Generation**: Sends request to OpenAI with OS context
4. **JSON Parsing**: Parses structured command plan from AI
5. **Presentation**: Shows plan with descriptions
6. **Confirmation**: Waits for user approval
7. **Execution**: Runs each command using the existing exec engine
8. **Verification**: Typically includes version check commands

### AI Prompt Engineering

The setup command uses a carefully crafted prompt that:

- Specifies the exact OS and package manager
- Requests JSON format for consistent parsing
- Asks for safe, commonly-used commands
- Includes verification steps
- Marks optional steps appropriately

### Error Handling

- If JSON parsing fails, shows debug output
- If a command fails, you can skip and continue
- Optional steps can be skipped without issues
- Full error messages displayed in real-time

## Use Cases

### 1. New System Setup

```bash
./livecli setup "git, vim, tmux, and htop"
```

### 2. Learning New Technologies

```bash
./livecli setup "kubernetes minikube for learning"
# See exactly how to install and verify
```

### 3. Team Onboarding

```bash
./livecli setup "company dev environment: docker, nodejs, postgresql"
```

### 4. Quick Prototyping

```bash
./livecli setup "redis and mongodb for prototyping"
```

### 5. CI/CD Setup

```bash
./livecli setup "jenkins ci/cd server"
```

## Limitations & Considerations

### Current Limitations

- Requires OpenAI API key (uses credits)
- AI-generated commands should be reviewed
- Some interactive installers may not work perfectly
- Best for common, well-documented software

### Best Practices

1. Always review the generated plan
2. Use `--dry-run` for unfamiliar software
3. Have sudo password ready if needed
4. Check system requirements first
5. Understand what each command does

### When NOT to Use

- Installing custom/proprietary software
- Complex multi-system setups
- When you need specific non-standard configurations
- In production without testing first

## Future Enhancements

Planned improvements:

- [ ] Save and reuse installation plans
- [ ] Custom templates for common setups
- [ ] Rollback capability if installation fails
- [ ] Integration with package managers for verification
- [ ] Support for configuration file generation
- [ ] Multi-machine setup orchestration

## FAQ

**Q: Is it safe to run AI-generated commands?**
A: The AI uses common, well-established commands, but you should ALWAYS review them before confirming. That's why each step requires confirmation.

**Q: Does it work on all operating systems?**
A: Currently optimized for Linux distributions. Support for macOS and Windows coming soon.

**Q: Can I customize the generated commands?**
A: Not directly, but you can use `--dry-run` to see the plan, then run commands manually with your modifications.

**Q: What if a step fails?**
A: You can skip the failed step and continue, or stop the entire process. All steps are independent when possible.

**Q: How much does it cost?**
A: It uses your OpenAI API credits. Each setup command typically costs less than $0.01 with GPT-3.5-turbo.

##Feedback & Contributions

We'd love your feedback! Try the new setup command and let us know:

- What worked well?
- What could be improved?
- What software did you install?
- Any edge cases we should handle?

## Get Started Now!

```bash
# Make sure you have an API key set
export OPENAI_API_KEY="your-key-here"

# Try it out!
./livecli setup "your favorite tool"

# See the guide for more examples
cat SETUP_GUIDE.md
```

---

**The future of software installation is here! 🚀**
