package dto

import "time"

type QueueProcessDTO struct {
	QueueMessage string
	Message      string
	Result       string
	CreatedAt    time.Time
}
