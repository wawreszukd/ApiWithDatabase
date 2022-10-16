package endPointHandle

import (
	"APIwithDatabase/dbHandler"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func HandleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/", homePage).Methods("GET")
	r.HandleFunc("/person", handleForm).Methods("GET")
	r.HandleFunc("/person", handleCreate).Methods("POST")
	r.HandleFunc("/persons", handleGetAll)
	r.HandleFunc("/person/{id}", handleEdit).Methods("PUT")
	r.HandleFunc("/person/{id}", handleDelete).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))
}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("homePage endpoint hit")
	http.ServeFile(w, r, "static/index.html")
}
func handleEdit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	intVar, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal(err)
	}
	db, err := dbHandler.InitiateDataBase()
	if err != nil {
		log.Fatal(err)
	}
	firstname := r.FormValue("first_name")
	lastname := r.FormValue("last_name")
	gender := r.FormValue("gender")
	dateofbirth := r.FormValue("date_of_birth")
	db.HandleEdit(firstname, lastname, gender, dateofbirth, intVar)
	http.Redirect(w, r, "/persons", http.StatusSeeOther)
}
func handleDelete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handleDelete endpoint hit")
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	intVar, err := strconv.Atoi(vars["id"])
	db, err := dbHandler.InitiateDataBase()
	if err != nil {
		log.Fatal(err)
	}
	db.HandleDelete(intVar)
	http.Redirect(w, r, "/persons", http.StatusSeeOther)
}
func handleGetAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handleGetAll endpoint hit")
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
func handleForm(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handleForm endpoint hit")
	if r.URL.Path != "/person" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		fmt.Fprintf(w, "404 Not Found")
	}
	http.ServeFile(w, r, "static/form.html")
	fmt.Println("HomePage endpoint hit")
}
func handleCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handleCreate endpoint hit")
	db, err := dbHandler.InitiateDataBase()
	if err != nil {
		log.Fatal(err)
	}
	if err = r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	firstname := r.FormValue("FirstName")
	lastname := r.FormValue("LastName")
	gender := r.FormValue("Gender")
	date := r.FormValue("Date")
	db.HandleInsert(firstname, lastname, gender, date)
	http.Redirect(w, r, "/persons", http.StatusSeeOther)
}
