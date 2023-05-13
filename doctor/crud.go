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

func UpdateDoctor(w http.ResponseWriter, r *http.Request) {
	m := mux.Vars(r)
	name := m["name"]
	var doc models.Doctor
	err := json.NewDecoder(r.Body).Decode(&doc)
	if err != nil {
		log.Fatal("Error occured during updating doctor's name", "error =", err)
	}
	db := utils.Connection.Database(utils.HMDB)
	collection := db.Collection(utils.DoctorColl)

	filter := bson.M{"name": name}
	res := collection.FindOne(r.Context(), filter)
	if res.Err() != nil {
		//
	}
	var docDB models.Doctor
	err = res.Decode(&docDB)
	if err != nil {
		//
	}

	if doc.Name == collection.Name() && doc.Contact == collection.Contact {
		updateID, err := collection.UpdateOne(context.TODO(), collection.Name(), doc.Name)
		if err != nil {
			log.Printf("error in inserting doctor's data  :: Error %s", err)
			return
		}
		log.Printf("docotor's data is updated successfully", "update ID = %s", updateID)
		w.WriteHeader(http.StatusOK)
	}

}
