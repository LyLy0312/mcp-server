package resources

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/kimnt93/xbrain_mcp_server/internal/config"
	"github.com/mark3labs/mcp-go/mcp"
	"io"
	"net/http"
)

type StockInfoForRecommend struct {
	Symbol       string  `json:"MA"`
	Price        float64 `json:"GIA"`
	Industry     string  `json:"NGANH"`
	Strength     string  `json:"SUCMANH"`
	MidTerm      string  `json:"TRUNGHAN"`
	LongTerm     string  `json:"DAIHAN"`
	RS           int     `json:"RS"`
	Strategy     string  `json:"CHIENLUOC"`
	Candles      string  `json:"candles"`
	Pattern      string  `json:"pattern"`
	AvgRating    float64 `json:"DG_bq"`
	BAT          float64 `json:"BAT"`
	Score        int     `json:"diemBinhquan"`
	BullVolume   int64   `json:"bulVol"`
	BearVolume   int64   `json:"bearVol"`
	RRG          string  `json:"rrg"`
	CommentTA    string  `json:"cmtTA"`
	SignalSMC    string  `json:"signalSMC"`
	AITrend      string  `json:"AiTrend"`
	AIPredict20D float64 `json:"AIPredict20d"`
	Liquidity    int64   `json:"THANHKHOAN"`
	Vol1DvsAvg   float64 `json:"KL1KLTB"`
	ShortUp      int     `json:"shortup"`
	MidUp        int     `json:"midup"`
	LongUp       int     `json:"longup"`
}

// TradingMapResponse Define top-level struct to match full JSON response
type TradingMapResponse struct {
	Stocks []StockInfoForRecommend `json:"stocks"`
}

var ResourceStockRecommend = mcp.NewResource(
	"stock://trading-map",
	"Stock Analysis : Recommendations",
	mcp.WithResourceDescription("Fetches stock recommendations from the trading map."),
	mcp.WithMIMEType("application/json"),
)

func ResourceStockRecommendHandler(ctx context.Context, req mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
	apiUrl := config.XnoV1ApiUrl + "/v2/tradingmap"
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

	var data TradingMapResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	out, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to re-encode response: %w", err)
	}

	return []mcp.ResourceContents{
		mcp.TextResourceContents{
			URI:      req.Params.URI,
			MIMEType: "application/json",
			Text:     string(out),
		},
	}, nil
}
