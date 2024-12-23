package models

import (
	"context"
	"lab-manager-api/config"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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
	Bookings  []Booking          `bson:"reservas"`
	Softwares []string           `bson:"softwares"`
}

func NewLab(name string, local string, acessible bool, pcNumbers int, status LabStatus, softwares []string) (Lab, error) {
	return Lab{
		Name:      name,
		Local:     local,
		Acessible: acessible,
		PcNumbers: pcNumbers,
		Status:    status,
		Bookings:  []Booking{},
		Softwares: softwares,
	}, nil
}

func SaveLab(lab Lab) (Lab, error) {
	collection := config.DB.Database("lab-manager").Collection("labs")

	// Ensure a new ObjectID is generated
	lab.ID = primitive.NewObjectID()

	result, err := collection.InsertOne(context.Background(), lab)
	if err != nil {
		return Lab{}, err
	}

	lab.ID = result.InsertedID.(primitive.ObjectID)
	if lab.Softwares == nil {
		lab.Softwares = []string{}
	}
	return lab, nil
}

func FindLab(id string) (Lab, error) {
	collection := config.DB.Database("lab-manager").Collection("labs")
	var lab Lab
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Lab{}, err
	}
	err = collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&lab)
	if err != nil {
		return Lab{}, err
	}
	return lab, nil
}

func FindAllLabs() ([]Lab, error) {
	collection := config.DB.Database("lab-manager").Collection("labs")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	var labs []Lab
	for cursor.Next(context.Background()) {
		var lab Lab
		cursor.Decode(&lab)
		labs = append(labs, lab)
	}
	return labs, nil
}

func UpdateLab(lab Lab) (Lab, error) {
	collection := config.DB.Database("lab-manager").Collection("labs")
	_, err := collection.UpdateOne(context.Background(), bson.M{"_id": lab.ID}, bson.M{"$set": lab})
	if err != nil {
		return Lab{}, err
	}
	return lab, nil
}

func DeleteLab(id string) error {
	collection := config.DB.Database("lab-manager").Collection("labs")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(context.Background(), bson.M{"_id": objectID})
	if err != nil {
		return err
	}
	return nil
}

func ReserveLab(labID, userID primitive.ObjectID, startTime, endTime time.Time, period string) (Booking, error) {
	booking, err := NewBooking(labID, userID, startTime, endTime, period, nil)
	if err != nil {
		return Booking{}, err
	}
	return SaveBooking(booking)
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

func SaveSoftwareRequest(request SoftwareRequest) (SoftwareRequest, error) {
	collection := config.DB.Database("lab-manager").Collection("software_requests")
	result, err := collection.InsertOne(context.Background(), request)
	if err != nil {
		return SoftwareRequest{}, err
	}
	request.ID = result.InsertedID.(primitive.ObjectID)
	return request, nil
}
