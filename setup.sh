#!/bin/bash

# LiveCLI Setup Script
# This script helps you set up LiveCLI and configure your OpenAI API key

set -e

echo "╔═══════════════════════════════════════════════════════════╗"
echo "║          🚀 LiveCLI Setup Script                          ║"
echo "╚═══════════════════════════════════════════════════════════╝"
echo ""

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Error: Go is not installed. Please install Go 1.21 or higher."
    echo "Visit: https://golang.org/doc/install"
    exit 1
fi

echo "✓ Go is installed: $(go version)"
echo ""

# Download dependencies
echo "📦 Downloading dependencies..."
go mod download
go mod tidy
echo "✓ Dependencies downloaded"
echo ""

# Build the application
echo "🔨 Building LiveCLI..."
go build -o livecli main.go
echo "✓ Build successful!"
echo ""

# Check for OpenAI API key
if [ -z "$OPENAI_API_KEY" ]; then
    echo "⚠️  OpenAI API key not found in environment variables"
    echo ""
    read -p "Would you like to set it now? (y/n): " answer
    
    if [ "$answer" = "y" ] || [ "$answer" = "Y" ]; then
        read -p "Enter your OpenAI API key: " api_key
        
        # Determine shell config file
        if [ -n "$ZSH_VERSION" ]; then
            SHELL_CONFIG="$HOME/.zshrc"
        elif [ -n "$BASH_VERSION" ]; then
            SHELL_CONFIG="$HOME/.bashrc"
        else
            SHELL_CONFIG="$HOME/.profile"
        fi
        
        echo "" >> "$SHELL_CONFIG"
        echo "# LiveCLI OpenAI API Key" >> "$SHELL_CONFIG"
        echo "export OPENAI_API_KEY=\"$api_key\"" >> "$SHELL_CONFIG"
        
        echo "✓ API key added to $SHELL_CONFIG"
        echo "  Run: source $SHELL_CONFIG"
        echo "  Or restart your terminal to use the key"
    else
        echo ""
        echo "You can set your API key later by:"
        echo "  export OPENAI_API_KEY='your-key-here'"
        echo "Or use the --api-key flag when running commands"
    fi
else
    echo "✓ OpenAI API key is set"
fi

echo ""
echo "╔═══════════════════════════════════════════════════════════╗"
echo "║          🎉 Setup Complete!                               ║"
echo "╚═══════════════════════════════════════════════════════════╝"
echo ""
echo "Next steps:"
echo "  1. Test the installation:"
echo "     ./livecli"
echo ""
echo "  2. Execute a command:"
echo "     ./livecli exec 'ls -la'"
echo ""
echo "  3. Ask AI a question:"
echo "     ./livecli ask 'How do I find large files?'"
echo ""
echo "  4. Start interactive mode:"
echo "     ./livecli interactive"
echo ""
echo "  5. Optional - Install globally:"
echo "     sudo mv livecli /usr/local/bin/"
echo ""
echo "For more information, see README.md"
echo ""
