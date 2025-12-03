#!/bin/bash

# LiveCLI Installation Script for Linux & macOS
# Usage: curl -fsSL https://raw.githubusercontent.com/yourusername/livecli/main/install.sh | bash

set -e

BINARY_NAME="livecli"
INSTALL_DIR="/usr/local/bin"
REPO_URL="https://github.com/yourusername/livecli"  # Update with your actual repo
VERSION="latest"

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

echo -e "${CYAN}"
echo "╔═══════════════════════════════════════════════════════════╗"
echo "║        🚀 LiveCLI Installation Script                    ║"
echo "╚═══════════════════════════════════════════════════════════╝"
echo -e "${NC}"

# Detect OS and Architecture
detect_platform() {
    OS=$(uname -s | tr '[:upper:]' '[:lower:]')
    ARCH=$(uname -m)
    
    case "$OS" in
        linux*)
            OS="linux"
            ;;
        darwin*)
            OS="darwin"
            ;;
        *)
            echo -e "${RED}❌ Unsupported operating system: $OS${NC}"
            exit 1
            ;;
    esac
    
    case "$ARCH" in
        x86_64|amd64)
            ARCH="amd64"
            ;;
        aarch64|arm64)
            ARCH="arm64"
            ;;
        *)
            echo -e "${RED}❌ Unsupported architecture: $ARCH${NC}"
            exit 1
            ;;
    esac
    
    echo -e "${GREEN}✓${NC} Detected platform: ${CYAN}$OS-$ARCH${NC}"
}

# Download binary
download_binary() {
    BINARY_FILE="${BINARY_NAME}-${OS}-${ARCH}"
    DOWNLOAD_URL="${REPO_URL}/releases/${VERSION}/download/${BINARY_FILE}"
    
    echo -e "\n${YELLOW}📥 Downloading ${BINARY_NAME}...${NC}"
    
    # For now, if building locally, use the build directory
    if [ -f "build/${BINARY_FILE}" ]; then
        echo -e "${GREEN}✓${NC} Using local build: build/${BINARY_FILE}"
        BINARY_PATH="build/${BINARY_FILE}"
    else
        # TODO: Implement remote download when releases are available
        echo -e "${RED}❌ Binary not found. Please build first with: make build-all${NC}"
        exit 1
    fi
}

# Install binary
install_binary() {
    echo -e "\n${YELLOW}📦 Installing ${BINARY_NAME} to ${INSTALL_DIR}...${NC}"
    
    # Check if we need sudo
    if [ -w "$INSTALL_DIR" ]; then
        cp "$BINARY_PATH" "$INSTALL_DIR/$BINARY_NAME"
        chmod +x "$INSTALL_DIR/$BINARY_NAME"
    else
        echo -e "${YELLOW}⚠️  Need sudo privileges to install to ${INSTALL_DIR}${NC}"
        sudo cp "$BINARY_PATH" "$INSTALL_DIR/$BINARY_NAME"
        sudo chmod +x "$INSTALL_DIR/$BINARY_NAME"
    fi
    
    echo -e "${GREEN}✓${NC} ${BINARY_NAME} installed successfully!"
}

# Verify installation
verify_installation() {
    echo -e "\n${YELLOW}🔍 Verifying installation...${NC}"
    
    if command -v $BINARY_NAME &> /dev/null; then
        VERSION_OUTPUT=$($BINARY_NAME --version 2>&1 || echo "version check not available")
        echo -e "${GREEN}✓${NC} ${BINARY_NAME} is ready to use!"
        echo -e "\n${CYAN}Try running:${NC} ${BINARY_NAME} --help"
    else
        echo -e "${RED}❌ Installation verification failed${NC}"
        exit 1
    fi
}

# Setup instructions
show_setup_info() {
    echo -e "\n${CYAN}╔═══════════════════════════════════════════════════════════╗${NC}"
    echo -e "${CYAN}║           ✅ Installation Complete!                       ║${NC}"
    echo -e "${CYAN}╚═══════════════════════════════════════════════════════════╝${NC}"
    
    echo -e "\n${YELLOW}📝 Next Steps:${NC}"
    echo -e "  1. Set your OpenAI API key:"
    echo -e "     ${CYAN}export OPENAI_API_KEY='your-api-key'${NC}"
    echo -e "     Add this to your ~/.bashrc or ~/.zshrc for persistence"
    echo
    echo -e "  2. Start using LiveCLI:"
    echo -e "     ${CYAN}livecli --help${NC}"
    echo -e "     ${CYAN}livecli setup \"docker\"${NC}"
    echo -e "     ${CYAN}livecli chat${NC}"
    echo
    echo -e "${GREEN}🎉 Happy coding!${NC}\n"
}

# Main installation flow
main() {
    detect_platform
    download_binary
    install_binary
    verify_installation
    show_setup_info
}

main
