package db

import (
	"context"
	"testing"

	"github.com/itsmeberwyn/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransaction(t *testing.T) Transfer {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	transaction, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transaction)
	require.Equal(t, arg.FromAccountID, transaction.FromAccountID)
	require.Equal(t, arg.ToAccountID, transaction.ToAccountID)
	require.Equal(t, arg.Amount, transaction.Amount)

	return transaction
}

func TestCreateTransaction(t *testing.T) {
	createRandomTransaction(t)
}

func TestGetTransaction(t *testing.T) {
	transaction1 := createRandomTransaction(t)
	transaction2, err := testQueries.GetTransfer(context.Background(), transaction1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transaction2)
	require.Equal(t, transaction1.FromAccountID, transaction2.FromAccountID)
	require.Equal(t, transaction1.ToAccountID, transaction2.ToAccountID)
	require.Equal(t, transaction1.Amount, transaction2.Amount)
}

func TestGetTransactions(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomTransaction(t)
	}

	arg := ListTransfersParams{
		Limit:  5,
		Offset: 5,
	}

    transfers, err := testQueries.ListTransfers(context.Background(), arg)
    require.NoError(t, err)
    require.Len(t, transfers, 5)

    for _, transfer := range transfers {
        require.NotEmpty(t, transfer)
    }
}
