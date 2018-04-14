package birthday

import (
	"encoding/json"
	"fmt"
	"gomail/database"
	"log"
	"net/http"
	"time"
)

//AllBirthdays la puta que te pario
func AllBirthdays(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint AllBirthdays")
	contacts, err := database.AllBirthdays()
	if err != nil {
		log.Fatal(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	json.NewEncoder(w).Encode(contacts)
	log.Println("Sent all contacts")
}

//bodyBirthday para guardar el decode del body
type bodyBirthday struct {
	Name string
	Born string
}

//InsertBirthday es una cagada
func InsertBirthday(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "Please send any body", http.StatusBadRequest)
		return
	}

	var birthday database.Birthday
	var body bodyBirthday

	database.NewConn("root:@tcp(127.0.0.1:3306)/test?parseTime=true", "birthday").InitializeDatabase()

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	//preparo modelo
	birthday.Name = body.Name
	birthday.Born, err = time.Parse(time.RFC3339, body.Born)

	if err != nil {
		fmt.Println("Error while parsing date :", err)
	}

	//insertBirthday
	log.Println("Endpoint InsertBirdthday")
	err = database.InsertBirthday(&birthday)
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	log.Println("Inserted new birthday")
}
