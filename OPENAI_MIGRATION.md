# Migration from Gemini to OpenAI - Summary

## Overview

LiveCLI has been successfully migrated from using Google's Gemini API back to OpenAI's API. This document summarizes the changes made during the migration.

## Key Changes

### Dependencies

**go.mod**:

- **Removed**: `github.com/google/generative-ai-go v0.18.0`
- **Removed**: `google.golang.org/api v0.213.0`
- **Added**: `github.com/sashabaranov/go-openai v1.20.4`
- **Removed**: All Google Cloud and Gemini-related indirect dependencies

### Environment Variables

**API Key**:

- **Old**: `GEMINI_API_KEY`
- **New**: `OPENAI_API_KEY`

**Default Model**:

- **Old**: `gemini-2.5-flash`
- **New**: `gpt-4o-mini`

### Code Changes

#### `cmd/root.go`

- Updated API key environment variable from `GEMINI_API_KEY` to `OPENAI_API_KEY`
- Changed default model from `gemini-2.5-flash` to `gpt-4o-mini`

#### `cmd/chat.go`

- Replaced Gemini imports with OpenAI SDK imports
- Rewrote `startChatSession()` to manually maintain conversation history (OpenAI doesn't have a built-in chat session)
- Created `getOpenAIResponse()` function to handle OpenAI responses
- Updated error messages to reference OpenAI

#### `cmd/ask.go`

- Complete rewrite to use OpenAI SDK
- Simplified response handling (OpenAI has simpler response structure)
- Updated error messages to reference OpenAI

#### `cmd/interactive.go`

- Rewrote to use OpenAI SDK with manual conversation history
- Maintains conversation history using message array
- Updated error messages to reference OpenAI

#### `cmd/setup.go`

- Updated imports to use OpenAI SDK
- Rewrote `generateSetupPlan()` to use OpenAI's API
- Simplified response parsing (no need to iterate through parts)
- Updated error messages to reference OpenAI

### Configuration Files

#### `.env.example`

- Updated all references from Gemini to OpenAI
- Changed API key variable name
- Updated model examples to OpenAI models

## API Differences

### Client Initialization

- **Gemini**: Required context and `option.WithAPIKey()` during client creation
- **OpenAI**: Simple `openai.NewClient(apiKey)` call

### Chat Sessions

- **Gemini**: Stateful `ChatSession` object that maintains history
- **OpenAI**: Stateless API, requires manual message history management

### Response Format

- **Gemini**: `resp.Candidates[0].Content.Parts` (requires iteration)
- **OpenAI**: `resp.Choices[0].Message.Content` (direct access)

### System Instructions

- **Gemini**: `model.SystemInstruction` property
- **OpenAI**: System message in the messages array

### Temperature & Tokens

- **Gemini**: Set on model object before generating content
- **OpenAI**: Passed in each request

## Available Models

Users can now use the following OpenAI models:

- `gpt-4o-mini` (default) - Fast and cost-effective
- `gpt-4o` - Most capable GPT-4 model
- `gpt-4-turbo` - Optimized GPT-4
- `gpt-3.5-turbo` - Fast and affordable

## Getting Your OpenAI API Key

Users can get their OpenAI API key from:

- https://platform.openai.com/api-keys

## Migration Steps for Users

Existing users need to:

1. Get a new OpenAI API key
2. Update their environment variables from `GEMINI_API_KEY` to `OPENAI_API_KEY`
3. Rebuild the application: `go build`

## Benefits of OpenAI

1. **Mature ecosystem**: OpenAI has a well-established API with extensive documentation
2. **Model variety**: Access to GPT-4, GPT-4 Turbo, and GPT-3.5 Turbo
3. **Reliability**: Industry-leading uptime and performance
4. **Features**: Advanced capabilities like function calling, JSON mode, etc.

## Quick Start After Migration

1. Set your OpenAI API key:

   ```bash
   export OPENAI_API_KEY="your-openai-api-key"
   ```

2. Build and run:

   ```bash
   go build
   ./livecli chat
   ```

## Documentation Updates

All documentation has been updated to reflect the change to OpenAI:

- README.md
- .env.example
- All error messages in the code
- Command help text

## Testing

The migration has been tested with:

- All commands (exec, chat, ask, interactive, setup, git)
- Different models
- Error handling
- Build process on multiple platforms
