package internal

type Item struct {
	ID          int    `json:"id" bson:"id,omitempty"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	IsComplete  bool   `json:"is_complete" bson:"isComplete"`
}
