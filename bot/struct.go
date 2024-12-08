package bot

// ApiResponse API响应的根对象结构体
type ApiResponse struct {
	Code    int    `json:"code"`    // 返回状态码，0表示成功
	Msg     string `json:"msg"`     // 错误信息
	Message string `json:"message"` // 错误信息（备用）
	Ttl     int    `json:"ttl"`     // 返回TTL
	Data    Data   `json:"data"`    // 数据主体
}

// Data 数据主体结构体
type Data struct {
	SessionList         []Session        `json:"session_list"`          // 会话列表
	HasMore             int              `json:"has_more"`              // 是否有更多会话
	AntiDisturbCleaning bool             `json:"anti_disturb_cleaning"` // 是否启用“一键防骚扰”
	IsAddressListEmpty  int              `json:"is_address_list_empty"` // 地址列表是否为空
	SystemMsg           map[string]int64 `json:"system_msg"`            // 系统消息会话
	ShowLevel           bool             `json:"show_level"`            // 是否显示用户等级
}

// Session 会话结构体
type Session struct {
	TalkerID    int64       `json:"talker_id"`    // 会话对象ID
	SessionType int         `json:"session_type"` // 会话类型
	UnreadCount int         `json:"unread_count"` // 未读消息数
	SessionTS   int64       `json:"session_ts"`   // 会话时间戳
	LastMsg     LastMsg     `json:"last_msg"`     // 最新消息
	GroupName   string      `json:"group_name"`   // 群组名称
	AccountInfo AccountInfo `json:"account_info"` // 账户信息
}

// LastMsg 最新消息结构体
type LastMsg struct {
	SenderUID int64  `json:"sender_uid"` // 消息发送者UID
	Content   string `json:"content"`    // 消息内容（JSON格式）
	Timestamp int64  `json:"timestamp"`  // 消息发送时间戳
	MsgType   int    `json:"msg_type"`   // 消息类型
	MsgSource int    `json:"msg_source"` // 消息来源
}

// AccountInfo 账户信息结构体
type AccountInfo struct {
	Name   string `json:"name"`    // 账户名称
	PicURL string `json:"pic_url"` // 账户头像URL
}

// NavResponse 用户信息结构体
type NavResponse struct {
	Code int `json:"code"` // API 响应状态码
	Data struct {
		Mid int64 `json:"mid"` // 用户的唯一 ID
	} `json:"data"`
}
