package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

// User represents a user in the system
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username string             `bson:"username" json:"username"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"-"`
}

// Movie represents a movie in the collection
type Movie struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Poster      string             `bson:"poster" json:"poster"`
	Trailer     string             `bson:"trailer" json:"trailer"`
	Actors      []string           `bson:"actors" json:"actors"`
	Genres      []string           `bson:"genres" json:"genres"`
	UserID      primitive.ObjectID `bson:"userId" json:"userId"`
}