# LiveCLI Installation Guide

Complete installation instructions for **Linux**, **macOS**, and **Windows**.

---

## Table of Contents

- [Quick Install](#quick-install)
- [Manual Installation](#manual-installation)
  - [Linux](#linux)
  - [macOS](#macos)
  - [Windows](#windows)
- [Building from Source](#building-from-source)
- [Post-Installation Setup](#post-installation-setup)
- [Troubleshooting](#troubleshooting)

---

## Quick Install

### Linux & macOS

```bash
# Using the installation script (recommended)
curl -fsSL https://raw.githubusercontent.com/mohd-aquib-razi-7853/livecli/main/install.sh | bash

# OR using wget
wget -qO- https://raw.githubusercontent.com/mohd-aquib-razi-7853/livecli/main/install.sh | bash
```

### Windows

```powershell
# Using PowerShell (run as Administrator recommended)
iwr -useb https://raw.githubusercontent.com/mohd-aquib-razi-7853/livecli/main/install.ps1 | iex

# OR download and run the script
Invoke-WebRequest -Uri "https://raw.githubusercontent.com/mohd-aquib-razi-7853/livecli/main/install.ps1" -OutFile "install.ps1"
PowerShell -ExecutionPolicy Bypass -File install.ps1
```

---

## Manual Installation

### Linux

#### Using Pre-built Binary

1. **Download the appropriate binary for your architecture:**

   ```bash
   # For x86_64 (most common)
   wget https://github.com/mohd-aquib-razi-7853/livecli/releases/latest/download/livecli-linux-amd64

   # For ARM64 (e.g., Raspberry Pi, AWS Graviton)
   wget https://github.com/mohd-aquib-razi-7853/livecli/releases/latest/download/livecli-linux-arm64
   ```

2. **Make it executable:**

   ```bash
   chmod +x livecli-linux-*
   ```

3. **Move to a directory in your PATH:**

   ```bash
   sudo mv livecli-linux-* /usr/local/bin/livecli
   ```

4. **Verify installation:**

   ```bash
   livecli --help
   ```

#### Using Package Managers

**Debian/Ubuntu (APT):**

```bash
# Coming soon - DEB package
# sudo dpkg -i livecli_1.0.0_amd64.deb
```

**Fedora/RHEL/CentOS (DNF/YUM):**

```bash
# Coming soon - RPM package
# sudo dnf install livecli-1.0.0.x86_64.rpm
```

**Arch Linux (AUR):**

```bash
# Coming soon - AUR package
# yay -S livecli
```

---

### macOS

#### Using Pre-built Binary

1. **Download the appropriate binary:**

   ```bash
   # For Intel Macs
   curl -LO https://github.com/mohd-aquib-razi-7853/livecli/releases/latest/download/livecli-darwin-amd64

   # For Apple Silicon (M1/M2/M3)
   curl -LO https://github.com/mohd-aquib-razi-7853/livecli/releases/latest/download/livecli-darwin-arm64
   ```

2. **Make it executable:**

   ```bash
   chmod +x livecli-darwin-*
   ```

3. **Move to a directory in your PATH:**

   ```bash
   sudo mv livecli-darwin-* /usr/local/bin/livecli
   ```

4. **Remove quarantine attribute (macOS security):**

   ```bash
   xattr -d com.apple.quarantine /usr/local/bin/livecli
   ```

5. **Verify installation:**

   ```bash
   livecli --help
   ```

#### Using Homebrew

```bash
# Coming soon - Homebrew formula
# brew tap mohd-aquib-razi-7853/livecli
# brew install livecli
```

---

### Windows

#### Using Pre-built Binary

1. **Download the appropriate binary:**

   - **x86_64**: [livecli-windows-amd64.exe](https://github.com/mohd-aquib-razi-7853/livecli/releases/latest/download/livecli-windows-amd64.exe)
   - **ARM64**: [livecli-windows-arm64.exe](https://github.com/mohd-aquib-razi-7853/livecli/releases/latest/download/livecli-windows-arm64.exe)

2. **Create installation directory:**

   ```powershell
   New-Item -ItemType Directory -Path "$env:LOCALAPPDATA\Programs\livecli" -Force
   ```

3. **Move the executable:**

   ```powershell
   Move-Item -Path ".\livecli-windows-*.exe" -Destination "$env:LOCALAPPDATA\Programs\livecli\livecli.exe"
   ```

4. **Add to PATH:**

   ```powershell
   # Add to User PATH
   $currentPath = [Environment]::GetEnvironmentVariable("Path", "User")
   $newPath = "$env:LOCALAPPDATA\Programs\livecli"
   [Environment]::SetEnvironmentVariable("Path", "$currentPath;$newPath", "User")
   ```

5. **Restart your terminal and verify:**

   ```powershell
   livecli --help
   ```

#### Using Chocolatey

```powershell
# Coming soon - Chocolatey package
# choco install livecli
```

#### Using Scoop

```powershell
# Coming soon - Scoop package
# scoop bucket add mohd-aquib-razi-7853 https://github.com/mohd-aquib-razi-7853/scoop-bucket
# scoop install livecli
```

---

## Building from Source

### Prerequisites

- **Go 1.21 or later** ([Download Go](https://go.dev/dl/))
- **Git**
- **Make** (optional, but recommended)

### Build Steps

1. **Clone the repository:**

   ```bash
   git clone https://github.com/mohd-aquib-razi-7853/livecli.git
   cd livecli
   ```

2. **Install dependencies:**

   ```bash
   go mod download
   ```

3. **Build for your platform:**

   ```bash
   # Quick build for current platform
   make build

   # OR build for all platforms
   make build-all
   ```

4. **Install locally:**

   ```bash
   # Linux/macOS
   make install

   # Windows (from build directory)
   copy build\livecli-windows-amd64.exe %LOCALAPPDATA%\Programs\livecli\livecli.exe
   ```

### Build Targets

The Makefile provides several build targets:

| Command          | Description                                                     |
| ---------------- | --------------------------------------------------------------- |
| `make build`     | Build for current platform                                      |
| `make build-all` | Build for all platforms (Linux, macOS, Windows - amd64 & arm64) |
| `make install`   | Build and install to `/usr/local/bin` (Linux/macOS)             |
| `make clean`     | Remove build artifacts                                          |
| `make test`      | Run tests                                                       |
| `make deps`      | Update dependencies                                             |

---

## Post-Installation Setup

### 1. Set OpenAI API Key

LiveCLI requires an OpenAI API key to function. Get your API key from [OpenAI Platform](https://platform.openai.com/api-keys).

**Linux/macOS:**

```bash
# Temporary (current session only)
export OPENAI_API_KEY='your-api-key-here'

# Permanent - add to ~/.bashrc or ~/.zshrc
echo 'export OPENAI_API_KEY="your-api-key-here"' >> ~/.bashrc
source ~/.bashrc
```

**Windows:**

```powershell
# Temporary (current session only)
$env:OPENAI_API_KEY = "your-api-key-here"

# Permanent (User environment variable)
setx OPENAI_API_KEY "your-api-key-here"

# Permanent (System environment variable - requires admin)
[System.Environment]::SetEnvironmentVariable('OPENAI_API_KEY', 'your-api-key-here', 'Machine')
```

### 2. Verify Installation

```bash
# Check version
livecli --version

# View help
livecli --help

# Test with a simple command
livecli ask "What is the current date command?"
```

### 3. Try Key Features

```bash
# AI-powered setup
livecli setup "docker"

# Execute commands
livecli exec "ls -la"

# Start chat mode
livecli chat

# Git automation
livecli git "Add installation guide"

# Interactive mode
livecli interactive
```

---

## Troubleshooting

### Command not found

**Problem:** Terminal says `livecli: command not found`

**Solution:**

- Ensure the installation directory is in your PATH
- Restart your terminal after installation
- Verify the binary is in the expected location:

  ```bash
  # Linux/macOS
  which livecli

  # Windows
  where livecli
  ```

### Permission denied

**Problem:** `Permission denied` when trying to run livecli

**Solution:**

```bash
# Linux/macOS - make sure it's executable
chmod +x /usr/local/bin/livecli
```

### macOS quarantine error

**Problem:** macOS blocks the binary saying it's from an unidentified developer

**Solution:**

```bash
xattr -d com.apple.quarantine /usr/local/bin/livecli
```

Or go to **System Preferences → Security & Privacy** and click "Open Anyway"

### Windows SmartScreen

**Problem:** Windows SmartScreen prevents running the executable

**Solution:**

- Click "More info" → "Run anyway"
- Or right-click the file → Properties → Check "Unblock" → Apply

### API Key Issues

**Problem:** Error about missing API key

**Solution:**

- Verify your API key is set: `echo $OPENAI_API_KEY` (Linux/macOS) or `echo $env:OPENAI_API_KEY` (Windows)
- Use the `--api-key` flag: `livecli --api-key=your-key chat`
- Check for typos in your API key

### Build from source fails

**Problem:** Build errors when compiling from source

**Solution:**

```bash
# Ensure you have the correct Go version
go version  # Should be 1.21+

# Clean and rebuild
make clean
go mod tidy
make build
```

---

## Upgrading

### Using Installation Script

Simply run the installation script again - it will overwrite the old version:

```bash
# Linux/macOS
curl -fsSL https://raw.githubusercontent.com/mohd-aquib-razi-7853/livecli/main/install.sh | bash

# Windows
iwr -useb https://raw.githubusercontent.com/mohd-aquib-razi-7853/livecli/main/install.ps1 | iex
```

### Manual Upgrade

Download the latest binary and replace the existing one following the [Manual Installation](#manual-installation) steps.

---

## Uninstalling

### Linux/macOS

```bash
sudo rm /usr/local/bin/livecli
```

### Windows

```powershell
Remove-Item "$env:LOCALAPPDATA\Programs\livecli" -Recurse -Force

# Remove from PATH manually via System Environment Variables
# Or using PowerShell:
$path = [Environment]::GetEnvironmentVariable("Path", "User")
$newPath = ($path.Split(';') | Where-Object { $_ -ne "$env:LOCALAPPDATA\Programs\livecli" }) -join ';'
[Environment]::SetEnvironmentVariable("Path", $newPath, "User")
```

---

## Support

- **Issues**: [GitHub Issues](https://github.com/mohd-aquib-razi-7853/livecli/issues)
- **Discussions**: [GitHub Discussions](https://github.com/mohd-aquib-razi-7853/livecli/discussions)
- **Documentation**: [README.md](README.md)

---

## License

LiveCLI is open source software. See [LICENSE](LICENSE) for details.
