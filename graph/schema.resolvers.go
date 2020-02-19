// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
package graph

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/iho/gumroad/auth"
	"github.com/iho/gumroad/graph/generated"
	"github.com/iho/gumroad/graph/model"
	"github.com/iho/gumroad/pg"
	"github.com/vektah/gqlparser/gqlerror"
)

func (r *mutationResolver) BuyProduct(ctx context.Context, input *model.BuyProduct) (*model.PayResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateProduct(ctx context.Context, input model.NewProduct) (*pg.Product, error) {
	user, ok := auth.ForContext(ctx)
	if !ok {
		return nil, gqlerror.Errorf("user is not logined :(")
	}

	product, err := r.Repository.CreateProduct(ctx, pg.CreateProductParams{
		Name:         input.Name,
		Price:        sql.NullInt32{Int32: input.Price, Valid: true},
		Description:  input.Description,
		Summary:      input.Summary,
		Calltoaction: input.CallToAction,
		Coverimage:   input.CoverImage,
		Slug:         input.Slug,
		Receipt:      input.Receipt,
		Content:      input.Content,
		Ispablished:  input.IsPablished,
		UserID:       sql.NullInt32{Int32: user.ID, Valid: true},
	})
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *mutationResolver) PublishProduct(ctx context.Context, input model.PublishProduct) (*pg.Product, error) {
	product, err := r.Repository.PublishProduct(ctx, input.ProductID)
	return &product, err
}

func (r *mutationResolver) Signup(ctx context.Context, email string, password string, username string, name *string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, email string, password string) (string, error) {
	_, tokenString, _ := auth.TokenAuth.Encode(jwt.MapClaims{"user_id": "1"})
	return tokenString, nil
}

func (r *productResolver) User(ctx context.Context, obj *pg.Product) (*pg.User, error) {
	userRow, err := r.Repository.GetUser(ctx, obj.UserID.Int32)
	if err != nil {
		return nil, err
	}
	user := pg.User{ID: userRow.ID, Name: userRow.Name, Username: userRow.Username}
	return &user, nil
}

func (r *productResolver) Price(ctx context.Context, obj *pg.Product) (int32, error) {
	return obj.Price.Int32, nil
}

func (r *queryResolver) Product(ctx context.Context, id int32) (*pg.Product, error) {
	product, err := r.Repository.GetProduct(ctx, id)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *queryResolver) Products(ctx context.Context, userID *int32, count *int32, after *int32) ([]pg.Product, error) {
	userId := sql.NullInt32{Int32: 0, Valid: false}
	if userID != nil {
		userId.Int32 = *userID
	}

	var afterID int32 = 0
	if after != nil {
		afterID = *after
	}

	var limit int32 = 100
	if count != nil {
		limit = *count
	}

	products, err := r.Repository.GetProducts(ctx, pg.GetProductsParams{
		UserID: userId,
		ID:     afterID,
		Limit:  limit,
	})
	return products, err
}

func (r *queryResolver) Me(ctx context.Context) (*model.ExtendedUser, error) {
	user, ok := auth.ForContext(ctx)
	if ok {
		return converUserToExtendedUser(user), nil
	}
	return nil, gqlerror.Errorf("user is not logined :(")

}

func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }
func (r *Resolver) Product() generated.ProductResolver   { return &productResolver{r} }
func (r *Resolver) Query() generated.QueryResolver       { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type productResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
