package repository

import (
	"context"

	"github.com/rbcorrea/meli-test/internal/domain/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	Collection *mongo.Collection
}

func NewMongoRepository(db *mongo.Database) *MongoRepository {
	return &MongoRepository{
		Collection: db.Collection("short_urls"),
	}
}

func (r *MongoRepository) Save(ctx context.Context, shortURL *entity.ShortURL) error {
	_, err := r.Collection.InsertOne(ctx, shortURL)
	return err
}

func (r *MongoRepository) FindByCode(ctx context.Context, code string) (*entity.ShortURL, error) {
	var result entity.ShortURL
	err := r.Collection.FindOne(ctx, bson.M{"code": code}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// func (r *MongoRepository) UpdateAccessData(ctx context.Context, code string, accessed int64, lastAccess time.Time) error {
// 	_, err := r.Collection.UpdateOne(
// 		ctx,
// 		bson.M{"code": code},
// 		bson.M{"$set": bson.M{
// 			"accessed":    accessed,
// 			"last_access": lastAccess,
// 		}},
// 	)
// 	return err
// }

func (r *MongoRepository) DeactivateByCode(ctx context.Context, code string) error {
	_, err := r.Collection.UpdateOne(
		ctx,
		bson.M{"code": code},
		bson.M{"$set": bson.M{"is_active": false}},
	)
	return err
}
