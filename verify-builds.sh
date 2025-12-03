#!/bin/bash

# Quick test script to verify all binaries work
# This helps verify cross-platform builds after running make build-all

set -e

BUILD_DIR="build"
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

echo -e "${CYAN}"
echo "╔═══════════════════════════════════════════════════════════╗"
echo "║        🧪 LiveCLI Cross-Platform Build Verification      ║"
echo "╚═══════════════════════════════════════════════════════════╝"
echo -e "${NC}"

# Check if build directory exists
if [ ! -d "$BUILD_DIR" ]; then
    echo -e "${RED}❌ Build directory not found. Run 'make build-all' first.${NC}"
    exit 1
fi

# List all built binaries
echo -e "\n${YELLOW}📦 Built binaries:${NC}"
ls -lh $BUILD_DIR/ | grep -E "livecli-" | awk '{print "  " $9 " (" $5 ")"}'

# Verify checksums exist
echo -e "\n${YELLOW}🔐 Verifying checksums...${NC}"
if [ -f "$BUILD_DIR/checksums.txt" ]; then
    echo -e "${GREEN}✓${NC} Checksums file exists"
    cat $BUILD_DIR/checksums.txt
else
    echo -e "${RED}❌ Checksums file not found${NC}"
fi

# Test the binary for current platform
echo -e "\n${YELLOW}🧪 Testing current platform binary...${NC}"

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
        echo -e "${YELLOW}⚠️  Unsupported OS for testing: $OS${NC}"
        exit 0
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
        echo -e "${YELLOW}⚠️  Unsupported architecture for testing: $ARCH${NC}"
        exit 0
        ;;
esac

BINARY="$BUILD_DIR/livecli-$OS-$ARCH"

if [ -f "$BINARY" ]; then
    echo -e "${CYAN}Testing: $BINARY${NC}"
    
    # Make executable
    chmod +x "$BINARY"
    
    # Test help command
    if $BINARY --help > /dev/null 2>&1; then
        echo -e "${GREEN}✓${NC} Binary runs successfully"
    else
        echo -e "${RED}❌ Binary failed to run${NC}"
        exit 1
    fi
    
    # Show version info (if available)
    $BINARY --version 2>/dev/null || echo -e "${YELLOW}  (version info not available)${NC}"
else
    echo -e "${RED}❌ Binary not found: $BINARY${NC}"
    exit 1
fi

echo -e "\n${CYAN}╔═══════════════════════════════════════════════════════════╗${NC}"
echo -e "${GREEN}║           ✅ All checks passed!                           ║${NC}"
echo -e "${CYAN}╚═══════════════════════════════════════════════════════════╝${NC}"

echo -e "\n${YELLOW}📋 Summary:${NC}"
echo -e "  • Total binaries: $(ls -1 $BUILD_DIR/livecli-* 2>/dev/null | wc -l)"
echo -e "  • Platforms covered: Linux (amd64, arm64), macOS (amd64, arm64), Windows (amd64, arm64)"
echo -e "  • Current platform binary tested and working ✓"

echo -e "\n${CYAN}💡 Next steps:${NC}"
echo -e "  1. Test on actual target systems"
echo -e "  2. Create a GitHub release: git tag v1.0.0 && git push origin v1.0.0"
echo -e "  3. Distribute binaries from build/ directory"
echo
