// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
package graph

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/iho/gumroad/auth"
	"github.com/iho/gumroad/dataloaders"
	"github.com/iho/gumroad/graph/generated"
	"github.com/iho/gumroad/graph/model"
	"github.com/iho/gumroad/pg"
	"github.com/vektah/gqlparser/gqlerror"
)

func (r *fileResolver) Path(ctx context.Context, obj *pg.File) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *fileResolver) IsBought(ctx context.Context, obj *pg.File) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *imageResolver) Path(ctx context.Context, obj *pg.Image) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *imageResolver) PreviewPath(ctx context.Context, obj *pg.Image) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) BuyProduct(ctx context.Context, input *model.BuyProduct) (*model.PayResponse, error) {
	panic("not implemented")
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
	// i, _ := strconv.Atoi(input.ProductID)
	product, err := r.Repository.PublishProduct(ctx, int32(input.ProductID))
	return &product, err
}

func (r *mutationResolver) Signup(ctx context.Context, email string, password string, username string, name *string) (*model.TokenResponse, error) {
	safeName := ""
	if name != nil {
		safeName = *name
	}
	hash, _ := auth.GenerateFromPassword(password)
	user, err := r.Repository.CreateUser(ctx, pg.CreateUserParams{
		Email:    sql.NullString{String: email, Valid: true},
		Password: sql.NullString{String: hash, Valid: true},
		Username: username,
		Name:     safeName,
	})
	if err == nil {
		_, tokenString, _ := auth.TokenAuth.Encode(jwt.MapClaims{"user_id": string(user.ID)})
		return &model.TokenResponse{Token: tokenString}, nil
	}
	return &model.TokenResponse{}, gqlerror.Errorf("Some error occurred")
}

func (r *mutationResolver) Login(ctx context.Context, email string, password string) (*model.TokenResponse, error) {
	userRow, err := r.Repository.GetUserByEmail(ctx, sql.NullString{String: email, Valid: true})
	if err == nil {
		match, err := auth.ComparePasswordAndHash(password, userRow.Password.String)
		if match && err == nil {
			_, tokenString, _ := auth.TokenAuth.Encode(jwt.MapClaims{"user_id": strconv.FormatInt(int64(userRow.ID), 10)})
			return &model.TokenResponse{Token: tokenString}, nil
		}
	}
	return &model.TokenResponse{}, gqlerror.Errorf("No user with current credentials.")
}

func (r *mutationResolver) ForgotPassword(ctx context.Context, email *string) (*model.BoolResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ChangePassword(ctx context.Context, password string) (*model.BoolResponse, error) {
	hash, _ := auth.GenerateFromPassword(password)
	user, ok := auth.ForContext(ctx)
	err := r.Repository.UpdatePassword(ctx, pg.UpdatePasswordParams{
		Password: sql.NullString{String: hash, Valid: true},
		ID:       user.ID,
	})
	if err != nil {
		return nil, err
	}
	return &model.BoolResponse{IsSuccess: ok}, nil
}

func (r *productResolver) ID(ctx context.Context, obj *pg.Product) (int, error) {
	return int(obj.ID), nil
}

func (r *productResolver) User(ctx context.Context, obj *pg.Product) (*pg.User, error) {
	return dataloaders.For(ctx).UserByID.Load(obj.UserID.Int32)
}

func (r *productResolver) Price(ctx context.Context, obj *pg.Product) (int32, error) {
	return obj.Price.Int32, nil
}

func (r *productResolver) Content(ctx context.Context, obj *pg.Product) (*pg.File, error) {
	res, err := r.Repository.GetProductFile(ctx, sql.NullInt32{Int32: obj.ID, Valid: true})
	return &res, err
}

func (r *productResolver) Images(ctx context.Context, obj *pg.Product) ([]pg.Image, error) {
	return r.Repository.GetProductImages(ctx, sql.NullInt32{Int32: obj.ID, Valid: true})
}

func (r *queryResolver) Product(ctx context.Context, username string, slug string) (*pg.Product, error) {
	product, err := r.Repository.GetProduct(ctx, pg.GetProductParams{
		Slug:     slug,
		Username: username,
	})
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *queryResolver) Products(ctx context.Context, username *string, count *int32, after *int32) ([]pg.Product, error) {
	var afterID int32 = 0
	if after != nil {
		afterID = *after
	}

	var limit int32 = 100
	if count != nil {
		limit = *count
	}
	var products []pg.Product
	var err error
	if username == nil {
		products, err = r.Repository.GetProducts(ctx, pg.GetProductsParams{
			ID:    afterID,
			Limit: limit,
		})

	} else {
		products, err = r.Repository.GetUserProducts(ctx, pg.GetUserProductsParams{
			Username: *username,
			ID:       afterID,
			Limit:    limit,
		})
	}

	return products, err
}

func (r *queryResolver) MyProducts(ctx context.Context, count *int32, after *int32) ([]pg.Product, error) {
	var afterID int32 = 0
	if after != nil {
		afterID = *after
	}

	var limit int32 = 100
	if count != nil {
		limit = *count
	}

	user, _ := auth.ForContext(ctx)

	return r.Repository.MyProducts(ctx, pg.MyProductsParams{
		ID:     afterID,
		Limit:  limit,
		UserID: sql.NullInt32{Int32: user.ID, Valid: true},
	})
}

func (r *queryResolver) Me(ctx context.Context) (*model.ExtendedUser, error) {
	user, _ := auth.ForContext(ctx)
	return converUserToExtendedUser(user), nil
}

func (r *queryResolver) MyImages(ctx context.Context) ([]pg.Image, error) {
	user, _ := auth.ForContext(ctx)
	return r.Repository.GetMyImages(ctx, sql.NullInt32{Int32: user.ID, Valid: true})
}

func (r *queryResolver) MyFiles(ctx context.Context) ([]pg.File, error) {
	user, _ := auth.ForContext(ctx)
	return r.Repository.GetMyFiles(ctx, sql.NullInt32{Int32: user.ID, Valid: true})
}

func (r *userResolver) ID(ctx context.Context, obj *pg.User) (int, error) {
	return int(obj.ID), nil
}

func (r *Resolver) File() generated.FileResolver         { return &fileResolver{r} }
func (r *Resolver) Image() generated.ImageResolver       { return &imageResolver{r} }
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }
func (r *Resolver) Product() generated.ProductResolver   { return &productResolver{r} }
func (r *Resolver) Query() generated.QueryResolver       { return &queryResolver{r} }
func (r *Resolver) User() generated.UserResolver         { return &userResolver{r} }

type fileResolver struct{ *Resolver }
type imageResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type productResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
