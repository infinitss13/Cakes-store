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
	doc := bson.D{
		primitive.E{
			Key: "ID", Value: 1,
		},
		primitive.E{
			Key:   "Title",
			Value: "Вечные 22",
		},
		primitive.E{
			Key:   "Price",
			Value: 34,
		},
		primitive.E{
			Key:   "ImgUrl",
			Value: "https://sun9-40.userapi.com/impg/3wRh8WzEZzvoWTm-saAAZSGzE24IF6_D1AAQ9g/tu6RfaSZedA.jpg?size=736x743&quality=95&sign=1ff680ee89970f8ccb08b590169cc2ef&type=album",
		},
		primitive.E{
			Key:   "Description",
			Value: "Один из самых вкусных тортов, который напомнит, что возраст это просто цифра",
		},
		primitive.E{
			Key:   "BiscuitType",
			Value: "Карамельный",
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
			Value: "Клубника",
		},
		primitive.E{
			Key:   "Weight",
			Value: "0.5 кг",
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

	doc = bson.D{
		primitive.E{
			Key: "ID", Value: 2,
		},
		primitive.E{
			Key:   "Title",
			Value: "А мне всё равно",
		},
		primitive.E{
			Key:   "Price",
			Value: 42,
		},
		primitive.E{
			Key:   "ImgUrl",
			Value: "https://sun1.velcom-by-minsk.userapi.com/impg/pJCbfvaCalEcz72QHjxofBdiqawDbNxUEX7RrA/R5pwU7c11is.jpg?size=564x564&quality=95&sign=b030e82a6a3971d8fe76b1e2da028a00&type=album",
		},
		primitive.E{
			Key:   "Description",
			Value: "Нежный и утонченный торт, который подарит радость и наслаждение",
		},
		primitive.E{
			Key:   "BiscuitType",
			Value: "Шоколадный",
		},
		primitive.E{
			Key:   "CreamType",
			Value: "Сметанный",
		},
		primitive.E{
			Key:   "ToppingType",
			Value: "Орехи",
		},
		primitive.E{
			Key:   "FillingType",
			Value: "Кокосовая начинка",
		},
		primitive.E{
			Key:   "Berries",
			Value: "Голубика",
		},
		primitive.E{
			Key:   "Weight",
			Value: "0.8 кг",
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

	doc = bson.D{
		primitive.E{
			Key: "ID", Value: 3,
		},
		primitive.E{
			Key:   "Title",
			Value: "Мечтатель",
		},
		primitive.E{
			Key:   "Price",
			Value: 51,
		},
		primitive.E{
			Key:   "ImgUrl",
			Value: "https://sun9-31.userapi.com/impg/CFlzyxJ1l5uUO4YRv4MfBhHzJythaVxnU-UPwA/NlUlRPimU9A.jpg?size=564x564&quality=95&sign=781749c1e57f6d80733ca2dd46ceba8c&type=album",
		},
		primitive.E{
			Key:   "Description",
			Value: "Пусть все ваши мечты осуществятся",
		},
		primitive.E{
			Key:   "BiscuitType",
			Value: "Ванильный",
		},
		primitive.E{
			Key:   "CreamType",
			Value: "Масляный",
		},
		primitive.E{
			Key:   "ToppingType",
			Value: "Бананы",
		},
		primitive.E{
			Key:   "FillingType",
			Value: "Карамель",
		},
		primitive.E{
			Key:   "Berries",
			Value: "Голубика",
		},
		primitive.E{
			Key:   "Weight",
			Value: "0.7 кг",
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

	doc = bson.D{
		primitive.E{
			Key: "ID", Value: 4,
		},
		primitive.E{
			Key:   "Title",
			Value: "Студенческая жизнь",
		},
		primitive.E{
			Key:   "Price",
			Value: 35,
		},
		primitive.E{
			Key:   "ImgUrl",
			Value: "https://sun9-42.userapi.com/impg/qnBg0Kpt0NaZV2_G5tqdqwQ3spb-TIYMaGzmiw/Hg1lscz47gI.jpg?size=564x752&quality=95&sign=422fb49f4d6c5582c9c68c2d5e0c9c8d&type=album",
		},
		primitive.E{
			Key:   "Description",
			Value: "Яркие воспоминания о студенческой жизни не оставят вас равнодушными. Ох, как бы в то время пригодился такой тортик",
		},
		primitive.E{
			Key:   "BiscuitType",
			Value: "Шоколадный",
		},
		primitive.E{
			Key:   "CreamType",
			Value: "Карамель",
		},
		primitive.E{
			Key:   "ToppingType",
			Value: "Карамель",
		},
		primitive.E{
			Key:   "FillingType",
			Value: "Шоколадные гранолы",
		},
		primitive.E{
			Key:   "Berries",
			Value: "Малина",
		},
		primitive.E{
			Key:   "Weight",
			Value: "0.6 кг",
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

	doc = bson.D{
		primitive.E{
			Key: "ID", Value: 5,
		},
		primitive.E{
			Key:   "Title",
			Value: "Влюбленность",
		},
		primitive.E{
			Key:   "Price",
			Value: 50,
		},
		primitive.E{
			Key:   "ImgUrl",
			Value: "https://sun9-37.userapi.com/impg/Ky4rCrodG0-fPEezpSI6zjm5EKJu6O3j1SCwiQ/Z1AtaUyW8q8.jpg?size=564x752&quality=95&sign=c361be94727cbeab1e1a5e1471d6fa4f&type=album",
		},
		primitive.E{
			Key:   "Description",
			Value: "Порой влюбленность заставляет нас совершать безумные поступки, поверьте, в этот тортик вы влюбитесь",
		},
		primitive.E{
			Key:   "BiscuitType",
			Value: "Шоколадный",
		},
		primitive.E{
			Key:   "CreamType",
			Value: "Карамель",
		},
		primitive.E{
			Key:   "ToppingType",
			Value: "Карамель",
		},
		primitive.E{
			Key:   "FillingType",
			Value: "Шоколадные гранолы",
		},
		primitive.E{
			Key:   "Berries",
			Value: "Малина",
		},
		primitive.E{
			Key:   "Weight",
			Value: "0.6 кг",
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

	doc = bson.D{
		primitive.E{
			Key: "ID", Value: 6,
		},
		primitive.E{
			Key:   "Title",
			Value: "Отряд звериных",
		},
		primitive.E{
			Key:   "Price",
			Value: 45,
		},
		primitive.E{
			Key:   "ImgUrl",
			Value: "https://sun9-65.userapi.com/impg/Qko0LhAAxebP7HyJqRpagUO7f_Z8GLA4OM8iHA/Cm_HgM5F7kM.jpg?size=564x752&quality=95&sign=8f7a8d13b2e7b54f5768060ad2207d8e&type=album",
		},
		primitive.E{
			Key:   "Description",
			Value: "После одного укуса данного тортика вы будете более могущественным, чем любой лев или тигр",
		},
		primitive.E{
			Key:   "BiscuitType",
			Value: "Шоколадный",
		},
		primitive.E{
			Key:   "CreamType",
			Value: "Сливочный",
		},
		primitive.E{
			Key:   "ToppingType",
			Value: "Карамель",
		},
		primitive.E{
			Key:   "FillingType",
			Value: "Шоколадные гранолы",
		},
		primitive.E{
			Key:   "Berries",
			Value: "Клубника",
		},
		primitive.E{
			Key:   "Weight",
			Value: "0.7 кг",
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

	doc = bson.D{
		primitive.E{
			Key: "ID", Value: 7,
		},
		primitive.E{
			Key:   "Title",
			Value: "Нежно-зеленый",
		},
		primitive.E{
			Key:   "Price",
			Value: 38,
		},
		primitive.E{
			Key:   "ImgUrl",
			Value: "https://sun9-15.userapi.com/impg/T5o6cnU_Rbsj6vzisI09xSYZdVKLq4uj2Qhg6A/M0dabMx1SYU.jpg?size=679x679&quality=95&sign=4884121e99233d7d9ba3896f5c38e9aa&type=album",
		},
		primitive.E{
			Key:   "Description",
			Value: "Что может быть приятнее, чем получить такой милы тортик на день рождения. Обязательно задумайтесь насчет такого подарка кому нибудь",
		},
		primitive.E{
			Key:   "BiscuitType",
			Value: "Ванильный",
		},
		primitive.E{
			Key:   "CreamType",
			Value: "Масляный",
		},
		primitive.E{
			Key:   "ToppingType",
			Value: "Фисташка",
		},
		primitive.E{
			Key:   "FillingType",
			Value: "Шоколадные гранолы",
		},
		primitive.E{
			Key:   "Berries",
			Value: "Клубника",
		},
		primitive.E{
			Key:   "Weight",
			Value: "0.8 кг",
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

	doc = bson.D{
		primitive.E{
			Key: "ID", Value: 8,
		},
		primitive.E{
			Key:   "Title",
			Value: "Космический",
		},
		primitive.E{
			Key:   "Price",
			Value: 42,
		},
		primitive.E{
			Key:   "ImgUrl",
			Value: "https://sun9-79.userapi.com/impg/tZC42A5YKIxvnALn487Sen90o6Bv8jev_cmRvw/3bsCPM_G7dU.jpg?size=1200x660&quality=95&sign=45616c16410cb9f6c795cd08a7ca0143&type=album",
		},
		primitive.E{
			Key:   "Description",
			Value: "С таким тортиком не страшны никакие приключения, хоть в магазин, хоть до луны",
		},
		primitive.E{
			Key:   "BiscuitType",
			Value: "Карамельный",
		},
		primitive.E{
			Key:   "CreamType",
			Value: "Сливочный",
		},
		primitive.E{
			Key:   "ToppingType",
			Value: "Джем",
		},
		primitive.E{
			Key:   "FillingType",
			Value: "Шоколадные гранолы",
		},
		primitive.E{
			Key:   "Berries",
			Value: "Клубника",
		},
		primitive.E{
			Key:   "Weight",
			Value: "0.8 кг",
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

func (database MongoDatabase) GetCatalogWithLimit(limit int) ([]entities.Cake, error) {
	cursor, err := database.collection.Find(context.Background(), bson.D{}, options.Find().SetLimit(int64(limit)))
	if err != nil {
		return nil, fmt.Errorf("AAAAAAAAA : %w", err)
	}

	var cakes []entities.Cake
	for cursor.Next(context.Background()) {
		var cake entities.Cake
		if err = cursor.Decode(&cake); err != nil {
			return nil, err
		}
		cakes = append(cakes, cake)
	}
	return cakes, nil
}
