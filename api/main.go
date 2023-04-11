package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/users", getUsers)
	fmt.Println("api is on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

type User struct {
	ID   int
	Name string
}

func getUsers(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content=Type", "aplication/json")
	json.NewEncoder(w).Encode([]User{
		{
			ID:   1,
			Name: "Marcio Sena",
		},
		{
			ID:   5,
			Name: "Tonico",
		},
	})
}
