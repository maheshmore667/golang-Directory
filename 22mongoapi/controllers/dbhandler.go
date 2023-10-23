package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maheshmore667/mongoapi/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017"

var dbName = "OttApp"
var collectionName = "Watchlist"
var collection *mongo.Collection

func init() {

	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mongo connected successfully")
	collection = client.Database(dbName).Collection(collectionName)
	fmt.Println("Collection instance is ready")
}

//mongoDb helpers : separate file

func insertOneRecord(movie models.OttApp) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Record is inserted successfully : ", inserted.InsertedID)
}

func updateOneRecord(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	result, updateErr := collection.UpdateOne(context.Background(), filter, update)
	if updateErr != nil {
		log.Fatal(updateErr)
	}
	fmt.Println("Updated Count : ", result.ModifiedCount)

}

func deleteOneRecord(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	deletedRecord, deleteErr := collection.DeleteOne(context.Background(), filter)
	if deleteErr != nil {
		log.Fatal(deleteErr)
	}
	fmt.Println("DeletedCount : ", deletedRecord.DeletedCount)
}

func deleteAllRecord() {
	deletedRecords, deleteErr := collection.DeleteMany(context.Background(), bson.D{{}})
	if deleteErr != nil {
		log.Fatal(deleteErr)
	}
	fmt.Println("DeletedCount : ", deletedRecords.DeletedCount)
}

func getAllMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M
	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cursor.Close(context.Background())
	return movies
}

// Actual controller file

func GetPlayList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("application-type", "application/json")
	playList := getAllMovies()
	json.NewEncoder(w).Encode(playList)
}

func CreateOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("application-type", "application/json")
	w.Header().Set("allow-control-allow-method", "POST")
	var movie models.OttApp
	json.NewDecoder(r.Body).Decode(&movie)
	insertOneRecord(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("application-type", "application/json")
	w.Header().Set("allow-control-allow-method", "POST")
	params := mux.Vars(r)
	updateOneRecord(params["id"])
	json.NewEncoder(w).Encode("updated the movie.")
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("application-type", "application/json")
	w.Header().Set("allow-control-allow-method", "DELETE")
	params := mux.Vars(r)
	deleteOneRecord(params["id"])
	json.NewEncoder(w).Encode("Deleted the movie.")
}

func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("application-type", "application/json")
	w.Header().Set("allow-control-allow-method", "DELETE")
	deleteAllRecord()
	json.NewEncoder(w).Encode("Wooho !! All movies gone :)")
}
