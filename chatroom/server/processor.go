package main

import (
	"fmt"
	"io"
	"net"

	"../common/message"
	"./process"
	"./utils"
)

// 創建一個Processor的結構體
type Processor struct {
	Conn net.Conn
}

// 編寫一個ServerProcessMes函數
// 功能: 根據客戶端發送消息種類不同，決定調用哪個函數來處理
func (p *Processor) serverProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		// 處理登入
		// 創建一個UserProcess實例
		up := &process.UserProcess{
			Conn: p.Conn,
		}
		err = up.ServerProcessLogin(mes)
		//fmt.Println("@@@serverProcessMes serverProcessLogin")
	case message.RegisterMesType:
		// 處理註冊
	default:
		fmt.Println("消息類型不存在，無法處理...")
	}
	return
}

func (p *Processor) processHandleDetil() (err error) {
	// 循環讀取客戶端發送的信息
	for {
		// 這裡我們將讀取數據包，直接封裝成一個函數readPkg(),返回Message,Err
		// 創建一個Transfer實例完成讀包任務
		tf := &utils.Transfer{
			Conn: p.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客戶端退出，服務器端也退出...")
			} else {
				fmt.Println("readPkg err = ", err)
			}
			return err
		}

		fmt.Println("mes = ", mes)

		err = p.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}
