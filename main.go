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
}

func main() {
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

	tmpl := template.Must(template.ParseFiles("temp.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		details := Subject{
			Gender: r.FormValue("gender"),
			Age:    r.FormValue("age"),
			Ioiha:  r.FormValue("ioiha"),
		}

		collection := client.Database("study").Collection("subject")

		insertResult, err := collection.InsertOne(context.TODO(), details)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Inserted a Document: ", insertResult.InsertedID)

		tmpl.Execute(w, struct{ Success bool }{true})
	})

	http.ListenAndServe(":8080", nil)
}
