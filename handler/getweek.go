package handler

import (
  "time"
)

func getWeekNumber(t time.Time) int {
	// startDate := time.Date(2025, 1, 1, 0, 0, 0, 0, t.Location())
	startDate := time.Date(2024, 12, 30, 0, 0, 0, 0, t.Location())
	diffDays := int(t.Sub(startDate).Hours() / 24)
	weekNum := diffDays/7 + 1
	if weekNum < 1 {
		return 1
	}
	if weekNum > 52 {
		return 52
	}
	return weekNum
}