package message

const (
	LoginMesType       = "LoginMes"
	LoginResMesType    = "LoginResMes"
	RegisterMesType    = "RegisterMes"
	RegisterResMesType = "RegisterResMes"
)

type Message struct {
	Type string `json:"type"` // 消息類型
	Data string `json:"data"` // 消息資料
}

type LoginMes struct {
	UserId   int    `json:"userId"`   // 用戶id
	UserPwd  string `json:"userPwd"`  // 用戶密碼
	UserName string `json:"userName"` // 用戶名
}

type LoginResMes struct {
	Code  int    `json:"code"`  // 返回狀態碼500表示該用戶未註冊、返回200表示登入成功
	Error string `json:"error"` // 返回錯誤信息
}

type RegisterMes struct {
	UserId   int    `json:"userId"`   // 用戶id
	UserPwd  string `json:"userPwd"`  // 用戶密碼
	UserName string `json:"userName"` // 用戶名
}

type RegisterResMes struct {
	Code  int    `json:"code"`  // 返回狀態碼500表示該用戶未註冊、返回200表示登入成功
	Error string `json:"error"` // 返回錯誤信息
}
