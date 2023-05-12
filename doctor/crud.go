package doctor

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/hospital-management/models"
	"github.com/hospital-management/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateDoctor(w http.ResponseWriter, r *http.Request) {
	var doc models.Doctor
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&doc)
	if doc.Name == "" {
		http.Error(w, "name field can not be empty", http.StatusBadRequest)
		return
	}
	db := utils.Connection.Database(utils.HMDB)
	collection := db.Collection(utils.DoctorColl)

	res, err := collection.InsertOne(context.Background(), doc)
	if err != nil {
		log.Printf("error occurred while inserting doctor data, for name=%s, :: ERROR:%v\n", doc.Name, err)
		http.Error(w, "error while creating data", http.StatusInternalServerError)
		return
	}
	log.Printf("Doctor created successfully with id:%v\n", res.InsertedID)
	w.WriteHeader(http.StatusOK)
}

func DeleteDoc(w http.ResponseWriter, r *http.Request) {
	m := mux.Vars(r)
	name := m["name"]
	db := utils.Connection.Database(utils.HMDB)
	collection := db.Collection(utils.DoctorColl)
	filter := bson.M{"name": name}
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		http.Error(w, "delete one ended with error", http.StatusInternalServerError)
		return
	}
	log.Printf("doctor=[%s] deleted successfully", name)
	w.WriteHeader(200)
}
