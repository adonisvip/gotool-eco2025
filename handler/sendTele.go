package handler

import (
  "os"
  "fmt"
	"net/http"
	"net/url"
  "log"
  "time"
)

func sendTelegramMessage(message string) {
	telegramBotToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	telegramChatID := os.Getenv("TELEGRAM_CHAT_ID")

	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", telegramBotToken)
	data := url.Values{}
	data.Set("chat_id", telegramChatID)
	data.Set("text", message)

	maxRetries := 3
	for i := 1; i <= maxRetries; i++ {
		resp, err := http.PostForm(apiURL, data) 
		if err != nil {
			log.Printf("Lỗi gửi tin nhắn Telegram (lần %d): %v\n", i, err)
			time.Sleep(5 * time.Second)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			log.Println("Gửi tin nhắn Telegram thành công:", message)
			return
		} else {
			log.Printf("Lỗi %d khi gửi tin nhắn Telegram (lần %d)\n", resp.StatusCode, i)
		}
	}
	log.Println("Không thể gửi tin nhắn sau 3 lần thử.")
}