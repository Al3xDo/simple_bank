package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/Al3xDo/simple_bank/util"

	"github.com/stretchr/testify/require"
)

func createEntries(t *testing.T) Entry {
	account := createRandomAccount(t)
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomAmount(account.Balance),
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, entry.AccountID, account.ID)
	require.Equal(t, entry.Amount, arg.Amount)
	require.NotZero(t, entry.CreatedAt)
	require.NotZero(t, entry.ID)

	return entry
}

func TestCreateEntry(t *testing.T) {
	createEntries(t)
}
func TestGetEntry(t *testing.T) {
	entry1 := createEntries(t)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}
func TestGetListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		createEntries(t)
	}
	arg := ListEntriessParams{
		Limit:  5,
		Offset: 5,
	}
	entries, err := testQueries.ListEntriess(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)
	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}

func TestDeleteEntries(t *testing.T) {
	entry1 := createEntries(t)
	err := testQueries.DeleteEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, entry2)
}
