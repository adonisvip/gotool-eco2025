package handler

import (
  "gotool-eco2025/config"
  "os"
  "fmt"
  "strconv"
  "net/url"
  "time"
  "net/http"
  "log"
  "encoding/json"
)

var lastUpdateID int64 = 0

func PollUpdates() {
	telegramBotToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates", telegramBotToken)

	for {
		resp, err := http.Get(fmt.Sprintf("%s?offset=%d&timeout=30", apiURL, lastUpdateID+1))
		if err != nil {
			log.Println("❌ Lỗi gọi API getUpdates:", err)
			time.Sleep(5 * time.Second) 
			continue
		}

		// defer resp.Body.Close()

		var result struct {
			OK      bool             `json:"ok"`
			Updates []config.TelegramUpdate `json:"result"`
		}

		err = json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			log.Println("❌ Lỗi giải mã JSON từ getUpdates:", err)
			time.Sleep(5 * time.Second)
			continue
		}

		for _, update := range result.Updates {
			if update.UpdateID > lastUpdateID {
				lastUpdateID = update.UpdateID
				go handleTelegramCommands(update) 
			}
		}

		time.Sleep(1 * time.Second) // Giảm tải API
	}
}

func handleTelegramCommands(update config.TelegramUpdate) {
	if update.Message == nil {
		return
	}

	chatID := update.Message.Chat.ID
	text := update.Message.Text

	location, _ := time.LoadLocation(os.Getenv("TIMEZONE"))
	now := time.Now().In(location)
	weekNum := getWeekNumber(now)
	totalSaved := calculateTotalSavings(weekNum)

	var responseText string

	switch text {
	case "/week":
		responseText = fmt.Sprintf("📅 Tuần hiện tại: %d\n💰 Số tiền cần tiết kiệm: %s VND", weekNum, formatCurrency(savingsPlan[weekNum-1]))
	case "/total":
		responseText = fmt.Sprintf("💰 Tổng số tiền đã tiết kiệm: %s VND", formatCurrency(totalSaved))
	case "/help":
		responseText = "📌 Danh sách lệnh:\n/week - Xem tuần hiện tại\n/total - Xem tổng tiền đã tiết kiệm\n/help - Danh sách lệnh"
	default:
		responseText = "❌ Lệnh không hợp lệ. Gõ /help để xem danh sách lệnh."
	}

	sendTelegramMessageToChat(chatID, responseText)
}

func sendTelegramMessageToChat(chatID int64, message string) {
	telegramBotToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", telegramBotToken)

	data := url.Values{}
	data.Set("chat_id", strconv.FormatInt(chatID, 10))
	data.Set("text", message)

	http.PostForm(apiURL, data)
}

// func ListenForTelegramUpdates() {
// 	telegramBotToken := os.Getenv("TELEGRAM_BOT_TOKEN")
// 	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates", telegramBotToken)

// 	for {
// 		resp, err := http.Get(apiURL)
// 		if err != nil {
// 			log.Println("Lỗi khi nhận tin nhắn Telegram:", err)
// 			time.Sleep(5 * time.Second)
// 			continue
// 		}

// 		defer resp.Body.Close()

// 		var result config.TelegramResponse
// 		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
// 			log.Println("Lỗi giải mã JSON:", err)
// 			time.Sleep(5 * time.Second)
// 			continue
// 		}

// 		for _, update := range result.Result {
// 			handleTelegramCommands(update)
// 		}

// 		time.Sleep(3 * time.Second) // Tránh spam request
// 	}
// }