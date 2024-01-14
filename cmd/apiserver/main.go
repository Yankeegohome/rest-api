package main

import (
	"context"
	"flag"
	"github.com/BurntSushi/toml"
	"log"
	"log/slog"
	"os"
	"rest-api/internal/app/apiserver"
	grpcv1Client "rest-api/internal/clients/grpc"
	"rest-api/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config1 := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config1)
	if err != nil {
		log.Fatal(err)
	}

	if err := apiserver.Start(config1); err != nil {
		log.Fatal(err)
	}
	cfg := config.MustLoad()
	log := setupLogger(cfg.Env)
	grpcClient, err := grpcv1Client.New(
		context.Background(),
		log,
		cfg.Clients.GRPC.Address,
		cfg.Clients.GRPC.Timeout,
		cfg.Clients.GRPC.RetriesCount,
	)
	if err != nil {
		log.Error("failed to init grpc client", err)
	}
	grpcClient.IsAdmin(context.Background(), 1)
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
