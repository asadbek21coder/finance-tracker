package entities

import "time"

type Expense struct {
	ID          int       `json:"id"`
	UserID      int       `json:"userId"`
	Category    string    `json:"category"`
	Amount      int       `json:"amount"` // in cents
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}
