package main

import (
	"fmt"
	"io"
	"net"
)

// func readPkg(conn net.Conn) (mes message.Message, err error) {
// 	buf := make([]byte, 8096)
// 	fmt.Println("讀取客戶端發送的數據...")

// 	// 讀取消息長度
// 	// conn.Read 在conn沒有被關閉的情況下，才會阻塞
// 	// 如果客戶端關閉了conn則就不會阻塞
// 	_, err = conn.Read(buf[:4])
// 	if err != nil {
// 		//err = errors.New("read pkg header error")
// 		return
// 	}

// 	// 根據讀取消息內容

// 	// 根據buf[:4]轉成一個uint32類型
// 	var pkgLen uint32
// 	pkgLen = binary.BigEndian.Uint32(buf[0:4])
// 	n, err := conn.Read(buf[:pkgLen])
// 	if n != int(pkgLen) || err != nil {
// 		//err = errors.New("read pkg body error")
// 		return
// 	}

// 	// 把pkgLen反序列化成 -> message.Message
// 	err = json.Unmarshal(buf[:pkgLen], &mes)
// 	if err != nil {
// 		//fmt.Println("json.Unmarsha err = ", err)
// 		err = errors.New("json.Unmarsha err")
// 		return
// 	}
// 	return
// }

// func writePkg(conn net.Conn, data []byte) (err error) {
// 	// 先發送一個長度給對方
// 	// 先獲取到data的長度 -> 轉成一個表示長度的byte切片
// 	var pkgLen uint32
// 	pkgLen = uint32(len(data))
// 	var buf [4]byte
// 	binary.BigEndian.PutUint32(buf[0:4], pkgLen)

// 	// 發送長度
// 	n, err := conn.Write(buf[:4])
// 	if n != 4 || err != nil {
// 		fmt.Println("conn.Write(buf) fail", err)
// 		return
// 	}

// 	// 發送data本身
// 	n, err = conn.Write(data)
// 	if n != int(pkgLen) || err != nil {
// 		fmt.Println("conn.Write(buf) fail", err)
// 		return
// 	}
// 	return
// }

// // 編寫一個ServerProcessLogin函數
// // 功能: 專門處理登入請求
// func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
// 	// 1. 先從mes中取出mes.Data，並直接反序列化成LoginMes
// 	var loginMes message.LoginMes
// 	json.Unmarshal([]byte(mes.Data), &loginMes)
// 	if err != nil {
// 		fmt.Println("json.Unmarshal fail err = ", err)
// 		return
// 	}

// 	// 1. 先聲名一個resMes
// 	var resMes message.Message
// 	resMes.Type = message.LoginResMesType

// 	// 2. 再聲名一個LoginResMes
// 	var loginResMes message.LoginResMes

// 	// 如果用戶id = 100, 密碼 = 123456，認為合法，否則不合法
// 	// fmt.Println("loginMes.UserId", loginMes.UserId)
// 	// fmt.Println("loginMes.UserPwd", loginMes.UserPwd)
// 	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
// 		// 合法
// 		loginResMes.Code = 200
// 	} else {
// 		// 不合法
// 		loginResMes.Code = 500
// 		loginResMes.Error = "該用戶不存在，請註冊再使用..."
// 	}
// 	// fmt.Println("loginResMes.Code", loginResMes.Code)
// 	// fmt.Println("loginResMes.Error", loginResMes.Error)

// 	// 3. 將loginResMes序列化
// 	data, err := json.Marshal(loginResMes)
// 	if err != nil {
// 		fmt.Println("json.Marshal fail", err)
// 		return
// 	}

// 	// 4. 將data賦值給resMes
// 	resMes.Data = string(data)

// 	// 5. 對resMes進行序列化，准備發送
// 	data, err = json.Marshal(resMes)
// 	if err != nil {
// 		fmt.Println("json.Marshal fail", err)
// 		return
// 	}

// 	// 6. 發送data，我們將其封裝到writePkg函數
// 	err = utils.WritePkg(conn, data)
// 	//fmt.Println("err = writePkg(conn, data)")

// 	return
// }

// 編寫一個ServerProcessMes函數
// 功能: 根據客戶端發送消息種類不同，決定調用哪個函數來處理
// func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
// 	switch mes.Type {
// 	case message.LoginMesType:
// 		// 處理登入
// 		err = process.ServerProcessLogin(conn, mes)
// 		//fmt.Println("@@@serverProcessMes serverProcessLogin")
// 	case message.RegisterMesType:
// 		// 處理註冊
// 	default:
// 		fmt.Println("消息類型不存在，無法處理...")
// 	}
// 	return
// }

// 處理和客戶端的通訊
func processHandle(conn net.Conn) {
	// 這裡需要延時關閉conn
	defer conn.Close()

	// 這裡調用總控，創建一個
	processor := &Processor{
		Conn: conn,
	}
	err := processor.processHandleDetil()
	if err != nil {
		if err != io.EOF {
			fmt.Println("客戶端和服務器通訊協程錯誤 err = ", err)
		}
		return
	}
}

func main() {

	// 提示信息
	fmt.Println("服務器[新的結構]在8889端口監聽......")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("net.Listen err = ", err)
		return
	}

	// 一旦監聽成功，就等待客戶端來連接服務器
	for {
		fmt.Println("等待客戶端來連接服務器......")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err = ", err)
			return
		}

		// 一旦連接成功，則啟動一個協程和客戶端保持通訊
		go processHandle(conn)
	}

}
