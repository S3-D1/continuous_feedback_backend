package main

import (
	"encoding/json"
	"fmt"
	. "github.com/S3-D1/continuous_feedback_backend/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"net/http"
	"os"
	"strconv"
)

var db *gorm.DB
var err error

func main() {
	initDb()
	generateExampleData()
	startRestServer()
}

func initDb() {
	host := getenv("CF_DB_HOST", "0.0.0.0")
	port := getenv("CF_DB_PORT", "5432")
	schema := getenv("CF_DB_SCHEMA", "cf")
	user := getenv("CF_DB_USER", "postgres")
	pw := getenv("CF_DB_PW", "postgres")
	db, err = gorm.Open(
		"postgres",
		"host="+host+" port="+port+" user="+user+
			" dbname="+schema+" sslmode=disable password="+pw)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	db.AutoMigrate(&SingleFeedback{}, &Group{}, &User{})
}

func startRestServer() {
	router := mux.NewRouter()

	router.HandleFunc("/version", GetVersion).Methods("GET")
	router.HandleFunc("/feedbacks/{userId}", GetFeedbacksFromUser).Methods("GET")
	router.HandleFunc("/feedbacks", CreateFeedbacksFromUser).Methods("POST")
	router.HandleFunc("/feedback/{userId}", GetUserFeedback).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+getenv("CF_PORT", "13193"), router))
}

func CreateFeedbacksFromUser(writer http.ResponseWriter, request *http.Request) {
	var feedback SingleFeedback
	json.NewDecoder(request.Body).Decode(&feedback)
	db.Create(&feedback)
	json.NewEncoder(writer).Encode(&feedback)
}

func GetUserFeedback(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	userId := params["userId"]
	json.NewEncoder(writer).Encode(userId)
}

func GetFeedbacksFromUser(writer http.ResponseWriter, request *http.Request) {
	var feedbacks []SingleFeedback
	var user User
	params := mux.Vars(request)
	userId, _ := strconv.ParseUint(params["userId"], 10, 0)
	db.First(&user, userId)
	db.Preload("Author").Preload("Recipient").Where(&SingleFeedback{AuthorID: uint(userId)}).Find(&feedbacks)
	json.NewEncoder(writer).Encode(feedbacks)
}

func GetVersion(writer http.ResponseWriter, request *http.Request) {
	json.NewEncoder(writer).Encode("v.0.0.1")
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
