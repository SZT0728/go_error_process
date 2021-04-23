package service

import (
	"fmt"
	"go_error_process/dao"
)

func AddUser(userName string, age int32, math, chinese, english float64) error {
	//忽略透传
	userId, err := dao.AddUserAndScore(userName, age, math, chinese, english)
	if err != nil {
		return err
	}
	// userId do something
	fmt.Printf("add user success userId=%d\n", userId)
	return nil
}
