package routes

import (
	"log"
	"net/http"

	"github.com/jeffqev/go-microservice-mvc/controllers"
)

// ServerUP .
func ServerUP() {
	http.HandleFunc("/users", controllers.GetUser)

	if err := http.ListenAndServe(":1323", nil); err != nil {
		log.Fatal(err)
	}
}
