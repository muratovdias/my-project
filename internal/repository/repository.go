package repository

import (
	"context"
	"github.com/muratovdias/my-project/internal/repository/postgres"
)

type AccountRepo interface {
	CreateAccount(ctx context.Context, arg postgres.CreateAccountParams) (postgres.Account, error)
	DeleteAccount(ctx context.Context, id int64) error
	UpdateAccount(ctx context.Context, arg postgres.UpdateAccountParams) (postgres.Account, error)
	GetAccount(ctx context.Context, id int64) (postgres.Account, error)
	ListAccounts(ctx context.Context, arg postgres.ListAccountsParams) ([]postgres.Account, error)
	AddAccountBalance(ctx context.Context, arg postgres.AddAccountBalanceParams) (postgres.Account, error)
}

type EntryRepo interface {
	CreateEntry(ctx context.Context, arg postgres.CreateEntryParams) (postgres.Entry, error)
	GetEntry(ctx context.Context, id int64) (postgres.Entry, error)
	ListEntries(ctx context.Context, arg postgres.ListEntriesParams) ([]postgres.Entry, error)
}

type Repository struct {
	AccountRepo
	EntryRepo
}

func NewRepository(queries *postgres.Queries) *Repository {
	return &Repository{
		AccountRepo: queries,
		EntryRepo:   queries,
	}
}
