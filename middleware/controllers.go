package middleware

import (
	"americanas/model"
	"americanas/repository"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

var planetCollection = repository.Db().Database("desafioAmericanas").Collection("planets")

func CreatePlanet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var planet model.Planet

	_ = json.NewDecoder(r.Body).Decode(&planet)
	planet.Id = primitive.NewObjectID()
	result, err := planetCollection.InsertOne(context.TODO(), planet)

	if err != nil {
		getError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func GetPlanet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var planets []model.Planet

	cur, err := planetCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		getError(err, w)
		return
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var planet model.Planet
		err := cur.Decode(&planet)

		if err != nil {
			log.Fatal(err)
		}

		planets = append(planets, planet)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(planets)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var planet model.Planet
	var params = mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex(params["id"])

	filter := bson.M{"id": id}
	err := planetCollection.FindOne(context.TODO(), filter).Decode(&planet)

	if err != nil {
		getError(err, w)
		return
	}

	json.NewEncoder(w).Encode(planet)
}

func getError(err error, w http.ResponseWriter) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
