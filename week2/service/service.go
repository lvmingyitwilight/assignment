package service

import (
	"assignment/week2/dao"
	xerrors "github.com/pkg/errors"
)

//controller层或service层
func QueryUserInfo(id int) {
	u, err := dao.QueryUserById(id)
	if err != nil {
		if xerrors.Is(err, dao.ErrNoRows) {
			//user not exist
		}
		//do something to handle other error
	}
	_ = u
	//do something if user has existed
}
