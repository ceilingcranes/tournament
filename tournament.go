package main

import (
	"fmt"
	"tournament/db"
)

/*
func main() {
	var inputData = "testdata.txt"
	bracketList := createBrackets(inputData)
	fmt.Println(bracketList)
}*/

func main() {
	database_val := db.NewDB()
	err := database_val.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Sucessfully connected")

	messages := db.ReadData()
	db.FillDB(database_val, messages)
	fmt.Println("We in there")
}
