package user

import (
	"Intersolusi_Teknologi_Asia/internal/domain"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	db *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{db: db.Collection("users")}
}

func (r *UserRepository) Register(username, password string, email string) error {
	var user domain.User
	err := r.db.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err == nil {
		return errors.New("user already exists")
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	_, err = r.db.InsertOne(context.Background(), domain.User{
		Username: username,
		Password: string(hashedPassword),
		Email:    email,
	})

	return err
}

func (r *UserRepository) Login(username, password string) (string, error) {
	var user domain.User
	err := r.db.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return "", errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	return user.ID, nil
}
