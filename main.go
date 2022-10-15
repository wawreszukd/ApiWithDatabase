package main

import (
	"APIwithDatabase/dbHandler"
	"APIwithDatabase/endPointHandler"
	"fmt"
)

func main() {
	db, err := dbHandler.InitiateDataBase()
	if err != nil {
		fmt.Println(err)
	}

	defer db.Db.Close()
	endPointHandle.HandleRequests()
}
