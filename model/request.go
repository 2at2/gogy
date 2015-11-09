package model

type Request struct {
	Query     string
	TimeStart int64
	TimeEnd   int64
	Size      int
}
