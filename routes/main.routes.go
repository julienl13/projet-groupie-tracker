package routes

import (
	"amiiboapi/controllers"
	"fmt"
	"net/http"
)

func MainServe() {
	http.HandleFunc("/", controllers.ListeAmiibo)
	fmt.Println("Serveur lancé : http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
