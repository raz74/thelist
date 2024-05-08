package notification

import "time"

type Notification struct {
	Id      uint
	Message string
	Date    time.Time
}
