package patient

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hospital-management/models"
	"github.com/hospital-management/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func CreatePatient(w http.ResponseWriter, r *http.Request) {
	var pat models.Patient
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&pat)
	if err != nil {
		log.Println("invalid input")
		return
	}
	db := utils.Connection.Database(utils.HMDB)
	collection := db.Collection(utils.Patients)
	succ, err := collection.InsertOne(r.Context(), pat)
	if err != nil {
		log.Println("name of patient is not found")
		http.Error(w, "error while creating data", http.StatusInternalServerError)
		return
	}
	log.Println("patient successfully added", succ.InsertedID)
	w.WriteHeader(200)

}

func DeletePatient(w http.ResponseWriter, r *http.Request) {
	A := mux.Vars(r)
	name := A["name"]
	collection := utils.Connection.Database(utils.HMDB).Collection(utils.Patients)
	filter := bson.M{"name": name}
	_, err := collection.DeleteOne(r.Context(), filter)
	if err != nil {
		http.Error(w, "can not delete this one", 400)
		return
	}
	log.Println("deleted successfully", name)
	w.WriteHeader(200)
}
