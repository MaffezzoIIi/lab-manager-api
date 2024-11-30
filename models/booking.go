package models

import (
	"context"
	"lab-manager-api/config"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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
	DaysOfWeek []time.Weekday     `bson:"days_of_week,omitempty"`
}

func NewBooking(labID, userID primitive.ObjectID, startTime, endTime time.Time, period string, daysOfWeek []time.Weekday) (Booking, error) {
	return Booking{
		LabID:      labID,
		UserID:     userID,
		StartTime:  startTime,
		EndTime:    endTime,
		Period:     period,
		DaysOfWeek: daysOfWeek,
	}, nil
}

func SaveBooking(booking Booking) (Booking, error) {
	collection := config.DB.Database("lab-manager").Collection("bookings")
	result, err := collection.InsertOne(context.Background(), booking)
	if err != nil {
		return Booking{}, err
	}
	booking.ID = result.InsertedID.(primitive.ObjectID)
	return booking, nil
}

func DeleteBooking(bookingID primitive.ObjectID) error {
	collection := config.DB.Database("lab-manager").Collection("bookings")
	_, err := collection.DeleteOne(context.Background(), bson.M{"_id": bookingID})
	return err
}

func CreateRecurringBookings(labID, userID primitive.ObjectID, startTime, endTime time.Time, period string, daysOfWeek []time.Weekday, count int) ([]Booking, error) {
	var bookings []Booking
	currentStartTime := startTime
	currentEndTime := endTime

	for i := 0; i < count; i++ {
		for _, day := range daysOfWeek {
			// Ajustar currentStartTime e currentEndTime para o próximo dia da semana especificado
			for currentStartTime.Weekday() != day {
				currentStartTime = currentStartTime.AddDate(0, 0, 1)
				currentEndTime = currentEndTime.AddDate(0, 0, 1)
			}

			booking, err := NewBooking(labID, userID, currentStartTime, currentEndTime, period, daysOfWeek)
			if err != nil {
				return nil, err
			}

			savedBooking, err := SaveBooking(booking)
			if err != nil {
				return nil, err
			}

			bookings = append(bookings, savedBooking)
		}

		// Atualizar currentStartTime e currentEndTime para a próxima semana
		currentStartTime = currentStartTime.AddDate(0, 0, 7)
		currentEndTime = currentEndTime.AddDate(0, 0, 7)
	}

	return bookings, nil
}
