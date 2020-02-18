// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
package graph

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/iho/gumroad/graph/generated"
	"github.com/iho/gumroad/graph/model"
	"github.com/iho/gumroad/pg"
)

func (r *mutationResolver) BuyProduct(ctx context.Context, input *model.BuyProduct) (*model.PayResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateProduct(ctx context.Context, input model.NewProduct) (*pg.Product, error) {
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
		UserID:       sql.NullInt32{Int32: 1},
	})
	fmt.Println(input)
	fmt.Println(product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *mutationResolver) PublishProduct(ctx context.Context, input model.PublishProduct) (*pg.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *productResolver) User(ctx context.Context, obj *pg.Product) (*pg.User, error) {
	user, err := r.Repository.GetUser(ctx, obj.UserID.Int32)
	fmt.Println("User")
	fmt.Println(obj)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *productResolver) Price(ctx context.Context, obj *pg.Product) (int32, error) {
	return obj.Price.Int32, nil
}

func (r *queryResolver) Product(ctx context.Context, id int32) (*pg.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Products(ctx context.Context, userID *int32, count *int32, after *int32) ([]pg.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Me(ctx context.Context) (*model.ExtendedUser, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }
func (r *Resolver) Product() generated.ProductResolver   { return &productResolver{r} }
func (r *Resolver) Query() generated.QueryResolver       { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type productResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.

// func (r *productResolver) Description(ctx context.Context, obj *pg.Product) (*string, error) {
// 	return &obj.Description, nil
// }
// func (r *productResolver) Summary(ctx context.Context, obj *pg.Product) (*string, error) {
// 	return &obj.Summary, nil
// }
// func (r *productResolver) CallToAction(ctx context.Context, obj *pg.Product) (*string, error) {
// 	return &obj.Calltoaction, nil
// }
// func (r *productResolver) CoverImage(ctx context.Context, obj *pg.Product) (*string, error) {
// 	return &obj.Coverimage, nil
// }
// func (r *productResolver) Slug(ctx context.Context, obj *pg.Product) (*string, error) {
// 	return &obj.Slug, nil
// }
// func (r *productResolver) Receipt(ctx context.Context, obj *pg.Product) (*string, error) {
// 	return &obj.Receipt, nil
// }
// func (r *productResolver) Content(ctx context.Context, obj *pg.Product) (*string, error) {
// 	return &obj.Content, nil
// }
