package registers

import (
	"github.com/kimnt93/xbrain_mcp_server/internal/handlers/tools"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func RegisterTools(s *server.MCPServer) {
	tool := mcp.NewTool("hello_world",
		mcp.WithDescription("Say hello to someone"),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Name of the person to greet"),
		),
	)

	// Add example tools
	s.AddTool(tool, tools.HelloHandler)

	// Tool Get Stock Recommend
	stockTool := mcp.NewTool("get_stock_recommend",
		mcp.WithDescription("Fetch stock recommendations from XNO Trading Map"),
	)
	s.AddTool(stockTool, tools.GetStockRecommendHandler)
}
