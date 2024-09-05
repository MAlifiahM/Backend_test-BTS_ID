package domain

type Checklist struct {
	ID    string          `bson:"_id,omitempty"`
	Title string          `bson:"title"`
	Items []ChecklistItem `bson:"items"`
}
