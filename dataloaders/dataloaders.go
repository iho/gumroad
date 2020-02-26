package dataloaders

import (
	"context"
	"log"
	"time"

	"github.com/iho/gumroad/pg"
)

//go:generate go run github.com/vektah/dataloaden UserLoader int32 *github.com/iho/gumroad/pg.User
//go:generate go run github.com/vektah/dataloaden ProductLoader int32 *github.com/iho/gumroad/pg.Product

type Loaders struct {
	UserByID    *UserLoader
	ProductByID *ProductLoader
}

type contextKey string

const key = contextKey("dataloaders")

func newLoaders(ctx context.Context, repo pg.Querier) *Loaders {
	return &Loaders{
		UserByID:    newUserByID(ctx, repo),
		ProductByID: newProductByID(ctx, repo),
	}
}

func newUserByID(ctx context.Context, repo pg.Querier) *UserLoader {
	return &UserLoader{
		maxBatch: 100,
		wait:     1 * time.Millisecond,
		fetch: func(userIDs []int32) ([]*pg.User, []error) {
			// db query
			res, err := repo.ListUsers(ctx, userIDs)
			if err != nil {
				return nil, []error{err}
			}
			// map
			groupByUserIDs := make(map[int32]*pg.User, len(userIDs))
			for _, r := range res {
				groupByUserIDs[r.ID] = &r
			}
			// order
			result := make([]*pg.User, len(userIDs))
			for i, ID := range userIDs {
				result[i] = groupByUserIDs[ID]
			}
			log.Println("result", result)
			return result, nil
		},
	}
}

func newProductByID(ctx context.Context, repo pg.Querier) *ProductLoader {
	return &ProductLoader{
		maxBatch: 100,
		wait:     1 * time.Millisecond,
		fetch: func(productIDs []int32) ([]*pg.Product, []error) {
			// db query
			res, err := repo.ListProducts(ctx, productIDs)
			if err != nil {
				return nil, []error{err}
			}
			// map
			groupByProductIDs := make(map[int32]*pg.Product, len(productIDs))
			for _, r := range res {
				groupByProductIDs[r.ID] = &r
			}
			// order
			result := make([]*pg.Product, len(productIDs))
			for i, ID := range productIDs {
				result[i] = groupByProductIDs[ID]
			}
			return result, nil
		},
	}
}
func For(ctx context.Context) *Loaders {
	return ctx.Value(key).(*Loaders)
}
