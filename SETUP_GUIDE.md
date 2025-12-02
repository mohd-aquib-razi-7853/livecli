# 🎯 Setup Command Examples

The `setup` command is an intelligent assistant that uses AI to generate and execute installation commands for you!

## How It Works

1. **You describe what you want**: "setup rust into my system"
2. **AI generates commands**: Creates a step-by-step plan
3. **Shows you the plan**: Explains each command
4. **Asks for confirmation**: You approve before execution
5. **Executes safely**: Runs each command one by one

## Basic Usage

```bash
./livecli setup "rust into my system"
./livecli setup "docker and docker-compose"
./livecli setup "vscode editor"
./livecli setup "nodejs version 18"
./livecli setup "python3 and pip"
```

## Example Session

### Example 1: Installing Rust

```bash
$ ./livecli setup "rust into my system"

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

❓ Do you want to proceed with this setup plan? (yes/no): yes

🚀 Starting setup process...

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Step 1/4
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

📌 Install Rust using rustup
💻 Command: curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y

❓ Execute this command? (yes/no/skip all): yes

▶ Executing...
─────────────────────────────────────────────────────────────
[Rust installation output...]

✓ Step 1/4 completed

[... continues for each step ...]

╔═══════════════════════════════════════════════════════════╗
║           ✅ Setup Complete!                              ║
╚═══════════════════════════════════════════════════════════╝

✓ Executed 4 steps successfully

💡 Tip: Verify the installation with relevant commands (e.g., version checks)
```

### Example 2: Installing VS Code

```bash
$ ./livecli setup "vscode editor"

╔═══════════════════════════════════════════════════════════╗
║           🤖 AI Setup Assistant                           ║
╚═══════════════════════════════════════════════════════════╝

📋 Task: vscode editor

⏳ Analyzing your request and generating setup plan...

📝 Setup Plan:
─────────────────────────────────────────────────────────────

1. Import Microsoft GPG key
   Command: wget -qO- https://packages.microsoft.com/keys/microsoft.asc | gpg --dearmor > packages.microsoft.gpg

2. Install the GPG key
   Command: sudo install -D -o root -g root -m 644 packages.microsoft.gpg /etc/apt/keyrings/packages.microsoft.gpg

3. Add VS Code repository
   Command: echo "deb [arch=amd64,arm64,armhf signed-by=/etc/apt/keyrings/packages.microsoft.gpg] https://packages.microsoft.com/repos/code stable main" | sudo tee /etc/apt/sources.list.d/vscode.list

4. Update package index
   Command: sudo apt update

5. Install VS Code
   Command: sudo apt install -y code

6. [OPTIONAL] Clean up GPG file
   Command: rm -f packages.microsoft.gpg

7. Verify installation
   Command: code --version

─────────────────────────────────────────────────────────────

❓ Do you want to proceed with this setup plan? (yes/no): yes
```

### Example 3: Installing Docker

```bash
$ ./livecli setup "docker and docker-compose"

📝 Setup Plan:
─────────────────────────────────────────────────────────────

1. Update package index
   Command: sudo apt update

2. Install prerequisites
   Command: sudo apt install -y apt-transport-https ca-certificates curl software-properties-common

3. Add Docker GPG key
   Command: curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg

4. Add Docker repository
   Command: echo "deb [arch=amd64 signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list

5. Update package index
   Command: sudo apt update

6. Install Docker
   Command: sudo apt install -y docker-ce docker-ce-cli containerd.io docker-compose-plugin

7. Start Docker service
   Command: sudo systemctl start docker

8. Enable Docker on boot
   Command: sudo systemctl enable docker

9. [OPTIONAL] Add user to docker group
   Command: sudo usermod -aG docker $USER

10. Verify Docker
    Command: docker --version

11. Verify Docker Compose
    Command: docker compose version

─────────────────────────────────────────────────────────────
```

## Advanced Usage

### Auto-confirm all steps

```bash
./livecli setup "nodejs version 18" --yes
# or
./livecli setup "nodejs version 18" -y
```

### Dry run (see plan without executing)

```bash
./livecli setup "postgresql database" --dry-run
```

### Use different AI model

```bash
./livecli setup "kubernetes kubectl" --model gpt-4
```

## More Examples

### Programming Languages

```bash
# Python
./livecli setup "python3 and pip"

# Go
./livecli setup "golang latest version"

# Java
./livecli setup "openjdk 17"

# Ruby
./livecli setup "ruby and rails"
```

### Development Tools

```bash
# Git
./livecli setup "git with latest version"

# Vim with plugins
./livecli setup "vim editor with common plugins"

# Tmux
./livecli setup "tmux terminal multiplexer"

# Postman
./livecli setup "postman api client"
```

### Databases

```bash
# MySQL
./livecli setup "mysql server"

# PostgreSQL
./livecli setup "postgresql latest stable"

# MongoDB
./livecli setup "mongodb community edition"

# Redis
./livecli setup "redis server"
```

### Web Servers

```bash
# Nginx
./livecli setup "nginx web server"

# Apache
./livecli setup "apache2 web server"
```

### DevOps Tools

```bash
# Terraform
./livecli setup "terraform infrastructure tool"

# Ansible
./livecli setup "ansible automation"

# Jenkins
./livecli setup "jenkins ci/cd"
```

### Multiple Packages

```bash
# Development environment
./livecli setup "nodejs, npm, and yarn"

# Docker stack
./livecli setup "docker, docker-compose, and docker-buildx"

# Python data science
./livecli setup "python3, jupyter notebook, and pandas"
```

## Understanding the Output

### Step Status Indicators

- **📌** - Step description
- **💻** - Actual command to execute
- **▶** - Currently executing
- **✓** - Step completed successfully
- **⏭️** - Step skipped
- **[OPTIONAL]** - Optional step (you can skip)

### Confirmation Options

When prompted, you can respond with:

- `yes` or `y` - Execute this step
- `no` or `n` - Skip this step
- `skip all` or `skip` - Skip remaining steps

## Safety Features

1. **AI Validation**: AI generates safe, commonly-used commands
2. **OS Detection**: Commands tailored to your operating system
3. **Step-by-step confirmation**: Approve each command before execution
4. **Optional steps marked**: Non-critical steps clearly indicated
5. **Dry run mode**: Preview without executing
6. **Official sources only**: Uses official repositories and downloads

## Tips

### 1. Be Specific

```bash
# ❌ Vague
./livecli setup "code editor"

# ✅ Specific
./livecli setup "vscode editor with C++ extensions"
```

### 2. Mention Version if Needed

```bash
./livecli setup "nodejs version 18"
./livecli setup "python 3.11"
```

### 3. Use Dry Run First

```bash
# Preview the plan
./livecli setup "complex software" --dry-run

# Then execute if it looks good
./livecli setup "complex software"
```

### 4. Combine Related Tools

```bash
./livecli setup "docker, docker-compose, and kubectl"
```

## Common Use Cases

### Fresh System Setup

```bash
./livecli setup "essential development tools: git, curl, wget, build-essential"
```

### Web Development Environment

```bash
./livecli setup "nodejs, npm, and vscode for web development"
```

### Data Science Setup

```bash
./livecli setup "python3, pip, jupyter, pandas, numpy, and matplotlib"
```

### DevOps Workstation

```bash
./livecli setup "docker, kubectl, terraform, and aws cli"
```

## Troubleshooting

### If a step fails:

1. **Read the error message** - It's shown in real-time
2. **Skip and continue** - Choose "skip" for failed optional steps
3. **Check prerequisites** - Some tools need others installed first
4. **Manual intervention** - Some installations may need interactive input

### If AI generates incorrect commands:

1. **Use dry run** - Check the plan first
2. **Be more specific** - Provide more details in your request
3. **Try different model** - Use `--model gpt-4` for better results

## Comparison with Traditional Installation

### Traditional Way:

```bash
# You need to know all these commands
sudo apt update
sudo apt install -y docker.io
sudo systemctl start docker
sudo systemctl enable docker
sudo usermod -aG docker $USER
docker --version
```

### LiveCLI Way:

```bash
# Just describe what you want
./livecli setup "docker into my system"

# AI figures out all the commands
# You just confirm each step!
```

## Safety Notes

⚠️ **Important:**

- Always review the commands before confirming
- Some commands require sudo and will ask for your password
- Optional steps can be safely skipped
- You can stop at any time with "skip all"
- Use `--dry-run` when unsure

## Integration with Other Commands

You can combine setup with other LiveCLI features:

```bash
# Use interactive mode for complex setups
./livecli interactive

> I need to setup a complete web development environment
> [AI suggests tools]
> /exit

# Then use setup for each tool
./livecli setup "nodejs and npm"
./livecli setup "vscode"
./livecli setup "git"
```

---

**The setup command makes installation a breeze! 🚀**

Just tell LiveCLI what you want, and it figures out how to install it!
