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
	if r.URL.Path != "/" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	db, err := dbHandler.InitiateDataBase()
	if err != nil {
		log.Fatal(err)
	}
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "static/Index.html")
	case "POST":
		http.ServeFile(w, r, "static/Index.html")
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm( err: %v", err)
			return
		}
		firstname := r.FormValue("FirstName")
		lastname := r.FormValue("LastName")
		gender := r.FormValue("Gender")
		date := r.FormValue("Date")
		db.HandleInsert(firstname, lastname, gender, date)

	}
	fmt.Println("HomePage endpoint hit")
}
