package main

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/gorilla/websocket"
	"github.com/iho/gumroad/auth"
	"github.com/iho/gumroad/graph"
	"github.com/iho/gumroad/graph/generated"
	"github.com/iho/gumroad/pg"

	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func Open(dataSourceName string) (*sql.DB, error) {
	return sql.Open("postgres", dataSourceName)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	db, err := Open("user=db_user dbname=db sslmode=disable password=password")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	router := chi.NewRouter()
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
		// Debug:            true,
	}).Handler)

	// initialize the repository
	repo := pg.New(db)
	router.Use(jwtauth.Verifier(auth.TokenAuth))
	// router.Use(jwtauth.Authenticator)
	router.Use(auth.Middleware(db, repo))

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		Repository: repo,
	}}))
	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// Check against your desired domains here
				return r.Host == "localhost"
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})

	router.Handle("/", playground.Handler("Starwars", "/query"))
	router.Handle("/query", srv)

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
