package db

import (
	"context"
	"testing"

	"github.com/itsmeberwyn/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) Entry {
    account1 := createRandomAccount(t)

    arg := CreateEntryParams {
        AccountID: account1.ID,
        Amount: util.RandomMoney(),
    }

    entry, err := testQueries.CreateEntry(context.Background(), arg)
    require.NoError(t, err)
    require.NotEmpty(t, entry)
    require.Equal(t, arg.AccountID, entry.AccountID)
    require.Equal(t, arg.Amount, entry.Amount)

    return entry
}

func TestCreateEntry(t *testing.T) {
    createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
    entry1 := createRandomEntry(t)
    entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
    require.NoError(t, err)
    require.NotEmpty(t, entry2)
    require.Equal(t, entry1.AccountID, entry2.AccountID)
    require.Equal(t, entry1.Amount, entry2.Amount)
}

func TestGetEntries(t *testing.T) {
    for i:=0; i<10; i++ {
        createRandomEntry(t)
    }

    arg := ListEntriesParams {
        Limit: 5,
        Offset: 5,
    }

    entries, err := testQueries.ListEntries(context.Background(), arg)
    require.NoError(t, err)
    require.Len(t, entries, 5)

    for _, entry := range entries {
        require.NotEmpty(t, entry)
    }
}

