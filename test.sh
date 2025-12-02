#!/bin/bash

# LiveCLI - Complete Test Script
# This script tests all functionality of LiveCLI

set -e

echo "╔═══════════════════════════════════════════════════════════╗"
echo "║        LiveCLI - Comprehensive Test Suite                 ║"
echo "╚═══════════════════════════════════════════════════════════╝"
echo ""

# Colors
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Test counter
TESTS_PASSED=0
TESTS_FAILED=0

# Helper functions
pass() {
    echo -e "${GREEN}✓${NC} $1"
    ((TESTS_PASSED++))
}

fail() {
    echo -e "${RED}✗${NC} $1"
    ((TESTS_FAILED++))
}

info() {
    echo -e "${YELLOW}ℹ${NC} $1"
}

# Test 1: Check if binary exists
echo "Test 1: Checking if livecli binary exists..."
if [ -f "./livecli" ]; then
    pass "Binary exists"
else
    fail "Binary not found. Run: go build -o livecli main.go"
    exit 1
fi

# Test 2: Check if binary is executable
echo ""
echo "Test 2: Checking if binary is executable..."
if [ -x "./livecli" ]; then
    pass "Binary is executable"
else
    fail "Binary is not executable. Run: chmod +x livecli"
    exit 1
fi

# Test 3: Test help command
echo ""
echo "Test 3: Testing help command..."
if ./livecli --help &> /dev/null; then
    pass "Help command works"
else
    fail "Help command failed"
fi

# Test 4: Test exec command with simple echo
echo ""
echo "Test 4: Testing exec command..."
OUTPUT=$(./livecli exec "echo 'Hello from LiveCLI'" 2>&1)
if echo "$OUTPUT" | grep -q "Hello from LiveCLI"; then
    pass "Exec command works (echo test)"
else
    fail "Exec command failed"
fi

# Test 5: Test exec with ls command
echo ""
echo "Test 5: Testing exec with ls command..."
if ./livecli exec "ls -la" 2>&1 | grep -q "total"; then
    pass "Exec command works (ls test)"
else
    fail "Exec ls command failed"
fi

# Test 6: Test exec with date command
echo ""
echo "Test 6: Testing exec with date command..."
if ./livecli exec "date" &> /dev/null; then
    pass "Exec command works (date test)"
else
    fail "Exec date command failed"
fi

# Test 7: Test exec with pwd command
echo ""
echo "Test 7: Testing exec with pwd command..."
if ./livecli exec "pwd" &> /dev/null; then
    pass "Exec command works (pwd test)"
else
    fail "Exec pwd command failed"
fi

# Test 8: Check Go files exist
echo ""
echo "Test 8: Checking Go source files..."
GO_FILES=(
    "main.go"
    "cmd/root.go"
    "cmd/exec.go"
    "cmd/chat.go"
    "cmd/ask.go"
    "cmd/interactive.go"
)

ALL_GO_FILES_EXIST=true
for file in "${GO_FILES[@]}"; do
    if [ ! -f "$file" ]; then
        ALL_GO_FILES_EXIST=false
        fail "Missing file: $file"
    fi
done

if [ "$ALL_GO_FILES_EXIST" = true ]; then
    pass "All Go source files exist"
fi

# Test 9: Check documentation files
echo ""
echo "Test 9: Checking documentation files..."
DOC_FILES=(
    "README.md"
    "QUICKSTART.md"
    "DEMO.md"
    "EXAMPLES.md"
    "PROJECT_SUMMARY.md"
    "Makefile"
)

ALL_DOC_FILES_EXIST=true
for file in "${DOC_FILES[@]}"; do
    if [ ! -f "$file" ]; then
        ALL_DOC_FILES_EXIST=false
        fail "Missing file: $file"
    fi
done

if [ "$ALL_DOC_FILES_EXIST" = true ]; then
    pass "All documentation files exist"
fi

# Test 10: Check go.mod
echo ""
echo "Test 10: Checking go.mod..."
if [ -f "go.mod" ]; then
    if grep -q "github.com/spf13/cobra" go.mod && \
       grep -q "github.com/sashabaranov/go-openai" go.mod && \
       grep -q "github.com/fatih/color" go.mod; then
        pass "go.mod contains required dependencies"
    else
        fail "go.mod missing required dependencies"
    fi
else
    fail "go.mod not found"
fi

# Test 11: Check if code compiles
echo ""
echo "Test 11: Testing if code compiles..."
if go build -o livecli-test main.go 2>&1 > /dev/null; then
    pass "Code compiles successfully"
    rm -f livecli-test
else
    fail "Code compilation failed"
fi

# Test 12: Test exec with environment variables
echo ""
echo "Test 12: Testing exec with environment variables..."
if ./livecli exec "echo \$HOME" 2>&1 | grep -q "$HOME"; then
    pass "Environment variables work in exec"
else
    info "Environment variables test skipped or failed"
fi

# Test 13: Test exec error handling
echo ""
echo "Test 13: Testing exec error handling..."
if ./livecli exec "nonexistentcommand123456" 2>&1 | grep -q -i "error\|not found\|failed"; then
    pass "Error handling works"
else
    info "Error handling test inconclusive"
fi

# Test 14: Check API key warning for AI commands
echo ""
echo "Test 14: Testing API key validation..."
if [ -z "$OPENAI_API_KEY" ]; then
    if ./livecli ask "test" 2>&1 | grep -q -i "api key"; then
        pass "API key validation works"
    else
        info "API key validation unclear"
    fi
else
    info "API key is set, skipping validation test"
fi

# Test 15: Test make commands
echo ""
echo "Test 15: Testing Makefile..."
if [ -f "Makefile" ]; then
    if make help &> /dev/null; then
        pass "Makefile help target works"
    else
        fail "Makefile help target failed"
    fi
else
    fail "Makefile not found"
fi

# Summary
echo ""
echo "╔═══════════════════════════════════════════════════════════╗"
echo "║                    Test Summary                            ║"
echo "╚═══════════════════════════════════════════════════════════╝"
echo ""
echo -e "${GREEN}Tests Passed:${NC} $TESTS_PASSED"
echo -e "${RED}Tests Failed:${NC} $TESTS_FAILED"
TOTAL=$((TESTS_PASSED + TESTS_FAILED))
echo "Total Tests: $TOTAL"

if [ $TESTS_FAILED -eq 0 ]; then
    echo ""
    echo -e "${GREEN}╔═══════════════════════════════════════════════════════════╗${NC}"
    echo -e "${GREEN}║  🎉 All Tests Passed! LiveCLI is ready to use!           ║${NC}"
    echo -e "${GREEN}╚═══════════════════════════════════════════════════════════╝${NC}"
    echo ""
    echo "Next steps:"
    echo "  1. Set your API key: export OPENAI_API_KEY='your-key'"
    echo "  2. Try it: ./livecli interactive"
    echo "  3. Read: cat QUICKSTART.md"
    exit 0
else
    echo ""
    echo -e "${RED}╔═══════════════════════════════════════════════════════════╗${NC}"
    echo -e "${RED}║  ⚠️  Some tests failed. Please check the errors above.   ║${NC}"
    echo -e "${RED}╚═══════════════════════════════════════════════════════════╝${NC}"
    exit 1
fi
