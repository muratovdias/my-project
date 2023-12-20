package models

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/muratovdias/my-project/internal/repository/postgres"
)

type Account struct {
	ID        int64              `db:"id" json:"id"`
	Owner     string             `db:"owner" json:"owner"`
	Balance   int64              `db:"balance" json:"balance"`
	Currency  string             `db:"currency" json:"currency"`
	CreatedAt pgtype.Timestamptz `db:"created_at" json:"created_at"`
}

func (a *Account) ToCreateAccount() postgres.CreateAccountParams {
	return postgres.CreateAccountParams{
		Owner:    a.Owner,
		Balance:  a.Balance,
		Currency: a.Currency,
	}
}
