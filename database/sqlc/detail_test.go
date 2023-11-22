package db

import (
	"bookkeeping-backend/database/helper/random"
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomDetail(t *testing.T, user User, category string) Detail {
	arg := CreateDetailParams{
		Username: user.Username,
		Category: category,
		Cost:     random.Cost(),
		Date:     random.Date(),
	}
	detail, err := testQueries.CreateDetail(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, detail)

	require.Equal(t, arg.Username, detail.Username)
	require.Equal(t, arg.Category, detail.Category)
	require.Equal(t, arg.Cost, detail.Cost)
	require.WithinDuration(t, arg.Date, detail.Date, time.Second)
	require.NotZero(t, detail.ID)
	require.NotZero(t, detail.CreatedAt)

	return detail
}

func TestCreateDetail(t *testing.T) {
	createRandomDetail(t, createRandomUser(t), random.Category())
}

func TestGetDetailById(t *testing.T) {
	detail1 := createRandomDetail(t, createRandomUser(t), random.Category())
	detail2, err := testQueries.GetDetail(context.Background(), detail1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, detail2)

	require.Equal(t, detail1.ID, detail2.ID)
	require.Equal(t, detail1.Username, detail2.Username)
	require.Equal(t, detail1.Category, detail2.Category)
	require.Equal(t, detail1.Cost, detail2.Cost)
	require.WithinDuration(t, detail1.Date, detail2.Date, time.Second)
	require.WithinDuration(t, detail1.CreatedAt, detail2.CreatedAt, time.Second)
}

func TestListDetailsByUserId(t *testing.T) {
	user := createRandomUser(t)
	for i := 0; i < 10; i += 1 {
		createRandomDetail(t, user, random.Category())
	}

	arg := ListDetailsByUserParams{
		Username: user.Username,
		Offset:   5,
		Limit:    5,
	}
	details, err := testQueries.ListDetailsByUser(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, details, 5)

	for _, detail := range details {
		require.NotEmpty(t, detail)
	}
}

func TestUpdateDetail(t *testing.T) {
	category := random.Category()
	detail1 := createRandomDetail(t, createRandomUser(t), random.Category())
	arg := UpdateDetailParams{
		ID:       detail1.ID,
		Category: category,
		Cost:     random.Cost(),
		Date:     random.Date(),
	}
	detail2, err := testQueries.UpdateDetail(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, detail2)

	require.Equal(t, detail1.ID, detail2.ID)
	require.Equal(t, detail1.Username, detail2.Username)
	require.Equal(t, arg.Category, detail2.Category)
	require.Equal(t, arg.Cost, detail2.Cost)
	require.WithinDuration(t, arg.Date, detail2.Date, time.Second)
	require.WithinDuration(t, detail1.CreatedAt, detail2.CreatedAt, time.Second)
}

func DeleteDetail(t *testing.T) {
	detail1 := createRandomDetail(t, createRandomUser(t), random.Category())
	err := testQueries.DeleteDetail(context.Background(), detail1.ID)
	require.NoError(t, err)

	detail2, err := testQueries.GetDetail(context.Background(), detail1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, detail2)
}
