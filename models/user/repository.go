package user

import (
	"context"
	"lab-manager-api/config"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewUser(name string, userType UserType, password string) (User, error) {
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return User{}, err
	}

	collection := config.DB.Database("lab-manager").Collection("users")

	result, err := collection.InsertOne(context.Background(), User{Name: name, UserType: userType, Password: hashedPassword})
	if err != nil {
		return User{}, err
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()

	return User{ID: id, Name: name, UserType: userType, Password: hashedPassword}, nil
}
