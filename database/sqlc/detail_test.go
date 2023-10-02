package db

import (
	"bookkeeping-backend/database/helper/random"
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomDetail(t *testing.T, user User, category Category) Detail {
	arg := CreateDetailParams{
		UserID:     user.ID,
		CategoryID: category.ID,
		Cost:       random.Cost(),
		Date:       random.Date(),
	}
	detail, err := testQueries.CreateDetail(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, detail)

	require.Equal(t, arg.UserID, detail.UserID)
	require.Equal(t, arg.CategoryID, detail.CategoryID)
	require.Equal(t, arg.Cost, detail.Cost)
	require.WithinDuration(t, arg.Date, detail.Date, time.Second)
	require.NotZero(t, detail.ID)
	require.NotZero(t, detail.CreatedAt)

	return detail
}

func TestCreateDetail(t *testing.T) {
	createRandomDetail(t, createRandomUser(t), createRandomCategory(t))
}

func TestGetDetailById(t *testing.T) {
	detail1 := createRandomDetail(t, createRandomUser(t), createRandomCategory(t))
	detail2, err := testQueries.GetDetailById(context.Background(), detail1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, detail2)

	require.Equal(t, detail1.ID, detail2.ID)
	require.Equal(t, detail1.UserID, detail2.UserID)
	require.Equal(t, detail1.CategoryID, detail2.CategoryID)
	require.Equal(t, detail1.Cost, detail2.Cost)
	require.WithinDuration(t, detail1.Date, detail2.Date, time.Second)
	require.WithinDuration(t, detail1.CreatedAt, detail2.CreatedAt, time.Second)
}

func TestListDetailsByUserId(t *testing.T) {
	user := createRandomUser(t)
	for i := 0; i < 10; i += 1 {
		createRandomDetail(t, user, createRandomCategory(t))
	}

	arg := ListDetailsByUserIdParams{
		UserID: user.ID,
		Offset: 5,
		Limit:  5,
	}
	details, err := testQueries.ListDetailsByUserId(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, details, 5)

	for _, detail := range details {
		require.NotEmpty(t, detail)
	}
}

func TestUpdateDetail(t *testing.T) {
	category := createRandomCategory(t)
	detail1 := createRandomDetail(t, createRandomUser(t), createRandomCategory(t))
	arg := UpdateDetailParams{
		ID:         detail1.ID,
		CategoryID: category.ID,
		Cost:       random.Cost(),
		Date:       random.Date(),
	}
	detail2, err := testQueries.UpdateDetail(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, detail2)

	require.Equal(t, detail1.ID, detail2.ID)
	require.Equal(t, detail1.UserID, detail2.UserID)
	require.Equal(t, arg.CategoryID, detail2.CategoryID)
	require.Equal(t, arg.Cost, detail2.Cost)
	require.WithinDuration(t, arg.Date, detail2.Date, time.Second)
	require.WithinDuration(t, detail1.CreatedAt, detail2.CreatedAt, time.Second)
}

func DeleteDetail(t *testing.T) {
	detail1 := createRandomDetail(t, createRandomUser(t), createRandomCategory(t))
	err := testQueries.DeleteDetail(context.Background(), detail1.ID)
	require.NoError(t, err)

	detail2, err := testQueries.GetDetailById(context.Background(), detail1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, detail2)
}
