package mongo

// https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo

import (
	"common"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DbClient *mongo.Client

const MONGO_URI string = "mongodb://localhost:27017"

// type logEntity struct {
// 	logType  string `bson:"logType"`
// 	logValue string `bson:"logValue"`
// 	dbPrefix string `bson:"dbPrefix"`
// 	time     string `bson:"time"`
// }

func Connect() *mongo.Client {

	if DbClient != nil {
		return DbClient
	}

	// credential := options.Credential{
	// 	Username: "<USER_NAME>",
	// 	Password: "<PASSWORD>",
	// }
	// clientOptions := options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(credential)
	clientOptions := options.Client().ApplyURI(MONGO_URI)
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

func InsertOne(item common.MongoItem) {
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

	// list_ids := result.InsertedIDs
	// fmt.Printf("Documents inserted: %v\n", len(list_ids))
	return len(result.InsertedIDs)

}

// mongoDB에 저장된 자격증 개수 조회
func GetCountCollectionCount() []int {
	var countList []int = make([]int, common.EndSeriseCode)

	const databaseName string = "go-project"
	const collectionName string = "certification"

	coll := DbClient.Database(databaseName).Collection(collectionName)

	/*
		bson.D: 하나의 BSON 도큐멘트. MongoDB command 처럼 순서가 중요한 경우에 사용합니다.
		bson.M: 순서가 없는 map 형태. 순서를 유지하지 않는다는 점을 빼면 D와 같습니다.
		bson.A: 하나의 BSON array 형태.
		bson.E: D 타입 내부에서 사용하는 하나의 엘레멘트.
	*/
	for i := common.Serise(common.BeginSeriseCode) + 1; i < common.Serise(common.EndSeriseCode); i++ {
		filter := bson.D{{"seriesNm", common.Serise(i).String()}}
		certCount, err := coll.CountDocuments(context.TODO(), filter)
		if err != nil {
			fmt.Println("GetCountCollectionCount() error (", err.Error(), ")")
			log.Fatal(err)
		}
		fmt.Println("certificate count in mongoDb:", certCount)
		countList[i] = int(certCount)
	}

	for i, v := range countList {
		fmt.Println(i, common.Serise(i).String(), v)
	}

	return countList
}
