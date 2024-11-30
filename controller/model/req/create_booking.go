package req

import (
	"time"
)

type CreateBooking struct {
	LabID     string    `json:"lab_id" binding:"required"`
	UserID    string    `json:"user_id" binding:"required"`
	StartTime time.Time `json:"start_time" binding:"required"`
	EndTime   time.Time `json:"end_time" binding:"required"`
	Period    string    `json:"period" binding:"required"`
}
