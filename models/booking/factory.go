package booking

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
