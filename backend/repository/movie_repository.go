package repository

import (
	"context"
	"github.com/AfomiaTadesse/Afomia_M/backend/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MovieRepository interface {
	Create(ctx context.Context, movie *domain.Movie) error
	GetByID(ctx context.Context, id string) (*domain.Movie, error)
	GetAll(ctx context.Context, page, size int) ([]domain.Movie, int64, error)
	SearchByTitle(ctx context.Context, title string, page, size int) ([]domain.Movie, int64, error)
	Update(ctx context.Context, id string, movie *domain.Movie) error
	Delete(ctx context.Context, id string) error
	GetByUserID(ctx context.Context, userID string, page, size int) ([]domain.Movie, int64, error)
}

type movieRepository struct {
	collection *mongo.Collection
}

func NewMovieRepository(db *mongo.Database) MovieRepository {
	return &movieRepository{
		collection: db.Collection("movies"),
	}
}

func (r *movieRepository) Create(ctx context.Context, movie *domain.Movie) error {
	_, err := r.collection.InsertOne(ctx, movie)
	return err
}

func (r *movieRepository) GetByID(ctx context.Context, id string) (*domain.Movie, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var movie domain.Movie
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&movie)
	if err != nil {
		return nil, err
	}

	return &movie, nil
}

func (r *movieRepository) GetAll(ctx context.Context, page, size int) ([]domain.Movie, int64, error) {
	skip := int64((page - 1) * size)
	opts := options.Find().
		SetSkip(skip).
		SetLimit(int64(size))

	cursor, err := r.collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var movies []domain.Movie
	if err = cursor.All(ctx, &movies); err != nil {
		return nil, 0, err
	}

	total, err := r.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, 0, err
	}

	return movies, total, nil
}

func (r *movieRepository) SearchByTitle(ctx context.Context, title string, page, size int) ([]domain.Movie, int64, error) {
	skip := int64((page - 1) * size)
	opts := options.Find().
		SetSkip(skip).
		SetLimit(int64(size))

	filter := bson.M{
		"title": bson.M{
			"$regex":   title,
			"$options": "i", // case insensitive
		},
	}

	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var movies []domain.Movie
	if err = cursor.All(ctx, &movies); err != nil {
		return nil, 0, err
	}

	total, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	return movies, total, nil
}

func (r *movieRepository) Update(ctx context.Context, id string, movie *domain.Movie) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.collection.UpdateOne(
		ctx,
		bson.M{"_id": objID},
		bson.M{"$set": movie},
	)
	return err
}

func (r *movieRepository) Delete(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}

func (r *movieRepository) GetByUserID(ctx context.Context, userID string, page, size int) ([]domain.Movie, int64, error) {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, 0, err
	}

	skip := int64((page - 1) * size)
	opts := options.Find().
		SetSkip(skip).
		SetLimit(int64(size))

	filter := bson.M{"userId": objID}

	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var movies []domain.Movie
	if err = cursor.All(ctx, &movies); err != nil {
		return nil, 0, err
	}

	total, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	return movies, total, nil
}