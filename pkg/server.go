package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/GodlyDad/Backend/pkg/graph"
	"github.com/GodlyDad/Backend/pkg/graph/model"
	"github.com/GodlyDad/Backend/pkg/helpers"
	"github.com/GodlyDad/Backend/pkg/service"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

/*
TODO:
 1. schema
 2. imports/docker (gorm, postgres, zap logger? etc)
 2. connection to resolvers.
 3. functions... GORM ect.
 4. data
 5. auto-population functions
*/

func main() {
	godotenv.Load()

	config := &helpers.Config{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		User:     os.Getenv("POSTGRES_USER"),
		SSLMode:  os.Getenv("POSTGRES_SSLMODE"),
		DBName:   os.Getenv("POSTGRES_NAME"),
	}
	db, err := helpers.NewConnection(config)
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(1)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(1)

	db.AutoMigrate(&model.Book{}, &model.Translation{}, &model.Chapter{})

	if err := helpers.PopulateTranslations(db); err != nil {
		log.Printf("auto-population has failed: %v", err)
		return
	}

	newService := service.NewBibleService(db)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{Service: newService}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, nil))
}
