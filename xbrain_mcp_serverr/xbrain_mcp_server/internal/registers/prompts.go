package registers

import (
	"github.com/kimnt93/xbrain_mcp_server/internal/handlers/prompts"
	"github.com/mark3labs/mcp-go/server"
)

func RegisterPrompts(s *server.MCPServer) {
	s.AddPrompt(
		prompts.PromptXnoChatBotDefaultGenerate,
		prompts.PromptXnoChatBotDefaultGenerateHandler,
	)
	s.AddPrompt(
		prompts.PromptConversationSummary,
		prompts.PromptConversationSummaryHandler,
	)
}
