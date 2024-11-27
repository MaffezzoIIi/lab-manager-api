package models

import "time"

type Period string

const (
	Morning Period = "morning"
	Evening Period = "evening"
	Night   Period = "night"
)

type Booking struct {
	ID     string    `bson:"_id,omitempty"`
	LabID  string    `bson:"lab_id"`
	UserID string    `bson:"usuario_id"`
	Date   time.Time `bson:"data"`
	Period Period    `bson:"periodo,omitempty"`
	Start  time.Time `bson:"inicio,omitempty"`
	End    time.Time `bson:"fim,omitempty"`
}
