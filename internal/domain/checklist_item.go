package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type ChecklistItem struct {
	ID     primitive.ObjectID `bson:"_id"`
	Title  string             `bson:"title"`
	Status string             `bson:"status"`
}
