package main

import (
	"github.com/kimnt93/xbrain_mcp_server/internal/config"
	"github.com/kimnt93/xbrain_mcp_server/internal/registers"
	"github.com/mark3labs/mcp-go/server"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	// Chuyển log sang stderr để không phá MCP protocol
	log.Logger = log.Output(os.Stderr)

	// Initialize configuration
	config.InitConfig()

	// Create a new MCP server
	s := server.NewMCPServer(
		"Demo 🚀",
		"1.0.0",
		server.WithToolCapabilities(false),
	)

	// Register the tools
	log.Info().Msg("Registering tools...")
	registers.RegisterTools(s)

	log.Info().Msg("Registering resources...")
	registers.RegisterResources(s)

	log.Info().Msg("Registering prompts...")
	registers.RegisterPrompts(s)

	log.Info().Msg("Starting MCP server...")
	//startWithGracefulShutdown(s)
	if err := server.ServeStdio(s); err != nil {
		log.Err(err).Msg("Failed to start MCP server")
	} else {
		log.Info().Msg("MCP server started successfully")
	}
}

//func startWithGracefulShutdown(s *server.MCPServer) {
//	// Setup signal handling
//	sigChan := make(chan os.Signal, 1)
//	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
//
//	// Start server in goroutine
//	go func() {
//		if err := server.ServeStdio(s); err != nil {
//			log.Err(err).Msg("Failed to start MCP server")
//		}
//	}()
//
//	// Wait for shutdown signal
//	<-sigChan
//	log.Info().Msg("Received shutdown signal, stopping server...")
//
//	// Graceful shutdown with timeout
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
//
//	if err := s.Shutdown(ctx); err != nil {
//		log.Err(err).Msg("Failed to shutdown MCP server gracefully")
//	} else {
//		log.Info().Msg("MCP server stopped gracefully")
//	}
//	// Exit the application
//}
