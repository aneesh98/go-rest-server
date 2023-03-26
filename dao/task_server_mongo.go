package dao

// Mongo Implementation
import (
	"context"
	"rest_servers/models"
	"rest_servers/mongo_checks"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskStore struct {
	tasks  map[int]models.Task
	nextId int
}

var mongoClient *mongo.Client
var coll *mongo.Collection

func InitConn() {
	mongoClient = mongo_checks.ConnectToMongoDB()
	coll = mongoClient.Database("ToDoListDev").Collection("todoLists")
}

func (ts *TaskStore) GetTask(id int) (models.Task, error) {
	filter := bson.D{{"id", id}}
	var task models.Task

	err := coll.FindOne(context.TODO(), filter).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Task{}, nil
		}
		panic(err)
	}

	return task, nil
}
