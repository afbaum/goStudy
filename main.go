package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Subject struct {
	Gender string `json:"gender" bson:"gender"`
	Age    string `json:"age" bson:"age"`
	Ioiha  string `json:"ioiha" bson:"ioiha"`
	HearingAid
}

type HearingAid struct {
	Make  string `json:"make" bson:"make, omitempty"`
	Model string `json:"model" bson:"model, omitempty"`
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

	hearingAidInfo := HearingAid{
		Make:  r.FormValue("make"),
		Model: r.FormValue("model"),
	}

	details := Subject{
		Gender:     r.FormValue("gender"),
		Age:        r.FormValue("age"),
		Ioiha:      r.FormValue("ioiha"),
		HearingAid: hearingAidInfo,
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

func infoPage(w http.ResponseWriter, r *http.Request) {
	pipeline := pipeline("gender")
	tmpl := template.Must(template.ParseFiles("aids.html"))
	tmpl.ExecuteTemplate(w, "aids.html", pipeline)
}

// handle the different request function for this study
func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/form", dataEntry)
	http.HandleFunc("/infoPage", infoPage)

	http.ListenAndServe(":8080", nil)
}

type Data struct {
	Ioiha string
}

// queries data from the mongodB database
func pipeline(aidMake string) []string {
	client := MongoConnection()
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	database := client.Database("study")
	collection := database.Collection("subject")
	// collection := client.Database("study").Collection("subject")

	filterCursor, err := collection.Find(ctx, bson.M{"gender": "male"})
	if err != nil {
		log.Fatal(err)
	}
	var filteredData []bson.M
	if err = filterCursor.All(ctx, &filteredData); err != nil {
		log.Fatal(err)
	}

	s := make([]string, len(filteredData))

	for i := 0; i < len(filteredData); i++ {
		m := filteredData[i]["ioiha"].(string)
		s[i] = m
	}

	return s
}

func main() {
	handleRequests()
}
