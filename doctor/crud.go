package doctor

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	"github.com/hospital-management/models"
	"github.com/hospital-management/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateDoctor(w http.ResponseWriter, r *http.Request) {
	var doc models.Doctor
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&doc)
	if err != nil {
		log.Println("JSON marshal err ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if doc.Name == "" {
		http.Error(w, "name field can not be empty", http.StatusBadRequest)
		return
	}
	doc.CreatedAt = time.Now()
	db := utils.Connection.Database(utils.HMDB)
	collection := db.Collection(utils.DoctorColl)

	res, err := collection.InsertOne(context.Background(), doc)
	if err != nil {
		http.Error(w, "error while creating data", http.StatusInternalServerError)
		return
	}
	log.Printf("Doctor created successfully with id:%v\n", res.InsertedID)
	w.WriteHeader(http.StatusCreated)
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
	w.WriteHeader(http.StatusOK)
}

// inp = doctor list, offset=10, limit
func ListDoctors(w http.ResponseWriter, r *http.Request) {
	// offset := r.URL.Query().Get("offset")
	limit := r.URL.Query().Get("limit")
	limitInt, _ := strconv.Atoi(limit)
	limitInt64 := int64(limitInt)
	collection := utils.Connection.Database(utils.HMDB).Collection(utils.DoctorColl)
	var options = options.FindOptions{Limit: &limitInt64}
	records, err := collection.Find(r.Context(), bson.M{}, &options)
	if err != nil {
		log.Fatal("record not found", err, http.StatusNotFound)
	}
	var foundedRecords []models.Doctor
	err = records.All(r.Context(), &foundedRecords)
	if err != nil {
		log.Fatal(err, http.StatusNotFound)
	}
	b, err := json.Marshal(foundedRecords)
	if err != nil {
		log.Fatal("error occured while converting into byte", err, http.StatusInternalServerError)
	}
	w.Write(b)
}

func UpdateDoctor(w http.ResponseWriter, r *http.Request) {
	m := mux.Vars(r)
	name := m["name"]
	contact := m["contact"]
	var docReq models.Doctor
	err := json.NewDecoder(r.Body).Decode(&docReq)
	if err != nil {
		log.Fatal("Error occured during updating doctor's name", "error =", err, http.StatusInternalServerError)
	}
	db := utils.Connection.Database(utils.HMDB)
	collection := db.Collection(utils.DoctorColl)

	filter := bson.M{"name": name, "contact": contact}
	res := collection.FindOne(r.Context(), filter)
	if res.Err() != nil {
		log.Fatal("Error occured during finding doctor's name from DB", "error =", err, http.StatusNotFound)
	}
	var docDB models.Doctor
	err = res.Decode(&docDB)
	if err != nil {
		log.Fatal("Error occured during decoding doctor's name of DB", "error =", err)
	}

	// 	res = collection.FindOneAndUpdate(r.Context(), filter, docReq)
	// 	if res.Err() != nil {
	// log.Fatal("Error occured during updating doctor's name", "error =", err)
	// 	}
	docDB.Name = docReq.Name
	docDB.Contact = docReq.Contact
	docDB.UpdatedAt = time.Now()
	_,err=collection.UpdateOne(context.TODO(), filter, docDB)
	if err !=nil{
log.Fatal("error occured during updating the record,  error =", err, http.StatusInternalServerError)
	}
w.WriteHeader(http.StatusOK)
}
