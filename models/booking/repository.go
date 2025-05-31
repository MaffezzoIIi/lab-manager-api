package booking

import (
	"context"
	"lab-manager-api/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
