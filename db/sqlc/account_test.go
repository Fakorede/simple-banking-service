package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/fakorede/simple-banking-service/utils"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	acc := CreateAccountParams {
		Owner: utils.RandomOwner(),
		Balance: utils.RandomAmount(),
		Currency: utils.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), acc)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, acc.Owner, account.Owner)
	require.Equal(t, acc.Balance, account.Balance)
	require.Equal(t, acc.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	// create account
	acc := createRandomAccount(t)

	account, err := testQueries.GetAccount(context.Background(), acc.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, acc.ID, account.ID)
	require.Equal(t, acc.Owner, account.Owner)
	require.Equal(t, acc.Balance, account.Balance)
	require.Equal(t, acc.Currency, account.Currency)

	require.WithinDuration(t, acc.CreatedAt, account.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	acc := createRandomAccount(t)

	arg := UpdateAccountParams {
		ID: acc.ID,
		Balance: utils.RandomAmount(),
	}

	account, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, acc.ID, account.ID)
	require.Equal(t, acc.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, acc.Currency, account.Currency)

	require.WithinDuration(t, acc.CreatedAt, account.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	acc := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), acc.ID)
	require.NoError(t, err)

	account, err:= testQueries.GetAccount(context.Background(), acc.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account)
}

func TestListAccounts(t *testing.T) {
	for i :=0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAcountsParams {
		Limit: 5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAcounts(context.Background(), arg)
	require.NoError(t, err)

	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}