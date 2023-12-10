package mongo

// https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo

import (
	"common"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DbClient *mongo.Client

// type logEntity struct {
// 	logType  string `bson:"logType"`
// 	logValue string `bson:"logValue"`
// 	dbPrefix string `bson:"dbPrefix"`
// 	time     string `bson:"time"`
// }

func MongoConn() *mongo.Client {

	if DbClient != nil {
		return DbClient
	}

	// credential := options.Credential{
	// 	Username: "<USER_NAME>",
	// 	Password: "<PASSWORD>",
	// }
	// clientOptions := options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(credential)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	DbClient, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = DbClient.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("[Step1] MongoDB Connection Complete!")
	return DbClient
}

func MongoInsertOne(item common.MongoItem) {
	coll := DbClient.Database("go-project").Collection("certification")

	result, err := coll.InsertOne(context.TODO(), item)
	if err != nil {
		fmt.Printf("db insert error with JmNm: %s\n", item.JmNm)
		return
	}

	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
}

func InsertMany(items []interface{}) int {
	const databaseName string = "go-project"
	const collectionName string = "certification"

	coll := DbClient.Database(databaseName).Collection(collectionName)

	result, err := coll.InsertMany(context.TODO(), items)
	if err != nil {
		fmt.Println("db insert error (", err.Error(), ")")
		log.Fatal(err)
	}

	list_ids := result.InsertedIDs
	fmt.Printf("Documents inserted: %v\n", len(list_ids))
}
