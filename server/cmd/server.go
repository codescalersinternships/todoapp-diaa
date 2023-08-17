package main

import (
	"flag"
	"fmt"
	_ "github.com/codescalersinternships/todo-diaa/docs"
	"log"

	App "github.com/codescalersinternships/todo-diaa/internal"
)

// @title Todo API
// @description     Server for todo app.
// @host      localhost:5000
func main() {

	var dbPath string
	var port int

	flag.StringVar(&dbPath, "db", "./todo.db", "database file path")

	flag.IntVar(&port, "p", 5000, "port that will be used to run the app")

	flag.Parse()

	app, err := App.NewApp(dbPath)

	if err != nil {
		log.Fatal(err)
	}

	address := fmt.Sprintf(":%d", port)
	err = app.Run(address)

	if err != nil {
		log.Fatal(err)
	}

}
