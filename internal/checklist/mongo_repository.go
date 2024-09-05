package checklist

import (
	"Intersolusi_Teknologi_Asia/internal/domain"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type ChecklistRepository struct {
	db *mongo.Collection
}

func NewChecklistRepository(db *mongo.Database) *ChecklistRepository {
	return &ChecklistRepository{db: db.Collection("checklists")}
}

func (r *ChecklistRepository) CreateChecklist(title string) error {
	_, err := r.db.InsertOne(context.Background(), domain.Checklist{
		Title: title,
		Items: []domain.ChecklistItem{},
	})
	return err
}

func (r *ChecklistRepository) AddItem(checklistID, item string) error {
	id, err := primitive.ObjectIDFromHex(checklistID)
	if err != nil {
		return fmt.Errorf("invalid checklistID: %w", err)
	}

	newItem := domain.ChecklistItem{
		ID:     primitive.NewObjectID(),
		Title:  item,
		Status: "incomplete",
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$push": bson.M{"items": newItem}}
	result, err := r.db.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("error updating document: %w", err)
	}

	if result.ModifiedCount == 0 {
		log.Println("No documents were updated")
		return fmt.Errorf("No Dcuments were updated.")
	} else {
		log.Println("Successfully updated the document")
	}
	return nil
}

func (r *ChecklistRepository) UpdateItem(checklistID, itemID, newTitle string) error {
	checklistObjectID, err := primitive.ObjectIDFromHex(checklistID)
	if err != nil {
		return fmt.Errorf("invalid checklistID: %w", err)
	}
	itemObjectID, err := primitive.ObjectIDFromHex(itemID)
	if err != nil {
		return fmt.Errorf("invalid itemID: %w", err)
	}

	filter := bson.M{"_id": checklistObjectID, "items._id": itemObjectID}
	update := bson.M{"$set": bson.M{"items.$.title": newTitle}}
	result, err := r.db.UpdateOne(context.Background(), filter, update)

	if err != nil {
		return fmt.Errorf("error updating item title: %w", err)
	}

	if result.ModifiedCount == 0 {
		log.Println("No documents were updated")
		return fmt.Errorf("No Dcuments were updated.")
	} else {
		log.Println("Successfully updated the item title")
	}
	return nil
}

func (r *ChecklistRepository) DeleteItem(checklistID, itemID string) error {

	checklistObjectID, err := primitive.ObjectIDFromHex(checklistID)
	if err != nil {
		return fmt.Errorf("invalid checklistID: %w", err)
	}

	itemObjectID, err := primitive.ObjectIDFromHex(itemID)
	if err != nil {
		return fmt.Errorf("invalid itemID: %w", err)
	}

	filter := bson.M{"_id": checklistObjectID}

	update := bson.M{"$pull": bson.M{"items": bson.M{"_id": itemObjectID}}}

	result, err := r.db.UpdateOne(context.Background(), filter, update)

	if err != nil {
		return fmt.Errorf("error deleting item: %w", err)
	}

	if result.ModifiedCount == 0 {
		log.Println("No documents were updated")
		return fmt.Errorf("No Dcuments were updated.")
	} else {
		log.Println("Successfully deleted the item")
	}

	return nil
}

func (r *ChecklistRepository) UpdateItemStatus(checklistID, itemID, newStatus string) error {

	checklistObjectID, err := primitive.ObjectIDFromHex(checklistID)
	if err != nil {
		return fmt.Errorf("invalid checklistID: %w", err)
	}

	itemObjectID, err := primitive.ObjectIDFromHex(itemID)
	if err != nil {
		return fmt.Errorf("invalid itemID: %w", err)
	}

	filter := bson.M{"_id": checklistObjectID, "items._id": itemObjectID}

	update := bson.M{"$set": bson.M{"items.$.status": newStatus}}

	result, err := r.db.UpdateOne(context.Background(), filter, update)

	if err != nil {
		return fmt.Errorf("error updating item status: %w", err)
	}

	if result.ModifiedCount == 0 {
		log.Println("No documents were updated")
		return fmt.Errorf("No Dcuments were updated.")
	} else {
		log.Println("Successfully updated the item status")
	}

	return nil
}

func (r *ChecklistRepository) GetChecklists() ([]domain.Checklist, error) {
	cur, err := r.db.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	var checklists []domain.Checklist
	for cur.Next(context.Background()) {
		var checklist domain.Checklist
		cur.Decode(&checklist)
		checklists = append(checklists, checklist)
	}
	return checklists, nil
}

func (r *ChecklistRepository) DeleteChecklist(checklistID string) error {

	id, err := primitive.ObjectIDFromHex(checklistID)
	if err != nil {
		return fmt.Errorf("invalid checklistID: %w", err)
	}

	filter := bson.M{"_id": id}

	result, err := r.db.DeleteOne(context.Background(), filter)

	if err != nil {
		return fmt.Errorf("error deleting checklist: %w", err)
	}

	if result.DeletedCount == 0 {
		log.Println("No documents were deleted")
		return fmt.Errorf("No Dcuments were deleted.")
	}

	return nil
}

func (r *ChecklistRepository) GetCheckListItem(checklistID string) (domain.Checklist, error) {

	id, err := primitive.ObjectIDFromHex(checklistID)
	if err != nil {
		return domain.Checklist{}, fmt.Errorf("invalid checklistID: %w", err)
	}

	filter := bson.M{"_id": id}

	var checklist domain.Checklist

	err = r.db.FindOne(context.Background(), filter).Decode(&checklist)
	if err != nil {
		return domain.Checklist{}, err
	}

	return checklist, nil
}
