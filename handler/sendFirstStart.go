package handler

import (
  "time"
  "fmt"
  "os"
)

func SendStartupNotification() {
	now := time.Now()
	location, _ := time.LoadLocation(os.Getenv("TIMEZONE"))
	now = now.In(location)

	weekNum := getWeekNumber(now)
	amount := savingsPlan[weekNum-1]

	message := fmt.Sprintf("🚀 Bot tiết kiệm đã khởi động!\n📅 Tuần hiện tại: %d\n💰 Số tiền cần tiết kiệm: %s VND", weekNum, formatCurrency(amount))
	sendTelegramMessage(message)
}