# 🚀 NEW FEATURE: Git Workflow Automation

## What's New?

LiveCLI now includes a powerful **git command** that automates your daily git workflow!

## The Problem It Solves

**Before**: You had to type three commands every time you wanted to save your work:

```bash
git add .
git commit -m "feat: my new feature"
git push
```

**Now**: Do it all in one command:

```bash
./livecli git "feat: my new feature"
```

## How It Works

### Step 1: Run the Command

```bash
./livecli git "update documentation"
```

### Step 2: Review the Plan

```
╔═══════════════════════════════════════════════════════════╗
║              🚀 Git Workflow Automator                    ║
╚═══════════════════════════════════════════════════════════╝

📋 Commit Message: update documentation

📝 Execution Plan:
─────────────────────────────────────────────────────────────

1. Stage all changes
   Command: git add .

2. Commit changes
   Command: git commit -m "update documentation"

3. Push to remote
   Command: git push

─────────────────────────────────────────────────────────────

❓ Proceed with these git operations? (yes/no): yes
```

### Step 3: Watch It Execute

```
🚀 Executing git workflow...

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Step 1/3: Stage all changes
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
✓ Stage all changes completed

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Step 2/3: Commit changes
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
[main 1234abc] update documentation
 1 file changed, 10 insertions(+)
✓ Commit changes completed

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Step 3/3: Push to remote
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
To github.com:user/repo.git
   abcd..1234  main -> main
✓ Push to remote completed

╔═══════════════════════════════════════════════════════════╗
║           ✅ Git Workflow Complete!                       ║
╚═══════════════════════════════════════════════════════════╝
```

## Key Features

### ⚡ Fast

- Combines 3 commands into 1
- Saves you keystrokes and time

### 🔒 Safe

- Shows you exactly what will happen
- Asks for confirmation before running
- Stops immediately if any step fails

### 🤖 Automatable

- Use `--yes` flag to skip confirmation
- Perfect for scripts and CI/CD

## Usage Examples

### Basic Usage

```bash
./livecli git "fix: login bug"
```

### Auto-Confirm

```bash
./livecli git "wip: save work" --yes
```

### In Scripts

```bash
#!/bin/bash
# Do some work...
echo "updated" > status.txt
# Save changes
./livecli git "chore: update status" --yes
```

## Tips

1. **Write Good Commit Messages**: Since it's so easy to commit, make sure your messages are still descriptive!
2. **Check Status First**: You might want to run `git status` (or `./livecli exec "git status"`) before running this to see what will be added.
3. **Use with Interactive Mode**:
   ```bash
   ./livecli interactive
   > /exec git status
   > /exec livecli git "feat: new feature"
   ```

---

**Happy Coding! 🚀**
