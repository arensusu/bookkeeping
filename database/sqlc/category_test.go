package db

import (
	"bookkeeping-backend/database/helper/random"
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomCategory(t *testing.T) Category {
	arg := random.Category()
	category, err := testQueries.CreateCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category)

	require.Equal(t, arg, category.Name)
	require.NotZero(t, category.ID)
	require.NotZero(t, category.CreatedAt)

	return category
}

func TestCreateCategory(t *testing.T) {
	createRandomCategory(t)
}

func TestGetCategoryById(t *testing.T) {
	category1 := createRandomCategory(t)
	category2, err := testQueries.GetCategoryById(context.Background(), category1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, category2)

	require.Equal(t, category1.ID, category2.ID)
	require.Equal(t, category1.Name, category2.Name)
	require.WithinDuration(t, category1.CreatedAt, category2.CreatedAt, time.Second)
}

func TestListCategorys(t *testing.T) {
	for i := 0; i < 10; i += 1 {
		createRandomCategory(t)
	}

	arg := ListCategorysParams{
		Offset: 5,
		Limit:  5,
	}
	categorys, err := testQueries.ListCategorys(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, categorys, 5)

	for _, category := range categorys {
		require.NotEmpty(t, category)
	}
}

func TestDeleteCategory(t *testing.T) {
	category1 := createRandomCategory(t)
	err := testQueries.DeleteCategory(context.Background(), category1.ID)
	require.NoError(t, err)

	category2, err := testQueries.GetCategoryById(context.Background(), category1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, category2)
}
