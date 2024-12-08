package main

import (
	"BiliEcho/bot"
	"fmt"
)

func main() {
	cookie := ""
	bot.SetCookie(cookie)

	// BOT 会话事件
	bot.OnSessions(func(s bot.Session) {
		// 在这里处理会话逻辑
		// bot.SendMessage()
		fmt.Println("处理会话:", s)
	})

}
