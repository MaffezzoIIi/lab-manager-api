package lab

import (
	"time"

	"lab-manager-api/models/booking"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SoftwareRequest struct {
	ID          primitive.ObjectID `bson:"_id"`
	LabID       primitive.ObjectID `bson:"labID"`
	UserID      primitive.ObjectID `bson:"userID"`
	Software    string             `bson:"software"`
	RequestedAt time.Time          `bson:"requestedAt"`
	Status      string             `bson:"status"`
}

type LabStatus string

const (
	Disponivel   LabStatus = "disponivel"
	Ocupado      LabStatus = "ocupado"
	EmManutencao LabStatus = "em manutenção"
	Bloqueado    LabStatus = "bloqueado"
)

type Lab struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	Local     string             `bson:"local"`
	Acessible bool               `bson:"acessible"`
	PcNumbers int                `bson:"pcNumbers"`
	Status    LabStatus          `bson:"status"`
	Bookings  []booking.Booking  `bson:"reservas"`
	Softwares []string           `bson:"softwares"`
}
