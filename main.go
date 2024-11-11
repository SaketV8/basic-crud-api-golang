package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"github.com/saketV8/basic-crud-api-golang/routes"
	"github.com/saketV8/basic-crud-api-golang/sqlite"
)

// for production
// gin.SetMode(gin.ReleaseMode)

// var rtr *gin.Engine = gin.Default()
// rtr.GET("/", handlers.HomePage)
// err := rtr.Run()

// using this so we can pass this down to handler function
// to query data, sherrrrr!

func main() {
	fmt.Println("-- CRUD API SERVICE --")
	fmt.Println("CREATED BY SAKET")
	fmt.Println()
	fmt.Println()

	// SETTING UP DATABASE
	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatal("ERORR IN DATABASE!!!")
		log.Fatal(err)
	}
	// Making database querable
	UserAccountsDbModel := &sqlite.UserAccountsDbModel{
		DB: db,
	}

	// Starting the server via Run() method
	// passing UserAccountsDbModel down to router
	// then router --> handler
	router := routes.SetupRouter(UserAccountsDbModel)
	err = router.Run()
	if err != nil {
		log.Fatal("ERORR IN SERVER!!!")
		log.Fatal(err)
	}
}
