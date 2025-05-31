package lab

import (
	"context"
	"lab-manager-api/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SaveLab(lab Lab) (Lab, error) {
	collection := config.DB.Database("lab-manager").Collection("labs")
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
	return err
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
