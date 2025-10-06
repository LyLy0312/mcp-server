package resources

import (
	"context"
	"fmt"
	"github.com/kimnt93/xbrain_mcp_server/internal/config"
	"github.com/mark3labs/mcp-go/mcp"
	"io"
	"net/http"
)

var ResourceStockAnalysisVolumeSpike = mcp.NewResource(
	"stock://volume-spike",
	"Stock Analysis : Volume Spike",
	mcp.WithResourceDescription("Fetches stock analysis data for volume spikes."),
	mcp.WithMIMEType("application/json"),
)

func ResourceStockAnalysisVolumeSpikeHandler(ctx context.Context, req mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
	apiUrl := config.XnoApiUrl + "/xno-platform/v1/stocks-analysis/1d/volume-spike?page=1&limit=20&order_by=day_volume&order_side=desc"
	resp, err := http.Get(apiUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check if the response status is OK
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	return []mcp.ResourceContents{
		mcp.TextResourceContents{
			URI:      req.Params.URI,
			MIMEType: "application/json",
			Text:     string(body),
		},
	}, nil
}
