package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/scortier/gopherbank/util"
	"github.com/stretchr/testify/require" // for test result assertion
)

// createRandomAccount creates a random account for testing.
// It generates random account parameters, creates an account, and verifies the results.
func createRandomAccount(t *testing.T) Account {
	user := createRandomUser(t)

	arg := CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err, "Error creating a random account")
	require.NotEmpty(t, account, "Created account is empty")

	// Verify the account details match the input parameters.
	require.Equal(t, arg.Owner, account.Owner, "Owner mismatch")
	require.Equal(t, arg.Balance, account.Balance, "Balance mismatch")
	require.Equal(t, arg.Currency, account.Currency, "Currency mismatch")

	// Verify that account ID and CreatedAt are not zero.
	require.NotZero(t, account.ID, "ID is zero")
	require.NotZero(t, account.CreatedAt, "CreatedAt is zero")

	return account
}

// TestCreateAccount tests the CreateAccount function.
func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

// TestGetAccount tests the GetAccount function.
func TestGetAccount(t *testing.T) {
	// Create a random account.
	account1 := createRandomAccount(t)

	// Retrieve the account by ID.
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err, "Error retrieving account by ID")
	require.NotEmpty(t, account2, "Retrieved account is empty")

	// Verify that the retrieved account details match the original account.
	require.Equal(t, account1.Owner, account2.Owner, "Owner mismatch")
	require.Equal(t, account1.Balance, account2.Balance, "Balance mismatch")
	require.Equal(t, account1.Currency, account2.Currency, "Currency mismatch")

	// Verify that account ID and CreatedAt are not zero.
	require.NotZero(t, account2.ID, "ID is zero")
	require.NotZero(t, account2.CreatedAt, "CreatedAt is zero")
}

// TestUpdateAccount tests the UpdateAccount function.
func TestUpdateAccount(t *testing.T) {
	// Create a random account.
	account1 := createRandomAccount(t)

	// Generate random balance for update.
	arg := UpdateAccountParams{
		ID:      account1.ID,
		Balance: util.RandomMoney(),
	}

	// Update the account balance.
	account2, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err, "Error updating account balance")
	require.NotEmpty(t, account2, "Updated account is empty")

	// Verify that the updated account details match the original account except for the balance.
	require.Equal(t, account1.ID, account2.ID, "ID mismatch")
	require.Equal(t, account1.Owner, account2.Owner, "Owner mismatch")
	require.Equal(t, arg.Balance, account2.Balance, "Updated balance mismatch")
	require.Equal(t, account1.Currency, account2.Currency, "Currency mismatch")

	// Verify that CreatedAt remains within a second of the original account's CreatedAt.
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)

	// Verify that account ID and CreatedAt are not zero.
	require.NotZero(t, account2.ID, "ID is zero")
	require.NotZero(t, account2.CreatedAt, "CreatedAt is zero")
}

// TestDeleteAccount tests the DeleteAccount function.
func TestDeleteAccount(t *testing.T) {
	// Create a random account.
	account1 := createRandomAccount(t)

	// Delete the account.
	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err, "Error deleting account")

	// Attempt to retrieve the deleted account should result in an error.
	_, err = testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err, "Expected error when retrieving deleted account")
	require.EqualError(t, err, sql.ErrNoRows.Error(), "Error message mismatch")
}

// TestListAccounts tests the ListAccounts function.
func TestListAccounts(t *testing.T) {
	var lastAccount Account
	for i := 0; i < 10; i++ {
		lastAccount = createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Owner:  lastAccount.Owner,
		Limit:  5,
		Offset: 0,
	}

	// Retrieve a list of accounts matching the last created account's owner.
	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err, "Error listing accounts")
	require.NotEmpty(t, accounts, "List of accounts is empty")

	// Verify that all listed accounts have the same owner as the last created account.
	for _, account := range accounts {
		require.NotEmpty(t, account, "Listed account is empty")
		require.Equal(t, lastAccount.Owner, account.Owner, "Owner mismatch")
	}
}
