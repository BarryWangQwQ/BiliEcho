package bot

import (
	"fmt"
	"github.com/CuteReimu/bilibili/v2"
	"github.com/goccy/go-json"
	"io"
	"net/http"
	"net/url"
	"time"
)

// GetNewSessions 获取会话列表
func getNewSessions(beginTs int64, size int) (*ApiResponse, error) {
	params := url.Values{}
	if beginTs > 0 {
		params.Add("begin_ts", fmt.Sprintf("%d", beginTs))
	}
	params.Add("size", fmt.Sprintf("%d", size))
	params.Add("build", "0")
	params.Add("mobi_app", "web")

	req, err := http.NewRequest("GET", baseSessionURL+"?"+params.Encode(), nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}
	req.Header.Set("Cookie", Cookie)

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应体失败: %w", err)
	}

	var apiResponse ApiResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return nil, fmt.Errorf("解析JSON失败: %w", err)
	}
	return &apiResponse, nil
}

// GetBiliBiliMid 获取用户ID
func getBiliBiliMid(cookie string) (int64, error) {
	req, err := http.NewRequest("GET", navURL, nil)
	if err != nil {
		return 0, fmt.Errorf("请求创建失败: %w", err)
	}
	req.Header.Set("Cookie", cookie)

	resp, err := httpClient.Do(req)
	if err != nil {
		return 0, fmt.Errorf("请求失败: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("读取响应体失败: %w", err)
	}

	var navResp NavResponse
	if err := json.Unmarshal(body, &navResp); err != nil {
		return 0, fmt.Errorf("解析JSON失败: %w", err)
	}
	if navResp.Code != 0 {
		return 0, fmt.Errorf("请求失败，错误码: %d", navResp.Code)
	}
	return navResp.Data.Mid, nil
}

// Filter 过滤函数
func Filter[T any](items []T, predicate func(T) bool) []T {
	j := 0
	for i := 0; i < len(items); i++ {
		if predicate(items[i]) {
			items[j] = items[i]
			j++
		}
	}
	return items[:j]
}

func SetCookie(cookie string) {
	Cookie = cookie
	client.SetRawCookies(cookie)
	uid, _ = getBiliBiliMid(Cookie)
}

// SendMessage 发送私信
func SendMessage(receiverId int64, msgType int, content string) (*bilibili.SendPrivateMessageResult, error) {
	message, err := client.SendPrivateMessage(bilibili.SendPrivateMessageParam{
		SenderUid:    int(uid),
		ReceiverId:   int(receiverId),
		ReceiverType: 1,
		MsgType:      msgType,
		Content:      content,
		Timestamp:    int(time.Now().Unix()),
	})

	if err != nil {
		return message, err
	}

	return message, err
}

// OnSessions 封装为函数，接收一个回调函数处理会话
func OnSessions(callback func(Session)) {
	// 初始化参数
	beginTs := time.Now().UnixMicro()

	for {
		// 调用获取新会话列表的函数
		response, err := getNewSessions(beginTs, defaultSize)
		if err != nil {
			fmt.Printf("请求失败: %v\n", err)
			time.Sleep(pollInterval) // 失败时等待后重试
			continue
		}

		// 过滤会话列表，排除特定 UID 的会话
		sessions := Filter(response.Data.SessionList, func(s Session) bool {
			return s.LastMsg.SenderUID != uid // 替换为实际 UID
		})

		// 使用 Goroutine 并发执行回调，不阻塞主循环
		for _, session := range sessions {
			go callback(session) // 为每个会话启动一个 Goroutine 执行回调
		}

		// 如果有新会话，更新 beginTs
		if len(sessions) > 0 {
			beginTs = sessions[0].SessionTS
		}

		// 按固定时间间隔进行轮询
		time.Sleep(pollInterval)
	}
}
