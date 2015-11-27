package worker

import (
	"fmt"
	"github.com/strebul/gogy/model"
	"time"
)

func Request(object string, message string, custom string, duration int) model.Request {
	var query string

	query += fmt.Sprintf(`object: "%s"`, object)

	if len(message) > 0 {
		query += fmt.Sprintf(` AND message: "%s"`, message)
	}

	if len(custom) > 0 {
		query += fmt.Sprintf(` AND %s`, custom)
	}
	//	fmt.Println(query)
	return model.Request{
		Query:     query,
		TimeStart: time.Now().Add(-time.Duration(duration) * time.Hour),
		TimeEnd:   time.Now(),
		Size:      1000,
		Order:     "desc",
	}
}
