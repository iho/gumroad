package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/gorilla/websocket"
	"github.com/iho/gumroad/auth"
	"github.com/iho/gumroad/dataloaders"
	"github.com/iho/gumroad/graph"
	"github.com/iho/gumroad/graph/generated"
	"github.com/iho/gumroad/handlers"
	"github.com/iho/gumroad/pg"
	"github.com/joho/godotenv"
	"github.com/vektah/gqlparser/gqlerror"

	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func Open(dataSourceName string) (*sql.DB, error) {
	return sql.Open("postgres", dataSourceName)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
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
		AllowedOrigins:   []string{"localhost:3000", "http://localhost:8080", "*"},
		AllowCredentials: true,
		Debug:            true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "POST", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
	}).Handler)

	// initialize the repository
	repo := pg.New(db)
	router.Use(jwtauth.Verifier(auth.TokenAuth))
	// router.Use(jwtauth.Authenticator)
	router.Use(auth.Middleware(db, repo))
	router.Use(dataloaders.Middleware(repo))
	config := generated.Config{Resolvers: &graph.Resolver{
		Repository: repo,
	}}
	config.Directives.Authorized = func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
		_, ok := auth.ForContext(ctx)
		if ok {
			return next(ctx)
		}
		return "nil", gqlerror.Errorf("user is not logined :(")
	}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(config))

	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// Check against your desired domains here
				return r.Host == "dream.market"
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})

	router.Handle("/", playground.Handler("Starwars", "/query"))
	router.HandleFunc("/upload_file", handlers.UploadFile)
	router.HandleFunc("/upload_image", handlers.UploadImage)
	router.Handle("/query", srv)

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
