# Cross-Platform Distribution Guide

This guide explains how to build, test, and distribute **LiveCLI** for Windows, Linux, and macOS.

---

## Quick Start

### Build for All Platforms

```bash
# Build binaries for all platforms (Linux, macOS, Windows - both amd64 and arm64)
make build-all
```

This will create binaries in the `build/` directory:

- `livecli-linux-amd64`
- `livecli-linux-arm64`
- `livecli-darwin-amd64` (Intel Macs)
- `livecli-darwin-arm64` (Apple Silicon M1/M2/M3)
- `livecli-windows-amd64.exe`
- `livecli-windows-arm64.exe`
- `checksums.txt` (SHA256 checksums for verification)

### Verify Builds

```bash
# Run verification script
./verify-builds.sh
```

---

## Platform Support Matrix

| Platform    | Architecture          | Status             | Binary Name                 |
| ----------- | --------------------- | ------------------ | --------------------------- |
| **Linux**   | x86_64 (amd64)        | ✅ Fully Supported | `livecli-linux-amd64`       |
| **Linux**   | ARM64                 | ✅ Fully Supported | `livecli-linux-arm64`       |
| **macOS**   | Intel (amd64)         | ✅ Fully Supported | `livecli-darwin-amd64`      |
| **macOS**   | Apple Silicon (arm64) | ✅ Fully Supported | `livecli-darwin-arm64`      |
| **Windows** | x86_64 (amd64)        | ✅ Fully Supported | `livecli-windows-amd64.exe` |
| **Windows** | ARM64                 | ✅ Fully Supported | `livecli-windows-arm64.exe` |

### Tested Platforms

- ✅ Ubuntu 20.04, 22.04, 24.04
- ✅ Debian 11, 12
- ✅ Fedora 38, 39, 40
- ✅ CentOS/RHEL 8, 9
- ✅ Arch Linux
- ✅ macOS 12 (Monterey), 13 (Ventura), 14 (Sonoma)
- ✅ Windows 10, 11
- ✅ Raspberry Pi (ARM64)
- ✅ AWS Graviton (ARM64)

---

## Building from Source

### Prerequisites

- **Go 1.21+** (required)
- **Make** (optional, but recommended)
- **Git** (for version control)

### Build Commands

```bash
# 1. Clone the repository
git clone https://github.com/mohd-aquib-razi-7853/livecli.git
cd livecli

# 2. Install dependencies
go mod download

# 3. Build for all platforms
make build-all

# OR build for specific platform
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o build/livecli-linux-amd64 main.go
GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o build/livecli-darwin-arm64 main.go
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o build/livecli-windows-amd64.exe main.go
```

### Build Flags Explained

- `-ldflags="-s -w"`: Strip debug information to reduce binary size
  - `-s`: Omit symbol table
  - `-w`: Omit DWARF debug information
- `CGO_ENABLED=0`: Disable CGO for pure Go build (better cross-compilation)

---

## Automated Builds (GitHub Actions)

The repository includes a GitHub Actions workflow (`.github/workflows/release.yml`) that automatically:

1. **Builds** binaries for all 6 platforms
2. **Generates** SHA256 checksums
3. **Runs** tests
4. **Creates** GitHub releases when you push a tag

### Creating a Release

```bash
# 1. Tag your release
git tag v1.0.0

# 2. Push the tag
git push origin v1.0.0

# 3. GitHub Actions will automatically:
#    - Build all binaries
#    - Run tests
#    - Create a GitHub release with all binaries attached
```

---

## Distribution Methods

### 1. GitHub Releases (Recommended)

**Pros:**

- Automatic builds via GitHub Actions
- Built-in checksums and verification
- Version history
- Easy to download

**Setup:**

1. Push a tag: `git tag v1.0.0 && git push origin v1.0.0`
2. Binaries are automatically built and attached to the release
3. Users download from: `https://github.com/mohd-aquib-razi-7853/livecli/releases`

### 2. Installation Scripts

**For Linux/macOS:**

```bash
curl -fsSL https://raw.githubusercontent.com/mohd-aquib-razi-7853/livecli/main/install.sh | bash
```

**For Windows:**

```powershell
iwr -useb https://raw.githubusercontent.com/mohd-aquib-razi-7853/livecli/main/install.ps1 | iex
```

### 3. Package Managers (Future)

#### Homebrew (macOS/Linux)

```ruby
# Formula: homebrew-tap/livecli.rb
class Livecli < Formula
  desc "AI-powered command-line interface"
  homepage "https://github.com/mohd-aquib-razi-7853/livecli"
  version "1.0.0"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/mohd-aquib-razi-7853/livecli/releases/download/v1.0.0/livecli-darwin-arm64"
      sha256 "..."
    else
      url "https://github.com/mohd-aquib-razi-7853/livecli/releases/download/v1.0.0/livecli-darwin-amd64"
      sha256 "..."
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/mohd-aquib-razi-7853/livecli/releases/download/v1.0.0/livecli-linux-arm64"
      sha256 "..."
    else
      url "https://github.com/mohd-aquib-razi-7853/livecli/releases/download/v1.0.0/livecli-linux-amd64"
      sha256 "..."
    end
  end

  def install
    bin.install "livecli-#{OS.kernel_name.downcase}-#{Hardware::CPU.arch}" => "livecli"
  end
end
```

#### Chocolatey (Windows)

```powershell
# Coming soon
choco install livecli
```

#### Snap (Linux)

```bash
# Coming soon
snap install livecli
```

#### AUR (Arch Linux)

```bash
# Coming soon
yay -S livecli
```

---

## Platform-Specific Notes

### Linux

**Package Managers:**

- Ubuntu/Debian: Use `apt` for system packages
- Fedora/RHEL: Use `dnf`/`yum`
- Arch: Use `pacman`

**Installation Location:** `/usr/local/bin/livecli`

**Common Issues:**

- Ensure binary is executable: `chmod +x livecli-linux-*`
- May need sudo for installation to `/usr/local/bin`

### macOS

**Architectures:**

- **Intel Macs**: Use `livecli-darwin-amd64`
- **Apple Silicon (M1/M2/M3)**: Use `livecli-darwin-arm64`

**Gatekeeper:**
macOS may block the binary. Remove quarantine:

```bash
xattr -d com.apple.quarantine /usr/local/bin/livecli
```

**Package Manager:**

- Homebrew (preferred): `brew install livecli` (when formula is ready)

### Windows

**PowerShell vs CMD:**

- LiveCLI works with both PowerShell and Command Prompt
- PowerShell recommended for better experience

**Installation Location:** `%LOCALAPPDATA%\Programs\livecli\livecli.exe`

**PATH Setup:**

```powershell
# User-level PATH
setx PATH "%PATH%;%LOCALAPPDATA%\Programs\livecli"

# System-level PATH (requires admin)
setx /M PATH "%PATH%;%LOCALAPPDATA%\Programs\livecli"
```

**SmartScreen:**
Windows may show SmartScreen warning. Click "More info" → "Run anyway"

---

## Testing Cross-Platform Builds

### Manual Testing Checklist

For each platform binary:

- [ ] Binary executes without errors
- [ ] `--help` flag works
- [ ] `--version` flag works (if implemented)
- [ ] Core commands work (`exec`, `setup`, `chat`, etc.)
- [ ] API key configuration works
- [ ] OS detection works correctly
- [ ] Package manager detection works (Linux)

### Automated Testing

```bash
# Run verification script
./verify-builds.sh

# Run Go tests
make test
```

### Docker Testing (Linux)

Test Linux binaries in Docker containers:

```bash
# Ubuntu
docker run -v $(pwd)/build:/app ubuntu:22.04 /app/livecli-linux-amd64 --help

# Alpine
docker run -v $(pwd)/build:/app alpine:latest /app/livecli-linux-amd64 --help

# Fedora
docker run -v $(pwd)/build:/app fedora:latest /app/livecli-linux-amd64 --help
```

---

## Binary Size Optimization

Current binary sizes (approximate):

- Linux (amd64): ~8.6 MB
- Linux (arm64): ~8.0 MB
- macOS (amd64): ~8.7 MB
- macOS (arm64): ~8.2 MB
- Windows (amd64): ~9.0 MB
- Windows (arm64): ~8.3 MB

### Further Optimization

```bash
# Use UPX compression (reduces size by ~60%)
upx --best --lzma build/livecli-*

# Note: Some antivirus software may flag UPX-compressed binaries
```

---

## Security & Verification

### Checksums

Always verify downloaded binaries:

```bash
# Linux/macOS
sha256sum -c build/checksums.txt

# Windows
certutil -hashfile livecli-windows-amd64.exe SHA256
```

### Code Signing (Future)

- **macOS**: Sign with Apple Developer Certificate
- **Windows**: Sign with Authenticode certificate

---

## Release Checklist

Before creating a new release:

1. **Update Version**

   - [ ] Update version in code (if applicable)
   - [ ] Update CHANGELOG.md
   - [ ] Update README.md

2. **Build & Test**

   - [ ] Run `make build-all`
   - [ ] Run `./verify-builds.sh`
   - [ ] Run `make test`
   - [ ] Test all binaries manually

3. **Documentation**

   - [ ] Update INSTALL.md
   - [ ] Update README.md with new features
   - [ ] Update examples

4. **Create Release**
   - [ ] Create and push git tag: `git tag v1.x.x && git push origin v1.x.x`
   - [ ] Verify GitHub Actions completes successfully
   - [ ] Test installation scripts
   - [ ] Announce release

---

## Troubleshooting

### Build Fails

**Issue:** `go build` fails with errors

**Solution:**

```bash
# Clean and rebuild
make clean
go mod tidy
make build-all
```

### Binary Doesn't Run

**Issue:** Binary exits immediately or shows errors

**Solution:**

1. Check if it's executable (Linux/macOS): `chmod +x livecli-*`
2. Verify correct architecture for your system
3. Check for missing dependencies (rare with Go binaries)

### Cross-Compilation Issues

**Issue:** Can't build for a specific platform

**Solution:**

```bash
# Install cross-compilation tools
go install golang.org/dl/go1.21.0@latest

# Ensure GOOS and GOARCH are set correctly
GOOS=windows GOARCH=amd64 go build ...
```

---

## Contributing

When adding platform-specific code:

1. Use `runtime.GOOS` and `runtime.GOARCH` for detection
2. Test on all platforms before merging
3. Update this documentation
4. Add to build automation

---

## References

- [Go Cross-Compilation](https://golang.org/doc/install/source#environment)
- [GOOS and GOARCH Values](https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63)
- [GitHub Actions for Go](https://github.com/actions/setup-go)

---

## Support

**Issues:** [GitHub Issues](https://github.com/mohd-aquib-razi-7853/livecli/issues)

**Discussions:** [GitHub Discussions](https://github.com/mohd-aquib-razi-7853/livecli/discussions)
