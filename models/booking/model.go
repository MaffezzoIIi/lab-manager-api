package booking

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Period string

const (
	Manha Period = "manha"
	Tarde Period = "tarde"
	Noite Period = "noite"
)

type Booking struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	LabID      primitive.ObjectID `bson:"lab_id"`
	UserID     primitive.ObjectID `bson:"user_id"`
	StartTime  time.Time          `bson:"start_time"`
	EndTime    time.Time          `bson:"end_time"`
	Period     string             `bson:"period,omitempty"`
	DaysOfWeek []time.Weekday     `bson:"daysOfWeek,omitempty"`
}
