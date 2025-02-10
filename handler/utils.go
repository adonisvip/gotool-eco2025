package handler

import (
  "strconv"
  "strings"
)

func calculateTotalSavings(week int) int {
	total := 0
	for i := 0; i < week; i++ {
		total += savingsPlan[i]
	}
	return total
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func formatCurrency(amount int) string {
	amountStr := strconv.Itoa(amount)
	n := len(amountStr)
	
	if n <= 3 {
		return amountStr
	}

	var result []string
	for i := n; i > 0; i -= 3 {
		start := max(0, i-3)
		result = append([]string{amountStr[start:i]}, result...)
	}

	return strings.Join(result, ".")
}
