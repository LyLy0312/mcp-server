package tools

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/kimnt93/xbrain_mcp_server/internal/config"
	"github.com/mark3labs/mcp-go/mcp"
)

func GetStockRecommendHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	apiUrl := config.XnoV1ApiUrl + "/v2/tradingmap"

	resp, err := http.Get(apiUrl)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to fetch stock data: %v", err)), nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to read response: %v", err)), nil
	}

	// Trả về dữ liệu dạng JSON text
	return mcp.NewToolResultText(string(body)), nil
}
