package eployees

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

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var employees models.Employees
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&employees)
	if err != nil {
		log.Println("invalid input")
		return
	}
	collection := utils.Connection.Database(utils.HMDB).Collection(utils.Employees)
	succ, err := collection.InsertOne(context.TODO(), employees)
	if err != nil {
		log.Println("invalid input")
		http.Error(w, "invALid input", 400)
		return
	}
	log.Println(succ.InsertedID)
	w.WriteHeader(200)

}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	empl := mux.Vars(r)
	name := empl["name"]
	collection := utils.Connection.Database(utils.HMDB).Collection(utils.Employees)
	filter := bson.M{"name": name}
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		http.Error(w, "cannot delete this, as input is invalid", http.StatusBadRequest)
		return
	}
	log.Printf("employee=[%s] successfully deleted", name)
	w.WriteHeader(http.StatusOK)
}
