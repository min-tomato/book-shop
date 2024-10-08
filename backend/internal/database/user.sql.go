// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user.sql

package database

import (
	"context"
)

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT USR_ID, USR_EMAIL FROM ` + "`" + `USERS` + "`" + ` WHERE USR_EMAIL = ? LIMIT 1
`

type GetUserByEmailRow struct {
	UsrID    uint32
	UsrEmail string
}

func (q *Queries) GetUserByEmail(ctx context.Context, usrEmail string) (GetUserByEmailRow, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, usrEmail)
	var i GetUserByEmailRow
	err := row.Scan(&i.UsrID, &i.UsrEmail)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT USR_ID, USR_EMAIL FROM ` + "`" + `USERS` + "`" + ` WHERE USR_ID = ? LIMIT 1
`

type GetUserByIDRow struct {
	UsrID    uint32
	UsrEmail string
}

func (q *Queries) GetUserByID(ctx context.Context, usrID uint32) (GetUserByIDRow, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, usrID)
	var i GetUserByIDRow
	err := row.Scan(&i.UsrID, &i.UsrEmail)
	return i, err
}
