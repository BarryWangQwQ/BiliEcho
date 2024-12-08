# BiliEcho 🎧 - 超轻量 Bilibili 自动私信回复 🚀

<br>

<center>
  <img src="https://cdn-fusion.imgcdn.store/i/2024/11ca4defc7a779df.webp" alt="UniBarrage">
</center>

<br>

**BiliEcho** 是一个基于 Go 的项目，用于与 Bilibili 的消息 API 交互，轻松获取和管理私信会话、发送消息，并实时更新新会话。超级轻量！💻✨

---

## 🌟 功能特点 | Features

- 📬 **获取新会话**：通过 Bilibili API 高效获取和管理会话列表。
- 📩 **发送私信**：向其他用户发送带有自定义内容的私信。
- 🔄 **会话轮询**：定期自动检查新会话。
- 🎛 **会话过滤**：根据自定义条件轻松过滤会话。
- 🔧 **自定义回调**：实时处理会话并执行用户定义的逻辑。

---

## 🚀 快速开始 | Getting Started

### 📋 前置条件 | Prerequisites

1. 安装 [Go](https://go.dev/)（推荐使用 1.19 或更高版本）。
2. 获取你的 Bilibili **Cookie**（用于认证）。

---

### 安装 | Installation

1. 克隆此仓库：
   ```bash  
   git clone https://github.com/yourusername/BiliEcho.git  
   cd BiliEcho  
   ```  

2. 安装依赖：
   ```bash  
   go mod tidy  
   ```  

---

### 使用 | Usage

1. 设置你的 Bilibili **Cookie**：  
   在 `main` 函数中更新 `cookie` 变量：
   ```go  
   cookie := "<your-bilibili-cookie>"  
   ```  

2. 编译并运行项目：
   ```bash  
   go run main.go  
   ```  

3. 在 `OnSessions` 回调中定义你的自定义会话处理逻辑：
   ```go  
   OnSessions(func(s Session) {  
       fmt.Println("Processing session:", s)  
       // 在这里添加逻辑，例如发送消息  
       // SendMessage(s.TalkerID, 1, "Hello!")  
   })  
   ```  

---

## 📚 项目结构 | Project Structure

- `main.go`：程序的入口，包含主要逻辑，如获取会话和发送消息。
- `bilibili/v2`：用于与 Bilibili API 交互的自定义库。
- 工具函数：
    - `GetNewSessions`：从 Bilibili 获取新会话。
    - `SendMessage`：发送私信。
    - `Filter`：根据用户定义的条件过滤会话。
    - `OnSessions`：通过回调轮询并处理会话。

---

## 🛠 配置 | Configuration

通过调整代码中的以下常量以满足你的需求：

- **轮询间隔**：调整程序轮询新会话的频率（`pollInterval`）。
- **请求超时**：设置 API 请求的超时时间（`requestTimeout`）。
- **默认分页大小**：控制每次请求获取的会话数量（`defaultSize`）。

---

## 💡 示例 | Examples

### 示例 1：向所有新会话发送问候消息 | Sending a Greeting Message to All New Sessions

```go  
OnSessions(func (s Session) {
if s.UnreadCount > 0 {
_, err := SendMessage(s.TalkerID, 1, "Hello! 👋")
if err != nil {
fmt.Printf("Failed to send message: %v\n", err)
} else {
fmt.Printf("Greeting sent to: %s\n", s.AccountInfo.Name)
}  
}  
})  
```  

### 示例 2：过滤特定用户的消息 | Filtering Messages from Specific Users

```go  
sessions := Filter(response.Data.SessionList, func(s Session) bool {
return s.LastMsg.SenderUID != 123456789 // 替换为你不想处理的 UID  
})  
```  

---

## ⚙️ 依赖 | Dependencies

- [goccy/go-json](https://github.com/goccy/go-json) - 高性能 JSON 库。
- [CuteReimu/bilibili](https://github.com/CuteReimu/bilibili) - Bilibili API 客户端库。

---

## 🛡️ 安全性 | Security

- **Cookie** 应妥善保管，请勿公开分享。
- 确保你的应用程序安全存储敏感信息。

---

## 🏗️ 未来计划 | Future Plans

- 🌈 支持更多 Bilibili 功能（例如，发送图片与视频链接）。
- 🌟 提供 Web 界面以便更轻松地交互。
- ⚡ 增强会话过滤功能。

---

## 🤝 贡献 | Contributions

欢迎贡献！你可以提交问题或 PR 来帮助改进此项目。

---

## 📄 许可证 | License

此项目采用 MIT 许可证授权。

---

## 🙌 致谢 | Acknowledgments

特别感谢社区提供的 API 文档和支持！🎉

---

❤️ Made with Love by BarryWang.  