package app

// RegistBody 注册结构体
type RegistBody struct {
	ID       string `json:"ID"`
	Name     string `json:"Name"`
	Password string `json:"Password"`
}

// LoginBody 登录结构体
type LoginBody struct {
	ID       string `json:"ID"`
	Password string `json:"Password"`
}

// OnetimeBody 一次性时间时间结构体
type OnetimeBody struct {
	OwnID   string `json:"OwnID"`
	DDL     int64  `json:"DDL"`
	Content string `json:"Content"`
}

// DailyBody 日常事件结构体
type DailyBody struct {
	OwnID   string `json:"OwnID" `
	Begin   int    `json:"Begin"`
	End     int    `json:"End"`
	Content string `json:"Content"`
}

// WeeklyBody 周常事件结构体
type WeeklyBody struct {
	OwnID   string `json:"OwnID" `
	Week    int    `json:"Week"`
	Begin   int    `json:"begin" bson:"bgin"`
	End     int    `json:"end" bson:"end"`
	Content string `json:"content" son:"content"`
}
