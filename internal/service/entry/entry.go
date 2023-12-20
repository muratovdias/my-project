package entry

import (
	"context"
	"github.com/muratovdias/my-project/internal/repository"
	"github.com/muratovdias/my-project/internal/repository/postgres"
	"golang.org/x/exp/slog"
)

type Service struct {
	repo   repository.EntryRepo
	logger *slog.Logger
}

func (s Service) CreateEntry(ctx context.Context, arg postgres.CreateEntryParams) (postgres.Entry, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) GetEntry(ctx context.Context, id int64) (postgres.Entry, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) ListEntries(ctx context.Context, arg postgres.ListEntriesParams) ([]postgres.Entry, error) {
	//TODO implement me
	panic("implement me")
}

func NewService(repo repository.EntryRepo, logger *slog.Logger) *Service {
	return &Service{
		repo:   repo,
		logger: logger,
	}
}
