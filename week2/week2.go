package week2

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	xerrors "github.com/pkg/errors"
)

type User struct{}

var (
	db                   *sql.DB
	QueryUserInfoByIdSql = "select * from user where id = ?"
)

//dao层
func QueryUserById(id int) (*User, error) {
	var u *User
	err := db.QueryRow(QueryUserInfoByIdSql, id).Scan(u)
	if errors.Is(err, sql.ErrNoRows) {
		//个人认为不需要将sql.ErrNoRows抛回上层,只要在上层能够判断是否查询到了数据即可
		//也可以再添加一个布尔类型的返回值用来判断
		return nil, nil
	}
	if err != nil {
		return nil, xerrors.Wrap(err, "QueryUserById failed")
	}
	return u, nil
}

//controller层或service层
func QueryUserInfo(id int) {
	u, err := QueryUserById(id)
	if err != nil {
		//do something to handle error
	}
	if u != nil {
		//do something if user has existed
	}
	//do something if user has not existed
}
