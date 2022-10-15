package endPointHandle

import (
	"APIwithDatabase/dbHandler"
	"fmt"
	"log"
	"net/http"
)

func HandleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/all", returnAllPersons)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func returnAllPersons(w http.ResponseWriter, r *http.Request) {
	db, err := dbHandler.InitiateDataBase()
	if err != nil {
		log.Fatal(err)
	}
	err, outstring := db.HandleSelect()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, outstring)
}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>HomePage</h1>")
	fmt.Println("HomePage endpoint hit")
}
