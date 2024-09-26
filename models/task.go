package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	AssignedTo  string             `json:"assigned_to"`
	Status      string             `json:"status"`
}
