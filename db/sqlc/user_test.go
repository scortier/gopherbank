package db

import (
	"context"
	"testing"

	"github.com/scortier/gopherbank/util"
	"github.com/stretchr/testify/require" // for test result assertion
)

// createRandomAccount creates a random account for testing.
// It generates random account parameters, creates an account, and verifies the results.
func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: "secret",
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err, "Error creating a random account")
	require.NotEmpty(t, user, "Created user is empty")

	// Verify the account details match the input parameters.
	require.Equal(t, arg.Username, user.Username, "Username mismatch")
	require.Equal(t, arg.HashedPassword, user.HashedPassword, "HashedPassword mismatch")
	require.Equal(t, arg.FullName, user.FullName, "FullName mismatch")
	require.Equal(t, arg.Email, user.Email, "Email mismatch")

	// Verify that account ID and CreatedAt are not zero.
	require.NotZero(t, user.CreatedAt, "CreatedAt is zero")
	require.True(t, user.PasswordChangedAt.IsZero(), "PasswordChangedAt is not zero")

	return user
}

// TestCreateAccount tests the CreateAccount function.
func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

// TestGetAccount tests the GetAccount function.
func TestGetUser(t *testing.T) {
	// Create a random account.
	user1 := createRandomUser(t)

	// Retrieve the account by ID.
	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err, "Error retrieving account by ID")
	require.NotEmpty(t, user2, "Retrieved account is empty")

	// Verify that the retrieved account details match the original account.
	require.Equal(t, user1.Username, user2.Username, "Username mismatch")
	require.Equal(t, user1.HashedPassword, user2.HashedPassword, "HashedPassword mismatch")
	require.Equal(t, user1.FullName, user2.FullName, "FullName mismatch")

	// Verify that account ID and CreatedAt are not zero.
	require.NotZero(t, user2.CreatedAt, "CreatedAt is zero")
	require.True(t, user2.PasswordChangedAt.IsZero(), "PasswordChangedAt is not zero")
}
