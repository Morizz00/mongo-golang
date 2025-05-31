package main

import (
	"context"
	"fmt"
	"log"
	"mongo-golang/controllers"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Origin", "GET,POST,DELETE,PUT,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
func main() {
	client, err := getClient()
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	uc := controllers.NewUserController(client)

	r := httprouter.New()
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	r.GET("/users", uc.GetAllUsers)

	r.ServeFiles("/static/*filepath", http.Dir("frontend/"))
	r.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		http.ServeFile(w, r, "frontend/index.html")
	})

	fmt.Println("Server running at http://localhost:8081")
	fmt.Println("Frontend available at http://localhost:8081")
	fmt.Println("API endpoints:")
	fmt.Println("  GET    /user/:id")
	fmt.Println("  POST   /user")
	fmt.Println("  DELETE /user/:id")
	log.Fatal(http.ListenAndServe("localhost:8081", CORS(r)))
}

func getClient() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}
if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}
