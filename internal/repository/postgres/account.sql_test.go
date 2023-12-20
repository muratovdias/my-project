package postgres

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestQueries_CreateAccount(t *testing.T) {
	arg := CreateAccountParams{
		Owner:    "tom",
		Balance:  100,
		Currency: "USD",
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
}
