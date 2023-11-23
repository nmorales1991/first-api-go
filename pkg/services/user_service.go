package services

import (
	"context"
	"fmt"
	"github.com/nmorales1991/first-api-go/pkg/models"
	"github.com/nmorales1991/first-api-go/pkg/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	db *repository.DB
}

func NewUserService(db *repository.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) GetUsers() ([]models.User, error) {
	collection := s.db.GetCollection("go_example", "users")
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	var users []models.User
	if err := cursor.All(context.Background(), &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) CreateUser(user models.User) error {
	collection := s.db.GetCollection("go_example", "users")
	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) UpdateUser(id string, user models.User) error {
	collection := s.db.GetCollection("go_example", "users")
	objectID, err := primitive.ObjectIDFromHex(id)
	result, err := collection.UpdateOne(context.Background(), bson.M{"_id": objectID}, bson.M{"$set": user})
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return fmt.Errorf("no document found with id %s", id)
	}
	return nil

}
