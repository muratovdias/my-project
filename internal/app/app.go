package app

import (
	"context"
	"database/sql"
	"github.com/muratovdias/my-project/internal/config"
	v1 "github.com/muratovdias/my-project/internal/delivery/http/v1"
	"github.com/muratovdias/my-project/internal/repository"
	"github.com/muratovdias/my-project/internal/repository/postgres"
	"github.com/muratovdias/my-project/internal/service"
	"golang.org/x/exp/slog"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
)

type App struct {
	server *http.Server
}

func InitializeApp(db *sql.DB, logger *slog.Logger, cfg *config.Config) *App {

	queries := postgres.New(db)

	repo := repository.NewRepository(queries)

	serv := service.NewService(repo, logger)

	handler := v1.NewHandler(serv, logger)

	return &App{
		server: &http.Server{
			Addr:    ":" + cfg.Port,
			Handler: handler.InitRouter(),
		},
	}
}

func (a *App) Start() {
	eg := &errgroup.Group{}

	eg.Go(func() error {
		return a.server.ListenAndServe()
	})
	// TODO: graceful shutdown
	err := eg.Wait()
	if err != nil {
		err := a.server.Shutdown(context.Background())
		if err != nil {
			log.Fatal(err)
		}
	}
}
