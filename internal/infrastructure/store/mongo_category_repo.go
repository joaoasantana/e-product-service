package store

import (
	"context"
	"github.com/joaoasantana/e-product-service/internal/domain/entity"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"time"
)

type MongoCategoryRepository struct {
	collection *mongo.Collection
}

func NewMongoCategoryRepository(mongoDatabase *mongo.Database) *MongoCategoryRepository {
	return &MongoCategoryRepository{
		collection: mongoDatabase.Collection("categories"),
	}
}

type MongoCategory struct {
	ID        bson.ObjectID `bson:"_id,omitempty"`
	Name      string        `bson:"name"`
	CreatedAt time.Time     `bson:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at"`
}

func (r *MongoCategoryRepository) Create(category *entity.Category) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	model := MongoCategory{
		Name:      category.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := r.collection.InsertOne(ctx, model)
	if err != nil {
		return err
	}

	return nil
}

func (r *MongoCategoryRepository) FindAll() ([]entity.Category, error) {
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

	var entities []entity.Category
	for _, model := range models {
		entities = append(entities, entity.Category{
			ID:   model.ID.Hex(),
			Name: model.Name,
		})
	}

	return entities, nil
}

func (r *MongoCategoryRepository) FindById(id string) (*entity.Category, error) {
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

	category := &entity.Category{
		ID:   model.ID.Hex(),
		Name: model.Name,
	}

	return category, nil
}

func (r *MongoCategoryRepository) Validate(name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result := r.collection.FindOne(ctx, bson.M{"name": name})

	return result.Err()
}
