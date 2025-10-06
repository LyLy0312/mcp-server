package registers

import (
	"github.com/kimnt93/xbrain_mcp_server/internal/handlers/resources"
	"github.com/mark3labs/mcp-go/server"
)

func RegisterResources(s *server.MCPServer) {
	s.AddResource(
		resources.ResourceStockAnalysisVolumeSpike,
		resources.ResourceStockAnalysisVolumeSpikeHandler,
	)
	s.AddResource(
		resources.ResourceStockRecommend,
		resources.ResourceStockRecommendHandler,
	)
}
