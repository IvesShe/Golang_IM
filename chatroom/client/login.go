package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"

	"../common/message"
)

// 寫一個函數，完成登錄
func login(userId int, userPwd string) (err error) {
	// 下一個就要開始定協議
	// fmt.Printf(" userId = %d userPwd = %s\n", userId, userPwd)
	// return nil

	// 1. 連接到服務器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err = ", err)
		return
	}

	// 延時關閉
	defer conn.Close()

	// 2. 準備通過conn發送消息給服務
	var mes message.Message
	mes.Type = message.LoginMesType

	// 3. 創建一個LoginMes結構體
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	// 4. 將loginMes序列化
	// 這邊得到的data是[]byte
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		return
	}

	// 5. 把data賦給mes.Data字段
	// 要先將[]byte的data轉成string
	mes.Data = string(data)

	// 6. 將mes進行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		return
	}

	// 7. 到這一步，data就是我們要發送的消息
	// 7.1 先把data的長度發給服務器
	// 先獲取到data的長度 -> 轉成一個表示長度的byte切片
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)

	// 發送長度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(buf) fail", err)
		return
	}

	fmt.Printf("客戶端發送消息的長度 = %d\n", len(data))
	fmt.Printf("內容是:\n %s\n", string(data))

	// 發送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) fail", err)
		return
	}

	// 休眠10秒
	// time.Sleep(10 * time.Second)
	// fmt.Println("休眠了10秒...")

	// 這裡還需要處理服務器端返回的消息
	mes, err = readPkg(conn)
	if err != nil {
		fmt.Println("readPkg(conn) err = ", err)
		return
	}

	//fmt.Println("readPkg(conn): ", mes)

	// 將mes的Data部分反序列化成LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		fmt.Println("登入成功")
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	} else {
		fmt.Println("不明原因失敗")
	}

	//fmt.Println("json.Unmarshal([]byte(mes.Data), &loginResMes): ", loginResMes.Code, loginResMes.Error)

	return
}
