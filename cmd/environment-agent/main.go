package main

import (
	"context"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/dcm-project/environment-agent/internal/api/server"
	"github.com/dcm-project/environment-agent/internal/apiserver"
	"github.com/dcm-project/environment-agent/internal/config"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()
	os.Exit(run(ctx))
}

func run(ctx context.Context) int {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("Environment Agent starting")

	cfg, err := config.Load()
	if err != nil {
		logger.Error("failed to load config", "error", err)
		return 1
	}
	if err := cfg.Validate(); err != nil {
		logger.Error("invalid configuration", "error", err)
		return 1
	}

	ln, err := net.Listen("tcp", cfg.Server.Address)
	if err != nil {
		logger.Error("failed to listen", "error", err, "address", cfg.Server.Address)
		return 1
	}
	defer ln.Close()

	handler := &server.Unimplemented{}
	srv := apiserver.New(cfg, logger, handler)

	if err := srv.Run(ctx, ln); err != nil {
		logger.Error("server error", "error", err)
		return 1
	}
	logger.Info("Environment Agent stopped")
	return 0
}
