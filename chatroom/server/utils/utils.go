package utils

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"

	"../../common/message"
)

// 這裡將這些方法關聯到結構體中
type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte //傳輸入使用的緩衝
}

func (t *Transfer) ReadPkg() (mes message.Message, err error) {
	//buf := make([]byte, 8096)
	fmt.Println("讀取客戶端發送的數據...")

	// 讀取消息長度
	// conn.Read 在conn沒有被關閉的情況下，才會阻塞
	// 如果客戶端關閉了conn則就不會阻塞
	_, err = t.Conn.Read(t.Buf[:4])
	if err != nil {
		//err = errors.New("read pkg header error")
		return
	}

	// 根據讀取消息內容

	// 根據buf[:4]轉成一個uint32類型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(t.Buf[0:4])
	n, err := t.Conn.Read(t.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		//err = errors.New("read pkg body error")
		return
	}

	// 把pkgLen反序列化成 -> message.Message
	err = json.Unmarshal(t.Buf[:pkgLen], &mes)
	if err != nil {
		//fmt.Println("json.Unmarsha err = ", err)
		err = errors.New("json.Unmarsha err")
		return
	}
	return
}

func (t *Transfer) WritePkg(data []byte) (err error) {
	// 先發送一個長度給對方
	// 先獲取到data的長度 -> 轉成一個表示長度的byte切片
	var pkgLen uint32
	pkgLen = uint32(len(data))
	//var buf [4]byte
	binary.BigEndian.PutUint32(t.Buf[0:4], pkgLen)

	// 發送長度
	n, err := t.Conn.Write(t.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(buf) fail", err)
		return
	}

	// 發送data本身
	n, err = t.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(buf) fail", err)
		return
	}
	return
}
