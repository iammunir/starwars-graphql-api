package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/iammunir/starwars-graphql-api/app"
	"github.com/iammunir/starwars-graphql-api/config"
	"github.com/iammunir/starwars-graphql-api/database"
	"github.com/iammunir/starwars-graphql-api/logger"
	"github.com/iammunir/starwars-graphql-api/seeder"
	"github.com/spf13/viper"
)

func init() {
	config.InitConfig("./")
}

func main() {

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

	log := logger.InitLogger()

	app := app.InitApp(db, log)

	port := viper.GetString("APP_PORT")
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: app,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("error server initializing: %s", err.Error())
		}
	}()

	log.Info("server listening at port: ", port)

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Info("shutdown server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Error("server shutdown with error: %s", err)
	}
	log.Info("server exiting")

}
