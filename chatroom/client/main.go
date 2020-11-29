package main

import (
	"fmt"
	"os"
)

var userId int
var userPwd string

func main() {

	// 接收用戶的選擇
	var key int

	// 判斷是否還繼續顯示菜單
	var loop = true

	for loop {
		fmt.Println("----------------歡迎登入多人聊天系統----------------")
		fmt.Println("\t\t\t 1. 登入聊天室")
		fmt.Println("\t\t\t 2. 註冊用戶")
		fmt.Println("\t\t\t 3. 退出系統")
		fmt.Println("\t\t\t 請選擇(1-3):")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登入聊天室")
			loop = false
		case 2:
			fmt.Println("註冊用戶")
			loop = false
		case 3:
			fmt.Println("退出系統")
			os.Exit(0)
		default:
			fmt.Println("您的輸入有誤，請重新輸入")
		}

	}

	// 更新用戶的輸入，顯示新的提示信息
	if key == 1 {
		// 說明用戶要登入
		fmt.Println("請輸入用戶的id")
		fmt.Scanf("%d\n", &userId)
		fmt.Println("請輸入用戶的密碼")
		fmt.Scanf("%v\n", &userPwd)

		// 先把登入的函數，寫到另外一個文件
		login(userId, userPwd)
		// if err != nil {
		// 	fmt.Println("登入失敗")
		// } else {
		// 	fmt.Println("登入成功")
		// }

	} else if key == 2 {
		fmt.Println("進行用戶註冊的邏輯")
	}
}
