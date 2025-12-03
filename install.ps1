# LiveCLI Installation Script for Windows
# Usage: 
#   PowerShell -ExecutionPolicy Bypass -File install.ps1
#   Or: iwr -useb https://raw.githubusercontent.com/mohd-aquib-razi-7853/livecli/main/install.ps1 | iex

param(
    [string]$InstallDir = "$env:LOCALAPPDATA\Programs\livecli",
    [string]$Version = "latest"
)

$ErrorActionPreference = "Stop"

$BinaryName = "livecli.exe"
$RepoUrl = "https://github.com/mohd-aquib-razi-7853/livecli" # Update with your actual repo

# Colors for output
function Write-ColorOutput {
    param(
        [string]$Message,
        [string]$Color = "White"
    )
    Write-Host $Message -ForegroundColor $Color
}

function Write-Header {
    Write-ColorOutput "`n╔═══════════════════════════════════════════════════════════╗" "Cyan"
    Write-ColorOutput "║        🚀 LiveCLI Installation Script                    ║" "Cyan"
    Write-ColorOutput "╚═══════════════════════════════════════════════════════════╝" "Cyan"
}

function Detect-Platform {
    Write-ColorOutput "`n✓ Detecting platform..." "Green"
    
    $arch = if ([Environment]::Is64BitOperatingSystem) {
        if ([System.Runtime.InteropServices.RuntimeInformation]::ProcessArchitecture -eq 'Arm64') {
            "arm64"
        } else {
            "amd64"
        }
    } else {
        Write-ColorOutput "❌ 32-bit Windows is not supported" "Red"
        exit 1
    }
    
    Write-ColorOutput "✓ Detected: Windows-$arch" "Green"
    return "windows-$arch"
}

function Download-Binary {
    param([string]$Platform)
    
    $binaryFile = "livecli-$Platform.exe"
    $buildPath = Join-Path $PSScriptRoot "build\$binaryFile"
    
    Write-ColorOutput "`n📥 Looking for binary..." "Yellow"
    
    # Check if building locally
    if (Test-Path $buildPath) {
        Write-ColorOutput "✓ Using local build: $buildPath" "Green"
        return $buildPath
    } else {
        # TODO: Implement remote download when releases are available
        Write-ColorOutput "❌ Binary not found. Please build first with: make build-all" "Red"
        Write-ColorOutput "   (Requires WSL or Git Bash with make installed)" "Yellow"
        exit 1
    }
}

function Install-Binary {
    param([string]$SourcePath)
    
    Write-ColorOutput "`n📦 Installing livecli..." "Yellow"
    
    # Create installation directory if it doesn't exist
    if (-not (Test-Path $InstallDir)) {
        New-Item -ItemType Directory -Path $InstallDir -Force | Out-Null
    }
    
    $destinationPath = Join-Path $InstallDir $BinaryName
    
    # Copy binary
    Copy-Item -Path $SourcePath -Destination $destinationPath -Force
    
    Write-ColorOutput "✓ Binary copied to: $destinationPath" "Green"
    
    # Add to PATH if not already there
    $currentPath = [Environment]::GetEnvironmentVariable("Path", "User")
    if ($currentPath -notlike "*$InstallDir*") {
        Write-ColorOutput "`n⚙️  Adding to PATH..." "Yellow"
        [Environment]::SetEnvironmentVariable(
            "Path",
            "$currentPath;$InstallDir",
            "User"
        )
        Write-ColorOutput "✓ Added to PATH (restart terminal for changes to take effect)" "Green"
        $env:Path += ";$InstallDir"  # Update current session
    }
}

function Verify-Installation {
    Write-ColorOutput "`n🔍 Verifying installation..." "Yellow"
    
    $livecliPath = Join-Path $InstallDir $BinaryName
    
    if (Test-Path $livecliPath) {
        Write-ColorOutput "✓ livecli is installed successfully!" "Green"
        
        # Try to run it
        try {
            $output = & $livecliPath --help 2>&1
            if ($LASTEXITCODE -eq 0 -or $output) {
                Write-ColorOutput "✓ livecli is working correctly!" "Green"
            }
        } catch {
            Write-ColorOutput "⚠️  Installation complete but verification failed" "Yellow"
        }
    } else {
        Write-ColorOutput "❌ Installation verification failed" "Red"
        exit 1
    }
}

function Show-SetupInfo {
    Write-ColorOutput "`n╔═══════════════════════════════════════════════════════════╗" "Cyan"
    Write-ColorOutput "║           ✅ Installation Complete!                       ║" "Cyan"
    Write-ColorOutput "╚═══════════════════════════════════════════════════════════╝" "Cyan"
    
    Write-ColorOutput "`n📝 Next Steps:" "Yellow"
    Write-Host "  1. Set your OpenAI API key:"
    Write-ColorOutput "     setx OPENAI_API_KEY `"your-api-key`"" "Cyan"
    Write-Host ""
    Write-Host "  2. Restart your terminal/PowerShell window"
    Write-Host ""
    Write-Host "  3. Start using LiveCLI:"
    Write-ColorOutput "     livecli --help" "Cyan"
    Write-ColorOutput "     livecli setup `"docker`"" "Cyan"
    Write-ColorOutput "     livecli chat" "Cyan"
    Write-Host ""
    Write-ColorOutput "🎉 Happy coding!`n" "Green"
}

# Main installation flow
function Main {
    Write-Header
    
    $platform = Detect-Platform
    $binaryPath = Download-Binary -Platform $platform
    Install-Binary -SourcePath $binaryPath
    Verify-Installation
    Show-SetupInfo
}

# Run installation
try {
    Main
} catch {
    Write-ColorOutput "`n❌ Installation failed: $_" "Red"
    exit 1
}
