package auth

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/jwtauth"
	"github.com/iho/gumroad/pg"
)

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

// A stand-in for our database backed user object

// type User struct {
// 	ID       int32
// 	Username string
// 	Name     string
// }

// Middleware decodes the share session cookie and packs the session into context
func Middleware(db *sql.DB, repo pg.Querier) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			_, claims, err := jwtauth.FromContext(r.Context())
			fmt.Println(err)
			fmt.Println(claims)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			userIDstring, _ := claims["user_id"]

			userID := userIDstring.(string)
			id, err := strconv.Atoi(userID)
			if err != nil {
				id = -1
			}
			user, err := repo.GetUser(r.Context(), int32(id))
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			ctx := context.WithValue(r.Context(), userCtxKey, &user)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) (*pg.User, bool) {
	user, ok := ctx.Value(userCtxKey).(*pg.User)
	return user, ok
}
