package db

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	account1 := createTestAccount(t)
	account2 := createTestAccount(t)
	fmt.Println(">> before:", account1.Balance, account2.Balance)
	// run n concurrent transfer transactions
	n := 2
	amount := int64(10)
	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		txName := fmt.Sprintf("tx-%d", i+1)
		go func(i int) {
			ctx := context.WithValue(context.Background(), txKey, txName)
			result, err := store.TransferTx(ctx, TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			})
			log.Printf("%+v", result)
			errs <- err
			results <- result
		}(i)
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// check transfer
		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, account1.ID, transfer.FromAccountID)
		require.Equal(t, account2.ID, transfer.ToAccountID)

		transfer, err = store.GetTransfer(context.Background(), transfer.ID)
		require.NotEmpty(t, transfer)

		// check entry
		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, -amount, fromEntry.Amount)

		entry, err := store.GetEntry(context.Background(), fromEntry.ID)
		require.NotEmpty(t, entry)

		// check accounts
		fromAccount := result.FromAccount
		require.NotEmpty(t, fromAccount)
		require.Equal(t, account1.ID, fromAccount.ID)

		toAccount := result.ToAccount
		require.NotEmpty(t, toAccount)
		require.Equal(t, account2.ID, toAccount.ID)
	}

}
