package service

import (
	"context"
	"github.com/muratovdias/my-project/internal/repository"
	"github.com/muratovdias/my-project/internal/repository/postgres"
	"github.com/muratovdias/my-project/internal/service/account"
	"github.com/muratovdias/my-project/internal/service/entry"
	"golang.org/x/exp/slog"
)

type Account interface {
	CreateAccount(ctx context.Context, arg postgres.CreateAccountParams) (postgres.Account, error)
	DeleteAccount(ctx context.Context, id int64) error
	UpdateAccount(ctx context.Context, arg postgres.UpdateAccountParams) (postgres.Account, error)
	GetAccount(ctx context.Context, id int64) (postgres.Account, error)
	ListAccounts(ctx context.Context, arg postgres.ListAccountsParams) ([]postgres.Account, error)
	AddAccountBalance(ctx context.Context, arg postgres.AddAccountBalanceParams) (postgres.Account, error)
}

type Entry interface {
	CreateEntry(ctx context.Context, arg postgres.CreateEntryParams) (postgres.Entry, error)
	GetEntry(ctx context.Context, id int64) (postgres.Entry, error)
	ListEntries(ctx context.Context, arg postgres.ListEntriesParams) ([]postgres.Entry, error)
}

type Service struct {
	Account
	Entry
}

func NewService(repo *repository.Repository, logger *slog.Logger) *Service {
	return &Service{
		Account: account.NewService(repo.AccountRepo, logger),
		Entry:   entry.NewService(repo.EntryRepo, logger),
	}
}
