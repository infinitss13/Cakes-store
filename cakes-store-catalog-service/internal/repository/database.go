package repository

import (
	"context"
	"fmt"

	"github.com/infinitss13/Cakes-store-catalog-service/config"
	"github.com/infinitss13/Cakes-store-catalog-service/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDatabase struct {
	collection *mongo.Collection
}

func NewDatabase(database *mongo.Database) MongoDatabase {
	mongoDB := MongoDatabase{
		collection: database.Collection(config.NewConnectionMongo().MongoCollection),
	}
	mongoDB.fullfillCatalog()
	return mongoDB
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
func (database MongoDatabase) fullfillCatalog() {
	var result entities.Cake
	err := database.collection.FindOne(context.Background(), bson.D{}).Decode(&result)
	if err != mongo.ErrNoDocuments {
		return
	}
	for i := 1; i < 10; i++ {
		doc := bson.D{
			primitive.E{
				Key: "ID", Value: i,
			},
			primitive.E{
				Key:   "Title",
				Value: "Торт 1",
			},
			primitive.E{
				Key:   "Price",
				Value: i * 100,
			},
			primitive.E{
				Key:   "ImgUrl",
				Value: "https://example.com/img1.png",
			},
			primitive.E{
				Key:   "Description",
				Value: "Description",
			},
			primitive.E{
				Key:   "BiscuitType",
				Value: "Шоколадный 1",
			},
			primitive.E{
				Key:   "CreamType",
				Value: "Крем-брюле",
			},
			primitive.E{
				Key:   "ToppingType",
				Value: "Фрукты",
			},
			primitive.E{
				Key:   "FillingType",
				Value: "Карамель",
			},
			primitive.E{
				Key:   "Berries",
				Value: "Малина",
			},
			primitive.E{
				Key:   "Weight",
				Value: "1 кг",
			},
			primitive.E{
				Key:   "IsCustom",
				Value: false,
			},
			primitive.E{
				Key:   "CustomText",
				Value: "",
			},
		}
		_, err = database.collection.InsertOne(context.TODO(), doc)
	}
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
