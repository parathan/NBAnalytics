package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectDB connects to the MongoDB database using the provided URI.
// It returns a pointer to the MongoDB client.
// It logs an error and terminates the program if there is any error in the connection process.
//
// Returns:
// - *mongo.Client: A pointer to the MongoDB client.
func ConnectDB() *mongo.Client {
	// Create a context with a timeout of 10 seconds.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // Cancel the context when the function returns.

	// Connect to the MongoDB database using the provided URI.
	defer cancel()
	
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		log.Fatal(err) // Log the error and terminate the program if there is any error in the connection process.
		log.Fatal(err)
	}

	// Ping the MongoDB server to check if the connection is successful.
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err) // Log the error and terminate the program if there is any error in pinging the server.
		log.Fatal(err)
	}

	// Print a message indicating that the connection to MongoDB is successful.
	fmt.Println("Connected to MongoDB")

	// Return the MongoDB client.
	return client
}



// EnvMongoURI loads the environment variables from the .env file and returns the
// value of the MONGO_URI environment variable.
//
// It logs an error and terminates the program if there is any error in loading the
// .env file.
//
// Returns:
// - string: The value of the MONGO_URI environment variable.
func EnvMongoURI() string {
	// Load the environment variables from the .env file.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Return the value of the MONGO_URI environment variable.
	return os.Getenv("MONGO_URI")
}

var Mongo_Client *mongo.Client