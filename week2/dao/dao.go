package dao

import (
	"assignment/week2/model"
	"database/sql"
	"errors"
	xerrors "github.com/pkg/errors"
)

var (
	ErrNoRows            = sql.ErrNoRows
	db                   *sql.DB
	QueryUserInfoByIdSql = "select * from user where id = ?"
)

//dao层
func QueryUserById(id int) (*model.User, error) {
	var u *model.User
	err := db.QueryRow(QueryUserInfoByIdSql, id).Scan(u)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNoRows
	}
	if err != nil {
		return nil, xerrors.Wrap(err, "QueryUserById failed")
	}
	return u, nil
}
