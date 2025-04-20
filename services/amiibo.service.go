package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type AllAmiibo struct {
	Items []struct {
		Name   string `json:"name"`
		Id     string `json:"head"`
		Image  string `json:"image"`
		Series string `json:"gameSeries"`
	} `json:"amiibo"`
}

func GetAllAmiibo() (AllAmiibo, int, error) {
	var result AllAmiibo

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequest("GET", "https://amiiboapi.com/api/amiibo/", nil)
	if err != nil {
		return result, http.StatusInternalServerError, fmt.Errorf("erreur lors de la préparation de la requête : %s", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return result, http.StatusInternalServerError, fmt.Errorf("erreur lors de l'envoi de la requête : %s", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return result, res.StatusCode, fmt.Errorf("erreur dans la réponse : %d, message : %s", res.StatusCode, res.Status)
	}

	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return result, http.StatusInternalServerError, fmt.Errorf("erreur lors du décodage des données : %s", err)
	}

	return result, http.StatusOK, nil
}
