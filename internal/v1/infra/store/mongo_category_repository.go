package store

import (
	"context"
	"errors"
	"github.com/joaoasantana/e-product-service/internal/v1/domain/core/category"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"time"
)

type MongoCategoryRepository struct {
	collection *mongo.Collection
}

func NewMongoCategoryRepository(db *mongo.Database) *MongoCategoryRepository {
	return &MongoCategoryRepository{
		collection: db.Collection("categories"),
	}
}

func (r *MongoCategoryRepository) Create(entity *category.Entity) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	model := MongoCategory{
		Name:      entity.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result, err := r.collection.InsertOne(ctx, model)
	if err != nil {
		return "", err
	}

	objectID, ok := result.InsertedID.(bson.ObjectID)
	if !ok {
		return "", errors.ErrUnsupported
	}

	return objectID.Hex(), nil
}

func (r *MongoCategoryRepository) FindAll() ([]category.Entity, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	var models []MongoCategory

	if err = cursor.All(ctx, &models); err != nil {
		return nil, err
	}

	var entities []category.Entity

	for _, model := range models {
		entities = append(entities, category.Entity{
			ID:   model.ID.Hex(),
			Name: model.Name,
		})
	}

	return entities, nil
}

func (r *MongoCategoryRepository) FindByID(id string) (*category.Entity, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	result := r.collection.FindOne(ctx, bson.M{"_id": objectID})
	if result.Err() != nil {
		return nil, result.Err()
	}

	var model MongoCategory

	if err = result.Decode(&model); err != nil {
		return nil, err
	}

	entity := &category.Entity{
		ID:   model.ID.Hex(),
		Name: model.Name,
	}

	return entity, nil
}

func (r *MongoCategoryRepository) FindByName(name string) (*category.Entity, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result := r.collection.FindOne(ctx, bson.M{"name": name})
	if result.Err() != nil {
		return nil, result.Err()
	}

	var model MongoCategory

	if err := result.Decode(&model); err != nil {
		return nil, err
	}

	entity := &category.Entity{
		ID:   model.ID.Hex(),
		Name: model.Name,
	}

	return entity, nil
}
