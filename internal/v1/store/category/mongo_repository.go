package category

import (
	"context"
	"errors"
	"github.com/joaoasantana/e-product-service/internal/v1/domain/core/category"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"time"
)

const (
	collectionName = "categories"
)

type MongoRepository struct {
	collection *mongo.Collection
}

func NewMongoRepository(db *mongo.Database) *MongoRepository {
	return &MongoRepository{
		collection: db.Collection(collectionName),
	}
}

func (r *MongoRepository) Create(entity *category.Entity) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	model := MongoModel{
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

func (r *MongoRepository) FindAll() ([]category.Entity, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	var models []MongoModel

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

func (r *MongoRepository) FindByID(id string) (*category.Entity, error) {
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

	var model MongoModel

	if err = result.Decode(&model); err != nil {
		return nil, err
	}

	entity := &category.Entity{
		ID:   model.ID.Hex(),
		Name: model.Name,
	}

	return entity, nil
}

func (r *MongoRepository) FindByName(name string) (*category.Entity, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result := r.collection.FindOne(ctx, bson.M{"name": name})
	if result.Err() != nil {
		return nil, result.Err()
	}

	var model MongoModel

	if err := result.Decode(&model); err != nil {
		return nil, err
	}

	entity := &category.Entity{
		ID:   model.ID.Hex(),
		Name: model.Name,
	}

	return entity, nil
}
