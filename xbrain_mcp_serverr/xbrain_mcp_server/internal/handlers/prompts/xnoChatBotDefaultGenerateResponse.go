package prompts

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"time"
)

const xnoChatBotDefaultGenerateResponse = `
You are Mr.XNO, a helpful AI assistant developed by the XNO Team. Follow the rules and facts below, and use the user's request and chat summary to generate a response:

**Rules:**
- Always respond in a friendly manner and never use offensive language.
- Do not fabricate information or provide false details. If you don't know the answer, say "I don't know" or "I don't have enough information to answer your request."
- By default, respond in **Vietnamese**.
- If the user's message is in another language, respond in that language.
- If the user explicitly requests a specific language, respond in that language.

**Facts:**
- Now is %s in year-month-day hour:minute:second format.
- The XNO home page is https://xno.vn.
- XNO is a platform that combines AI and human intelligence to turn investment ideas into automated trading robots. It helps users save time on analysis, optimize profits, and reduce risks. XNO is a pioneer in applying AI to stock investing in Vietnam.

**User Request**: %s  
**Chat Summary**: %s  
**Response**: `

func getStringArg(args map[string]string, key, defaultValue string) string {
	if val, exists := args[key]; exists {
		return val
	}
	return defaultValue
}

var PromptXnoChatBotDefaultGenerate = mcp.NewPrompt(
	"prompt.xno_chatbot_default_generate_response",
	mcp.WithPromptDescription("Generates a response for the XNO chatbot based on user request and chat summary."),
	mcp.WithArgument("user_request",
		mcp.RequiredArgument(),
		mcp.ArgumentDescription("User's request for the chatbot to respond to"),
	),
	mcp.WithArgument("chat_summary",
		mcp.ArgumentDescription("Summary of the chat conversation so far, to provide context for the response"),
	),
)

func PromptXnoChatBotDefaultGenerateHandler(ctx context.Context, req mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
	userRequest := getStringArg(req.Params.Arguments, "user_request", "")
	chatSummary := getStringArg(req.Params.Arguments, "chat_summary", "")

	// Fill the prompt with the current time, user request, and chat summary
	now := time.Now().Format("2006-01-02 15:04:05") // or any other preferred format
	prompt := fmt.Sprintf(xnoChatBotDefaultGenerateResponse, now, userRequest, chatSummary)
	return &mcp.GetPromptResult{
		Messages: []mcp.PromptMessage{
			{
				Role:    mcp.RoleUser,
				Content: mcp.NewTextContent(prompt),
			},
		},
	}, nil
}
