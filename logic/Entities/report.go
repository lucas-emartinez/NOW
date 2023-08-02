package Entities

import "time"

type Report struct {
	ID       int
	UserID   int
	StreetID int
	Date     time.Time
}
