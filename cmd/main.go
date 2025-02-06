package main

import (
	"fmt"
	"log"
	"os"
	"time"

  "gotool-eco2025/handler"
)

func main() {

	timezone := os.Getenv("TIMEZONE")
	if timezone == "" {
		timezone = "Asia/Bangkok"
	}
	location, err := time.LoadLocation(timezone)
	if err != nil {
		log.Fatalf("Lỗi tải múi giờ: %v", err)
	}
	time.Local = location

	fmt.Println("Bot tiết kiệm 2025 đang chạy...")
  handler.SendStartupNotification()
	go handler.ScheduleNotifications()

	select {}
}
