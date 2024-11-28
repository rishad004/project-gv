package domain

type Stream struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Chat struct {
	User   string `json:"user"`
	Type   string `json:"type"`
	Text   string `json:"text"`
	Amount string `json:"amount"`
}
