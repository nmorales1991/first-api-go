package services

import (
	"context"
	"first-api-go/pkg/models"
	"first-api-go/pkg/repository"
	"go.mongodb.org/mongo-driver/bson"
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
	_, err := collection.UpdateOne(context.Background(), bson.M{"id": id}, bson.M{"$set": user})
	if err != nil {
		return err
	}
	return nil

}
