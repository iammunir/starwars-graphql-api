package main

import (
	"flag"
	"log"
	"os"

	"github.com/iammunir/starwars-graphql-api/database"
	"github.com/iammunir/starwars-graphql-api/seeder"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	flag.Parse()
	args := flag.Args()

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalln(err)
	}
	dbsql, err := db.DB()
	if err != nil {
		log.Fatalln(err)
	}
	defer dbsql.Close()

	if len(args) >= 1 {
		switch args[0] {
		case "seed":
			seeder.Execute(dbsql, args[1:]...)
			os.Exit(0)
		}
	}

}
