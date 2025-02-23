package handler

import (
  "time"
  "fmt"
  "os"
)

var savingsPlan = []int{
	10000, 20000, 30000, 40000, 50000, 60000, 70000, 80000, 90000, 100000,
	110000, 120000, 130000, 140000, 150000, 160000, 170000, 180000, 190000, 200000,
	210000, 220000, 230000, 240000, 250000, 260000, 270000, 280000, 290000, 300000,
	310000, 320000, 330000, 340000, 350000, 360000, 370000, 380000, 390000, 400000,
	410000, 420000, 430000, 440000, 450000, 460000, 470000, 480000, 490000, 500000,
	510000, 520000,
}

func ScheduleNotifications() {
	for {
		now := time.Now()
		location, _ := time.LoadLocation(os.Getenv("TIMEZONE"))
		now = now.In(location)

		weekday := now.Weekday()
		hour := now.Hour()
		minute := now.Minute()
		second := now.Second()

		if (weekday == time.Monday || weekday == time.Sunday) && hour == 21 && minute == 0 && second == 0 {
			weekNum := getWeekNumber(now)
			amount := savingsPlan[weekNum-1]
			message := fmt.Sprintf("📢 Tuần %d: Đến giờ chuyển khoản tiết kiệm!\n💰 Số tiền cần tiết kiệm: %s VND", weekNum, formatCurrency(amount))
			sendTelegramMessage(message)
		}

		time.Sleep(1 * time.Second) 
	}
}