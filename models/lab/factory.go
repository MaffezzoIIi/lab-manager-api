package lab

import (
	"lab-manager-api/models/booking"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewLab(name string, local string, acessible bool, pcNumbers int, status LabStatus, softwares []string, description string) (Lab, error) {
	return Lab{
		Name:        name,
		Local:       local,
		Acessible:   acessible,
		PcNumbers:   pcNumbers,
		Status:      status,
		Bookings:    []booking.Booking{},
		Softwares:   softwares,
		Description: description,
	}, nil
}

func NewSoftwareRequest(labID, userID primitive.ObjectID, software string) (SoftwareRequest, error) {
	return SoftwareRequest{
		LabID:       labID,
		UserID:      userID,
		Software:    software,
		RequestedAt: time.Now(),
		Status:      "pending",
	}, nil
}
