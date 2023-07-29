package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/DevPio/gographql/graph"
	"github.com/DevPio/gographql/internal/database"
	_ "github.com/mattn/go-sqlite3"
)

const defaultPort = "8080"

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	fmt.Println(err)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	categoryDb := database.NewCategory(db)
	courseDb := database.NewCourse(db)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{CategoryDb: categoryDb, CourseDb: courseDb}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
