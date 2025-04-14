package store

import (
	"context"
	"github.com/joaoasantana/e-product-service/internal/domain/entity"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"time"
)

type MongoProductRepository struct {
	collection *mongo.Collection
}

func NewMongoProductRepository(database *mongo.Database) *MongoProductRepository {
	return &MongoProductRepository{
		collection: database.Collection("products"),
	}
}

type MongoProduct struct {
	ID          bson.ObjectID `bson:"_id,omitempty"`
	Name        string        `bson:"name"`
	Description string        `bson:"description,omitempty"`
	CategoryID  bson.ObjectID `bson:"category_id"`
	CreatedAt   time.Time     `bson:"created_at"`
	UpdatedAt   time.Time     `bson:"updated_at"`
}

func (r *MongoProductRepository) Create(product *entity.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	categoryID, err := bson.ObjectIDFromHex(product.CategoryID)
	if err != nil {
		return err
	}

	model := MongoProduct{
		Name:        product.Name,
		Description: product.Description,
		CategoryID:  categoryID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	_, err = r.collection.InsertOne(ctx, model)
	if err != nil {
		return err
	}

	return nil
}

func (r *MongoProductRepository) FindAll() ([]entity.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	var models []MongoProduct
	if err = cursor.All(ctx, &models); err != nil {
		return nil, err
	}

	var entities []entity.Product
	for _, model := range models {
		entities = append(entities, entity.Product{
			ID:          model.ID.Hex(),
			Name:        model.Name,
			Description: model.Description,
			CategoryID:  model.CategoryID.Hex(),
		})
	}

	return entities, nil
}

func (r *MongoProductRepository) FindById(id string) (*entity.Product, error) {
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

	category := &entity.Product{
		ID:          model.ID.Hex(),
		Name:        model.Name,
		Description: model.Description,
		CategoryID:  model.CategoryID.Hex(),
	}

	return category, nil
}

func (r *MongoProductRepository) Validate(name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result := r.collection.FindOne(ctx, bson.M{"name": name})

	return result.Err()
}
