package main

import (
	"context"
	"encoding/json"
	"fmt"
	"rest_servers/models"
	"rest_servers/mongo_checks"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type TodoItem struct {
	Id   int
	Text string
	Tags []string
	Due  time.Time
}

func checkData() {
	client := mongo_checks.ConnectToMongoDB()
	coll := client.Database("ToDoListDev").Collection("todoLists")
	// filter := bson.D{{"text", "Learn Python"}}
	var results []models.Task
	cursor, err := coll.Find(context.TODO(), bson.D{{"id", 2}})
	if err != nil {
		fmt.Println(err)
	}
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	for _, result := range results {
		res, _ := json.Marshal(result)
		fmt.Println(string(res))
	}
	client.Disconnect(context.TODO())
}
