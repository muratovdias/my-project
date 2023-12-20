package account

import (
	"context"
	"github.com/muratovdias/my-project/internal/repository"
	"github.com/muratovdias/my-project/internal/repository/postgres"
	"golang.org/x/exp/slog"
)

type Service struct {
	repo   repository.AccountRepo
	logger *slog.Logger
}

func (s *Service) CreateAccount(ctx context.Context, arg postgres.CreateAccountParams) (postgres.Account, error) {
	return s.repo.CreateAccount(ctx, arg)
}

func (s *Service) DeleteAccount(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}

func (s *Service) UpdateAccount(ctx context.Context, arg postgres.UpdateAccountParams) (postgres.Account, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) GetAccount(ctx context.Context, id int64) (postgres.Account, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) ListAccounts(ctx context.Context, arg postgres.ListAccountsParams) ([]postgres.Account, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) AddAccountBalance(ctx context.Context, arg postgres.AddAccountBalanceParams) (postgres.Account, error) {
	//TODO implement me
	panic("implement me")
}

func NewService(repo repository.AccountRepo, logger *slog.Logger) *Service {
	return &Service{
		repo:   repo,
		logger: logger,
	}
}
