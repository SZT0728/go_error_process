package main

import (
	"fmt"
	"go_error_process/dao"
	"go_error_process/service"
)

/*
获取用户id=1的基本信息：
包括各科成绩，姓名,以及最后等级（优：三科成绩满90、良：三颗成绩有一科不满90）

*/

/*
先说一下自我总结：
除了错误增加了堆栈信息后比较容易定位
但是dao层、service层以及最外应用层(main)中的if err != nil 一行都没少写
请助教帮忙看看这样的代码有什么问题？
*/

func main() {
	err := dao.Init("root:123@(localhost)/test1")
	if err != nil{
		fmt.Printf("%+v",err)
		return
	}

	err = service.AddUser("张三",18,91,92,93)
	if err != nil{
		fmt.Printf("%+v\n",err)
		return
	}

	userInfo,err := service.GetUserInfo(1)
	if err != nil{
		fmt.Printf("%+v\n",err)
		return
	}

	userInfo.ToString()

}
