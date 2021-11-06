package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Subject struct {
	Gender string `json:"gender" bson:"gender"`
	Age    string `json:"age" bson:"age"`
	Ioiha  string `json:"ioiha" bson:"ioiha"`
	Make   string `json:"make" bson:"make"`
	RightEar
	LeftEar
}

type RightEar struct {
	Right250  string `json:"r350" bson:"r350"`
	Right500  string `json:"r500" bson:"r500"`
	Right750  string `json:"r750" bson:"r750"`
	Right1000 string `json:"r1000" bson:"r1000"`
	Right1500 string `json:"r1500" bson:"r1500"`
	Right2000 string `json:"r2000" bson:"r2000"`
	Right3000 string `json:"r3000" bson:"r3000"`
	Right4000 string `json:"r4000" bson:"r4000"`
	Right6000 string `json:"r6000" bson:"r6000"`
	Right8000 string `json:"r8000" bson:"r8000"`
}

type LeftEar struct {
	Left250  string `json:"l350" bson:"l350"`
	Left500  string `json:"l500" bson:"l500"`
	Left750  string `json:"l750" bson:"l750"`
	Left1000 string `json:"l1000" bson:"l1000"`
	Left1500 string `json:"l1500" bson:"l1500"`
	Left2000 string `json:"l2000" bson:"l2000"`
	Left3000 string `json:"l3000" bson:"l3000"`
	Left4000 string `json:"l4000" bson:"l4000"`
	Left6000 string `json:"l6000" bson:"l6000"`
	Left8000 string `json:"l8000" bson:"l8000"`
}

// MongoConnection is a function that opens a connection to a MongoDB and returns the open connection so it can be used for CRUD operations
func MongoConnection() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb+srv://afbaum69:mic530gra@cluster0.fbpz2.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client
}

// Display the home page for the study
func homePage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("homepage.html"))
	tmpl.Execute(w, nil)
}

// Display the entry form for the study and collect the data into the database
func dataEntry(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("temp.html"))

	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	rightDetails := RightEar{
		Right250:  r.FormValue("r250"),
		Right500:  r.FormValue("r500"),
		Right750:  r.FormValue("r750"),
		Right1000: r.FormValue("r1000"),
		Right1500: r.FormValue("r1500"),
		Right2000: r.FormValue("r2000"),
		Right3000: r.FormValue("r3000"),
		Right4000: r.FormValue("r4000"),
		Right6000: r.FormValue("r6000"),
		Right8000: r.FormValue("r8000"),
	}

	leftDetails := LeftEar{
		Left250:  r.FormValue("l250"),
		Left500:  r.FormValue("l500"),
		Left750:  r.FormValue("l750"),
		Left1000: r.FormValue("l1000"),
		Left1500: r.FormValue("l1500"),
		Left2000: r.FormValue("l2000"),
		Left3000: r.FormValue("l3000"),
		Left4000: r.FormValue("l4000"),
		Left6000: r.FormValue("l6000"),
		Left8000: r.FormValue("l8000"),
	}

	details := Subject{
		Gender:   r.FormValue("gender"),
		Age:      r.FormValue("age"),
		Ioiha:    r.FormValue("ioiha"),
		Make:     r.FormValue("make"),
		RightEar: rightDetails,
		LeftEar:  leftDetails,
	}

	// use function to connect to DB and insert data
	client := MongoConnection()

	collection := client.Database("study").Collection("subject")

	insertResult, err := collection.InsertOne(context.TODO(), details)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a Document: ", insertResult.InsertedID)

	tmpl.Execute(w, nil)
}

// handle the different request function for this study
func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/form", dataEntry)
}

func main() {

	handleRequests()

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	http.ListenAndServe(":8080", nil)
}
