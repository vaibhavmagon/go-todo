package db

import (
    "log"
    "context"
    "fmt"
    "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go-todo/server/utils"
)

var connectionString = "mongodb://" + utils.ViperEnvVariable("MONGO_INITDB_HOST") + ":27017"

// Database Name
const DbName = "todo"

var Client *mongo.Client

// create connection with mongo db
func init() {
    // Set client options
    clientOptions := options.Client().ApplyURI(connectionString)

    // connect to MongoDB
    client, err := mongo.Connect(context.TODO(), clientOptions)

    if err != nil {
        log.Fatal(err)
    }

    // Check the connection
    errP := client.Ping(context.TODO(), nil)

    if errP != nil {
        log.Fatal(errP)
    }else{
        fmt.Println("Connected to MongoDB!")
    }
    Client = client
}

func DbConnect() *mongo.Client{
    return Client
}

