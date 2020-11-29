package process

import (
	"encoding/json"
	"fmt"
	"net"

	"../../common/message"
	"../utils"
)

type UserProcess struct {
	Conn net.Conn
}

// 編寫一個ServerProcessLogin函數
// 功能: 專門處理登入請求
func (u *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	// 1. 先從mes中取出mes.Data，並直接反序列化成LoginMes
	var loginMes message.LoginMes
	json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err = ", err)
		return
	}

	// 1. 先聲名一個resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	// 2. 再聲名一個LoginResMes
	var loginResMes message.LoginResMes

	// 如果用戶id = 100, 密碼 = 123456，認為合法，否則不合法
	// fmt.Println("loginMes.UserId", loginMes.UserId)
	// fmt.Println("loginMes.UserPwd", loginMes.UserPwd)
	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		// 合法
		loginResMes.Code = 200
	} else {
		// 不合法
		loginResMes.Code = 500
		loginResMes.Error = "該用戶不存在，請註冊再使用..."
	}
	// fmt.Println("loginResMes.Code", loginResMes.Code)
	// fmt.Println("loginResMes.Error", loginResMes.Error)

	// 3. 將loginResMes序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}

	// 4. 將data賦值給resMes
	resMes.Data = string(data)

	// 5. 對resMes進行序列化，准備發送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}

	// 6. 發送data，我們將其封裝到writePkg函數
	// 因為使用分層模式(MVC), 先創建一個Transfer實例，然後讀取
	tf := &utils.Transfer{
		Conn: u.Conn,
	}
	err = tf.WritePkg(data)
	//fmt.Println("err = writePkg(conn, data)")

	return
}
