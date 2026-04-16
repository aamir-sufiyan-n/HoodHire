package responses

import "time"

type HirerStats struct {
	Total         int64 `json:"total"`
	Subscribed    int64 `json:"subscribed"`
	NotSubscribed int64 `json:"not_subscribed"`
}

type MonthlyRevenue struct {
	Month  time.Time
	Amount int64
}