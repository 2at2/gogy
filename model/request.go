package model

import "time"

type Request struct {
	Query     string
	TimeStart time.Time
	TimeEnd   time.Time
	Size      int
	Order     string
}
