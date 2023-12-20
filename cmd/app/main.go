package main

import (
	"fmt"
	"github.com/joho/godotenv"
	app2 "github.com/muratovdias/my-project/internal/app"
	"github.com/muratovdias/my-project/internal/config"
	db "github.com/muratovdias/my-project/pkg/db/postgres"
	"golang.org/x/exp/slog"
	"log"
	"os"
)

const (
	devMode  = "dev"
	prodMode = "prod"
)

func main() {
	if err := prepareENV(); err != nil {
		log.Fatalf("prepare env error: %s", err.Error())
	}

	cfg := config.PrepareConfig()
	fmt.Println(cfg)

	logger := setUpLogger(cfg.Mode)

	dbConn, err := db.NewPostgresDB(cfg.DSN)
	if err != nil {
		log.Fatalf("database connection failed %s", err.Error())
	}

	app := app2.InitializeApp(dbConn, logger, cfg)

	app.Start()
}

func setUpLogger(mode string) *slog.Logger {
	var logger *slog.Logger

	switch mode {
	case devMode:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelDebug,
			}),
		)
	case prodMode:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelInfo,
			}),
		)
	}

	return logger
}

func prepareENV() error {
	return godotenv.Load()
}
