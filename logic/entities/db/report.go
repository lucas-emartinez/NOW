package db

import "time"

type Report struct {
	ID          int // PK
	UserID      int // FK
	Description string
	Coordinates string
	Severity    int
	Date        time.Time
	Timestamp   time.Time
}
