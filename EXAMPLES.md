# LiveCLI Examples & Use Cases

This document provides real-world examples and use cases for LiveCLI.

## Table of Contents

1. [System Administration](#system-administration)
2. [Development Workflows](#development-workflows)
3. [Learning & Education](#learning--education)
4. [DevOps Tasks](#devops-tasks)
5. [Quick Troubleshooting](#quick-troubleshooting)

---

## System Administration

### Example 1: Disk Space Management

```bash
# Check disk usage
./livecli exec "df -h"

# Find large files
./livecli exec "du -sh /* 2>/dev/null | sort -rh | head -10"

# Ask AI for help
./livecli ask "How can I safely clean up disk space on Linux?"

# Interactive workflow
./livecli interactive
> /exec df -h
> The /var partition is 95% full. What should I do?
> /exec du -sh /var/* | sort -rh | head -10
> @ask How do I safely remove old log files?
> /exec sudo find /var/log -name "*.log" -mtime +30 -type f
```

### Example 2: User Management

```bash
# List users
./livecli exec "cat /etc/passwd | cut -d: -f1"

# Check current users
./livecli exec "who"

# Get AI guidance
./livecli ask "How do I create a new user with sudo privileges?"

# Interactive session
./livecli interactive
> @ask What's the difference between useradd and adduser?
> /exec which adduser
> Can you give me the exact command to add a user named 'johndoe'?
```

### Example 3: Process Management

```bash
# View running processes
./livecli exec "ps aux | head -20"

# Find specific process
./livecli exec "ps aux | grep nginx"

# Get help with processes
./livecli ask "How do I kill a hanging process safely?"

# Interactive debugging
./livecli interactive
> /exec top -b -n 1 | head -20
> Process 1234 is using 90% CPU. What should I do?
> @ask How do I investigate high CPU usage?
> /exec ps -p 1234 -o pid,ppid,cmd,%mem,%cpu
```

---

## Development Workflows

### Example 4: Git Operations

```bash
# Check git status
./livecli exec "git status"

# Get commit message help
./livecli ask "What makes a good git commit message?"

# Interactive git workflow
./livecli interactive
> /exec git status
> I have uncommitted changes. Should I commit or stash?
> @ask What's the difference between git stash and git commit?
> /exec git diff
> Can you help me write a commit message for these changes?
> /exec git add .
> /exec git commit -m "feat: add user authentication module"
```

### Example 5: Code Review & Debugging

```bash
# Check code
./livecli exec "cat main.go"

# Get AI review
./livecli ask "What are common Go code smells to avoid?"

# Interactive debugging
./livecli interactive
> /exec go build
> I'm getting this error: "undefined: myFunction". What does it mean?
> @ask How do I check if a function is exported in Go?
> Can you explain Go's visibility rules?
> /exec grep -r "myFunction" .
```

### Example 6: Testing & CI/CD

```bash
# Run tests
./livecli exec "go test -v ./..."

# Get testing advice
./livecli ask "What are best practices for writing Go unit tests?"

# Interactive test debugging
./livecli interactive
> /exec go test -v ./cmd/...
> Test TestExecuteCommand is failing. How do I debug Go tests?
> @ask What flags can I use with go test?
> /exec go test -v -run TestExecuteCommand
```

---

## Learning & Education

### Example 7: Learning New Commands

```bash
# Learn about a command
./livecli ask "What does the 'awk' command do?"

# Get examples
./livecli ask "Give me 5 practical examples of using grep"

# Interactive learning
./livecli interactive
> I want to learn about sed command. Can you explain?
> /exec echo "Hello World" | sed 's/World/Universe/'
> That's cool! Can you show me more examples?
> @ask What's the difference between sed and awk?
> /exec echo -e "line1\nline2\nline3" | sed '2d'
```

### Example 8: Programming Concepts

```bash
# Understand concepts
./livecli ask "Explain goroutines in Go with examples"

# Compare languages
./livecli ask "What's the difference between Go and Rust for CLI tools?"

# Interactive learning
./livecli chat
You> What are Go interfaces?
AI> [explains interfaces]
You> Can you give me a practical example?
AI> [provides example]
You> How is this different from classes in Java?
AI> [explains differences]
```

---

## DevOps Tasks

### Example 9: Docker Operations

```bash
# Check Docker
./livecli exec "docker ps"

# Get Docker help
./livecli ask "How do I optimize Docker image size?"

# Interactive Docker workflow
./livecli interactive
> /exec docker images
> I have many <none> images. What are these?
> @ask How do I clean up Docker images safely?
> /exec docker system df
> Can you explain this output?
> /exec docker system prune -a --dry-run
```

### Example 10: Kubernetes Debugging

```bash
# Check pods
./livecli exec "kubectl get pods"

# Debug failing pod
./livecli interactive
> /exec kubectl get pods
> Pod 'myapp-123' is CrashLoopBackOff. What does this mean?
> @ask How do I debug a CrashLoopBackOff pod?
> /exec kubectl describe pod myapp-123
> Can you help me understand this output?
> /exec kubectl logs myapp-123
```

### Example 11: Server Monitoring

```bash
# Check system load
./livecli exec "uptime"

# Memory usage
./livecli exec "free -h"

# Interactive monitoring
./livecli interactive
> /exec uptime
> Load average is 15.0. Is this bad?
> @ask What's a healthy load average?
> /exec free -h
> I'm using 90% of memory. How do I find what's using it?
> /exec ps aux --sort=-%mem | head -10
```

---

## Quick Troubleshooting

### Example 12: Network Issues

```bash
# Check connectivity
./livecli exec "ping -c 3 google.com"

# Get help
./livecli ask "How do I troubleshoot network connectivity issues?"

# Interactive troubleshooting
./livecli interactive
> My website isn't loading. How do I troubleshoot?
> /exec ping -c 3 mywebsite.com
> Ping works. What should I check next?
> @ask How do I check if a web server is responding?
> /exec curl -I https://mywebsite.com
> Getting 502 error. What does this mean?
```

### Example 13: Permission Issues

```bash
# Check permissions
./livecli exec "ls -la myfile.txt"

# Get help
./livecli ask "I'm getting 'permission denied'. How do I fix this?"

# Interactive fixing
./livecli interactive
> /exec ls -la script.sh
> Can't execute this script. What's wrong?
> @ask How do I make a file executable?
> /exec chmod +x script.sh
> /exec ls -la script.sh
> Perfect! Now it shows -rwxr-xr-x. What does this mean?
```

### Example 14: Database Queries

```bash
# Quick database check
./livecli exec "mysql -u root -p -e 'SHOW DATABASES;'"

# Get SQL help
./livecli ask "How do I optimize a slow MySQL query?"

# Interactive database work
./livecli interactive
> I need to write a query to find duplicate emails in users table
> /exec mysql -u root -p mydb -e "SELECT email, COUNT(*) FROM users GROUP BY email HAVING COUNT(*) > 1;"
> Can you help me write a query to delete the duplicates?
> @ask What's the safest way to remove duplicate rows in MySQL?
```

---

## Advanced Use Cases

### Example 15: Automated Workflows

```bash
#!/bin/bash
# Script using livecli for automated tasks

# Generate deployment plan
PLAN=$(./livecli ask "What are the steps to deploy a Go application?")
echo "$PLAN" > deployment_plan.txt

# Execute with AI guidance
./livecli interactive <<EOF
/exec git pull origin main
Am I ready to deploy?
/exec go test ./...
Tests passed! What's next?
/exec go build -o myapp
How do I verify the build?
/exit
EOF
```

### Example 16: Code Generation

```bash
# Get code help
./livecli ask "Generate a Go HTTP server with basic routing"

# Interactive code development
./livecli chat
You> I need to create a REST API in Go. Can you help?
AI> [provides guidance]
You> Show me code for a simple GET endpoint
AI> [provides code]
You> How do I add middleware for logging?
AI> [provides middleware code]
You> Can you put it all together?
AI> [provides complete example]
```

### Example 17: Log Analysis

```bash
# Analyze logs
./livecli exec "tail -100 /var/log/syslog"

# Interactive log analysis
./livecli interactive
> /exec tail -100 /var/log/nginx/error.log
> I see many 404 errors. How do I analyze this?
> @ask How do I find the most common 404 URLs?
> /exec grep "404" /var/log/nginx/access.log | awk '{print $7}' | sort | uniq -c | sort -rn | head
> Can you explain this output?
```

---

## Workflow Templates

### Template 1: New Project Setup

```bash
./livecli interactive
> I want to start a new Go CLI project. What's the best structure?
> /exec mkdir myproject && cd myproject
> /exec go mod init github.com/username/myproject
> What files should I create first?
> /exec touch main.go README.md .gitignore
> Can you give me a template for main.go?
> [AI provides template]
> /exec git init
> What should be in my .gitignore for Go?
```

### Template 2: System Health Check

```bash
./livecli interactive
> Run a complete system health check for me
> /exec df -h
> /exec free -h
> /exec uptime
> /exec systemctl --failed
> Based on these outputs, are there any issues I should address?
> What's my next step?
```

### Template 3: Security Audit

```bash
./livecli interactive
> I want to do a basic security audit of my server
> @ask What should I check for a basic security audit?
> /exec sudo netstat -tulpn
> /exec sudo ufw status
> /exec cat /etc/ssh/sshd_config | grep PermitRootLogin
> Are there any security issues here?
```

---

## Tips for Maximum Productivity

1. **Combine Commands**: Use interactive mode to chain related tasks
2. **Learn While Working**: Ask AI to explain outputs you don't understand
3. **Save Common Workflows**: Create scripts with common interactive sessions
4. **Use History**: Readline history means you can recall previous commands
5. **Experiment Safely**: Use `--dry-run` flags with AI guidance before real execution

---

## Getting Creative

### Example: AI Pair Programming

```bash
./livecli interactive
> I'm building a URL shortener. Let's design it together
> What database should I use?
> /exec which sqlite3
> Great, SQLite is available. Can you design the schema?
> What API endpoints do I need?
> Let's start with the main.go structure
> /exec touch main.go handlers.go db.go
> Show me the code for db.go
> [Copy the code AI provides]
> Now handlers.go
> [Continue iteratively]
```

### Example: Learning Through Doing

```bash
./livecli interactive
> Teach me about Linux file permissions by example
> /exec touch testfile
> /exec ls -l testfile
> What does -rw-r--r-- mean?
> /exec chmod 755 testfile
> /exec ls -l testfile
> What changed and why?
> Show me all possible permission combinations
```

---

**Remember**: LiveCLI is most powerful when you combine execution with AI assistance in interactive mode!
