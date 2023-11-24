package entities

type Budget struct {
	ID          int    `json:"id"`
	UserID      int    `json:"userId"`
	Category    string `json:"category"`
	Amount      int    `json:"amount"` // cents
	Description string `json:"description"`
}


