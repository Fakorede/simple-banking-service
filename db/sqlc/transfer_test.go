package db

import (
	"context"
	"testing"
	"time"

	"github.com/fakorede/simple-banking-service/utils"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, accounts []Account) Transfer {
	arg := CreateTransferParams{
		FromAccountID: accounts[0].ID,
		ToAccountID:   accounts[1].ID,
		Amount:        utils.RandomAmount(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func createAccounts(t *testing.T) []Account {
	accounts := []Account{}

	for i := 0; i < 2; i++ {
		account := createRandomAccount(t)
		accounts = append(accounts, account)
	}

	return accounts
}

func TestCreateTransfer(t *testing.T) {
	accounts := createAccounts(t)

	createRandomTransfer(t, accounts)
}

func TestGetTransfer(t *testing.T) {
	accounts := createAccounts(t)
	trf := createRandomTransfer(t, accounts)

	transfer, err := testQueries.GetTransfer(context.Background(), trf.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, trf.FromAccountID, transfer.FromAccountID)
	require.Equal(t, trf.ToAccountID, transfer.ToAccountID)
	require.Equal(t, trf.Amount, transfer.Amount)
	require.Equal(t, trf.ID, transfer.ID)

	require.WithinDuration(t, trf.CreatedAt, transfer.CreatedAt, time.Second)

}

func TestListTransfers(t *testing.T) {
	accounts := createAccounts(t)

	for i := 0; i < 10; i++ {
		createRandomTransfer(t, accounts)
	}
	args := ListTransfersParams{
		FromAccountID: accounts[0].ID,
		ToAccountID:   accounts[1].ID,
		Offset:        5,
		Limit:         5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), args)
	require.NoError(t, err)

	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}

}
