package controllers

import (
	"amiiboapi/services"
	"fmt"
	"html/template"
	"net/http"
)

func ListeAmiibo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("📥 Requête reçue")

	data, statusCode, err := services.GetAllAmiibo()
	if err != nil {
		fmt.Println("❌ Erreur récupération Amiibo:", err)
		http.Error(w, fmt.Sprintf("Erreur : %s", err), statusCode)
		return
	}

	tmpl, err := template.ParseFiles("./templates/listAmiibo.html")
	if err != nil {
		fmt.Println("❌ Erreur chargement template:", err)
		http.Error(w, fmt.Sprintf("Erreur template : %s", err), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Println("❌ Erreur rendering template:", err)
		http.Error(w, fmt.Sprintf("Erreur exécution template : %s", err), http.StatusInternalServerError)
		return
	}

	fmt.Println("✅ Données affichées avec succès")
}
