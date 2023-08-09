package db

import "time"

type Report struct {
	ID          int // PK
	UserID      int // FK
	Description string
	Coordinates string
	Severity    string
	Date        time.Time
	Timestamp   time.Time
}
