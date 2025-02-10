package config

type TelegramResponse struct {
	Result []TelegramUpdate `json:"result"`
}

type TelegramUpdate struct {
	UpdateID int64         `json:"update_id"`
	Message  *TelegramMessage `json:"message"`
}

type TelegramMessage struct {
	MessageID int64          `json:"message_id"`
  From      TelegramUser  `json:"from"`
	Chat      TelegramChat   `json:"chat"`
	Text      string         `json:"text"`
}

type TelegramChat struct {
	ID int64 `json:"id"`  // Sửa từ int -> int64
}

type TelegramUser struct {
	ID int64 `json:"id"`
}