package store

import (
	"context"
	"errors"
	"github.com/joaoasantana/e-product-service/internal/v1/domain/core/product"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"time"
)

type MongoProductRepository struct {
	collection *mongo.Collection
}

func NewMongoProductRepository(db *mongo.Database) *MongoProductRepository {
	return &MongoProductRepository{
		collection: db.Collection("products"),
	}
}

func (r *MongoProductRepository) Create(entity *product.Entity) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	parsedID, err := bson.ObjectIDFromHex(entity.CategoryID)
	if err != nil {
		return "", err
	}

	model := MongoProduct{
		Name:        entity.Name,
		Description: entity.Description,
		CategoryID:  parsedID,
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

func (r *MongoProductRepository) FindAll() ([]product.Entity, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	results, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var models []MongoProduct

	if err = results.All(ctx, &models); err != nil {
		return nil, err
	}

	var entities []product.Entity

	for _, model := range models {
		entities = append(entities, product.Entity{
			ID:          model.ID.Hex(),
			Name:        model.Name,
			Description: model.Description,
			CategoryID:  model.CategoryID.Hex(),
		})
	}

	return entities, nil
}

func (r *MongoProductRepository) FindByID(id string) (*product.Entity, error) {
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

	var model MongoProduct

	if err = result.Decode(&model); err != nil {
		return nil, err
	}

	entity := &product.Entity{
		ID:          model.ID.Hex(),
		Name:        model.Name,
		Description: model.Description,
		CategoryID:  model.CategoryID.Hex(),
	}

	return entity, nil
}

func (r *MongoProductRepository) FindByName(name string) (*product.Entity, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result := r.collection.FindOne(ctx, bson.M{"name": name})
	if result.Err() != nil {
		return nil, result.Err()
	}

	var model MongoProduct

	if err := result.Decode(&model); err != nil {
		return nil, err
	}

	entity := &product.Entity{
		ID:          model.ID.Hex(),
		Name:        model.Name,
		Description: model.Description,
		CategoryID:  model.CategoryID.Hex(),
	}

	return entity, nil
}
