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
var userCollection *mongo.Collection

func init(){
    var client *mongo.Client = db.DbConnect()
    userCollection = client.Database(db.DbName).Collection("user")
}

// get all users from the DB and return it
func GetAllUsers() []primitive.M {
	cur, err := userCollection.Find(context.Background(), bson.D{{}})
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

func GetOneUser(user string) bson.M{
    var result bson.M
    id, _ := primitive.ObjectIDFromHex(user)
	err := userCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

// Insert one task in the DB
func InsertUser(user models.User) {
	insertResult, err := userCollection.InsertOne(context.Background(), user)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single Record ", insertResult.InsertedID)
}

// user update method, update task's status to false
func UpdateUser(user string, body models.User) int64 {
	id, _ := primitive.ObjectIDFromHex(user)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"name": body.Name, "mobile": body.Mobile}}
	result, err := userCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result)
	return result.ModifiedCount
}

// delete one task from the DB, delete by ID
func DeleteUser(user string) int64 {
	fmt.Println(user)
	id, _ := primitive.ObjectIDFromHex(user)
	filter := bson.M{"_id": id}
	d, err := userCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted Document", d.DeletedCount)
	return d.DeletedCount
}