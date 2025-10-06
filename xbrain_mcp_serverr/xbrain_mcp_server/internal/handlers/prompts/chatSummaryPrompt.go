package prompts

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"time"
)

const chatSummaryPrompt = `
Here is a chat between a user (User) and an AI assistant (AI). Please summarize the Conversation in a few sentences. Today is ?.

Conversation: ?
Summary: `

var PromptConversationSummary = mcp.NewPrompt(
	"prompt.conversation_summary",
	mcp.WithPromptDescription("Generates a summary of the conversation between a user and an AI assistant."),
	mcp.WithArgument("conversation",
		mcp.RequiredArgument(),
		mcp.ArgumentDescription("The conversation text to summarize, formatted as a dialogue between User and AI"),
	),
)

func PromptConversationSummaryHandler(ctx context.Context, req mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
	conversation := getStringArg(req.Params.Arguments, "conversation", "")
	now := time.Now().Format("2006-01-02 15:04:05") // or any other preferred format
	prompt := fmt.Sprintf(chatSummaryPrompt, now, conversation)
	return &mcp.GetPromptResult{
		Messages: []mcp.PromptMessage{
			{
				Role:    mcp.RoleUser,
				Content: mcp.NewTextContent(prompt),
			},
		},
	}, nil
}
