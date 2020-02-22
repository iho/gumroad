package main

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

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
	"github.com/iho/gumroad/pg"
	"github.com/vektah/gqlparser/gqlerror"

	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func Open(dataSourceName string) (*sql.DB, error) {
	return sql.Open("postgres", dataSourceName)
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Uploading File")
	if err := r.ParseMultipartForm(100 << 20); err != nil {
		fmt.Println(err)
	}
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	saveName := path.Join("/tmp", "ihor", path.Base(handler.Filename))
	savef, err := os.Create(saveName)
	if err != nil {
		// Failed to create file on server, handle err
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}
	defer savef.Close()

	if _, err := io.Copy(savef, file); err != nil {
		fmt.Println("Error")
	}
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
	router.HandleFunc("/upload", uploadFile)
	router.Handle("/query", srv)

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
