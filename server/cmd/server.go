package main

import (
	"flag"
	"fmt"
	"log"
	_ "github.com/codescalersinternships/todo-diaa/docs"

	App "github.com/codescalersinternships/todo-diaa/internal"
)

// @title Todo API
// @description     Server for todo app.
// @host      localhost:5000
func main() {

	var db_path string
	var port int

	flag.StringVar(&db_path, "db", "./todo.db", "database file path")

	flag.IntVar(&port, "p", 5000, "port that will be used to run the app")

	flag.Parse()

	app := App.NewApp()

	err := app.StartDB(db_path)

	if err != nil {
		log.Fatal(err)
	}

	address := fmt.Sprintf(":%d", port)
	err = app.Run(address)

	if err != nil {
		log.Fatal(err)
	}

}
