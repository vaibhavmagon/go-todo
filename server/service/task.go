package service

import (
	"context"
	"fmt"
	"log"
    "go-todo/server/db"
	"go-todo/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Collection Created
var collection *mongo.Collection

func init(){
    var client *mongo.Client = db.DbConnect()
    collection = client.Database(db.DbName).Collection("todolist")
}

// get all task from the DB and return it
func GetAllTask() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var results []primitive.M

	for cur.Next(context.Background()) { // Loop on cur for all the values in it.
		var result bson.M
		errN := cur.Decode(&result)
		if errN != nil {
			log.Fatal(errN)
		}
		results = append(results, result)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.Background())
	return results
}

func GetOneTask(task string) bson.M{
    var result bson.M
    id, _ := primitive.ObjectIDFromHex(task)
	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

// Insert one task in the DB
func InsertOneTask(task models.ToDoList) {
	insertResult, err := collection.InsertOne(context.Background(), task)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single Record ", insertResult.InsertedID)
}

// task complete method, update task's status to true
func TaskComplete(task string, body models.ToDoList) int64{ // PUT
	fmt.Println(task)
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result)
	return result.ModifiedCount
}

// task undo method, update task's status to false
func UndoTask(task string) {
	fmt.Println(task)
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": false}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)
}

// task undo method, update task's status to false
func UpdateTaskData(task string, body models.ToDoList) int64 {
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": body.Status, "task": body.Task}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result)
	return result.ModifiedCount
}

// delete one task from the DB, delete by ID
func DeleteOneTask(task string) int64 {
	fmt.Println(task)
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	d, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted Document", d.DeletedCount)
	return d.DeletedCount
}

// delete all the tasks from the DB
func DeleteAllTask() int64 {
	d, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted Document", d.DeletedCount)
	return d.DeletedCount
}