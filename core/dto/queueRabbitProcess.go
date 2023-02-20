package dto

import "time"

type QueueProcessDTO struct {
	Message   string
	Result    string
	CreatedAt time.Time
}
