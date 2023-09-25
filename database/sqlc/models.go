// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0

package db

import (
	"database/sql"
	"time"
)

type Category struct {
	ID        int64        `json:"id"`
	Name      string       `json:"name"`
	CreatedAt sql.NullTime `json:"created_at"`
}

type Detail struct {
	ID         int64         `json:"id"`
	UserID     sql.NullInt64 `json:"user_id"`
	CategoryID sql.NullInt64 `json:"category_id"`
	Cost       int64         `json:"cost"`
	Date       time.Time     `json:"date"`
	CreatedAt  sql.NullTime  `json:"created_at"`
}

type User struct {
	ID        int64        `json:"id"`
	Username  string       `json:"username"`
	Password  string       `json:"password"`
	IsAdmin   sql.NullBool `json:"is_admin"`
	CreatedAt sql.NullTime `json:"created_at"`
}
