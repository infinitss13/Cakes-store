package repository

import (
	"context"
	"fmt"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/infinitss13/Cakes-store-card-service/config"
	"github.com/infinitss13/Cakes-store-card-service/entities"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/bson"
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

type UserCakes struct {
	UserID int             `json:"userId"`
	Cakes  []entities.Cake `json:"cakes"`
}

func (db MongoDatabase) CreateCard(userId int, Cake entities.Cake) error {
	var cakes []entities.Cake
	cakes = append(cakes, Cake)
	userCart := UserCakes{
		UserID: userId,
		Cakes:  cakes,
	}
	_, err := db.collection.InsertOne(context.TODO(), userCart)
	if err != nil {
		return err
	}
	return nil
}

func (db MongoDatabase) AddToCart(userID int, Cake entities.Cake) error {
	filter := bson.M{"userid": userID}

	// Define the update to add the given cake to the cakes array
	update := bson.M{"$push": bson.M{"cakes": Cake}}

	// Update the user document
	_, err := db.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (db MongoDatabase) IsCartEmpty(userID int) (bool, error) {
	filter := bson.M{"userid": userID}
	var userMap map[string]interface{}
	err := db.collection.FindOne(context.Background(), filter).Decode(&userMap)
	if err == mongo.ErrNoDocuments {
		return true, nil
	}
	if err != nil {
		return true, err
	}
	return false, nil
}

func (db MongoDatabase) GetCartItems(userId int) (map[string]interface{}, error) {
	filter := bson.M{"userid": userId}

	// Find the user document
	result := db.collection.FindOne(context.Background(), filter)
	if result.Err() != nil {
		return nil, result.Err()
	}

	// Decode the user document into a map
	var userMap map[string]interface{}
	err := result.Decode(&userMap)
	if err != nil {
		return nil, err
	}
	return userMap, nil

}
