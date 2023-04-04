package repository

import (
	"context"
	"fmt"

	"github.com/infinitss13/Cakes-store-catalog-service/config"
	"github.com/infinitss13/Cakes-store-catalog-service/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDatabase struct {
	collection *mongo.Collection
}

func NewDatabase(database *mongo.Database) MongoDatabase {
	return MongoDatabase{
		collection: database.Collection(config.NewConnectionMongo().MongoCollection),
	}
}

func NewClientMongo(ctx context.Context) (db *mongo.Database, err error) {
	var mongoDBURL string
	newConnection := config.NewConnectionMongo()

	mongoDBURL = fmt.Sprintf("mongodb://%s:%s", newConnection.MongoHost, newConnection.MongoPort)
	clientOptions := options.Client().ApplyURI(mongoDBURL)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	return client.Database(newConnection.MongoDBName), nil
}

func (database MongoDatabase) GetCatalog() ([]entities.Cake, error) {
	cursor, err := database.collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, fmt.Errorf("AAAAAAAAA : %w", err)
	}

	var cakes []entities.Cake
	for cursor.Next(context.Background()) {
		var cake entities.Cake
		if err := cursor.Decode(&cake); err != nil {
			return nil, err
		}
		cakes = append(cakes, cake)
	}
	return cakes, nil
}
