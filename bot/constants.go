package bot

import (
	"github.com/CuteReimu/bilibili/v2"
	"net/http"
	"time"
)

// 常量定义
const (
	// navURL 是获取用户个人信息（如 mid 等）的 BiliBili API 地址
	navURL = "https://api.bilibili.com/x/web-interface/nav"

	// baseSessionURL 是获取用户会话列表的 BiliBili API 地址
	baseSessionURL = "https://api.vc.bilibili.com/session_svr/v1/session_svr/new_sessions"

	// defaultSize 定义每次请求获取的会话数量上限，用于分页控制
	defaultSize = 100

	// pollInterval 定义会话轮询的时间间隔，确保定期检查新会话
	pollInterval = 3 * time.Second

	// requestTimeout 定义 HTTP 请求的超时时间，防止请求卡住导致程序阻塞
	requestTimeout = 10 * time.Second
)

// 全局变量
var (
	// uid 初始化
	uid = int64(0)

	// Cookie 初始化
	Cookie = ""

	// B站 客户端 初始化
	client = bilibili.New()

	// HTTP 客户端 初始化
	httpClient = &http.Client{Timeout: requestTimeout}
)
