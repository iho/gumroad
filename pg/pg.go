package pg

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // required
)

// Repository is the application's data layer functionality.
type Repository interface {
	CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error)
	GetProduct(ctx context.Context, id int32) (GetProductRow, error)
	GetProducts(ctx context.Context) ([]GetProductsRow, error)
	GetUser(ctx context.Context, id int32) (GetUserRow, error)
	PublishProduct(ctx context.Context, id int32) (Product, error)
}

type repoSvc struct {
	*Queries
	db *sql.DB
}

func (r *repoSvc) withTx(ctx context.Context, txFn func(*Queries) error) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = txFn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			err = fmt.Errorf("tx failed: %v, unable to rollback: %v", err, rbErr)
		}
	} else {
		err = tx.Commit()
	}
	return err
}

func (r *repoSvc) CreateProduct(ctx context.Context, productArg CreateProductParams) (Product, error) {
	product := new(Product)
	err := r.withTx(ctx, func(q *Queries) error {
		res, err := q.CreateProduct(ctx, productArg)
		if err != nil {
			return err
		}
		// for _, authorID := range authorIDs {
		// 	if err := q.SetBookAuthor(ctx, SetBookAuthorParams{
		// 		BookID:   res.ID,
		// 		AuthorID: authorID,
		// 	}); err != nil {
		// 		return err
		// 	}
		// }
		product = &res
		return nil
	})
	return *product, err
}

// NewRepository is used
func NewRepository(db *sql.DB) Repository {
	return &repoSvc{
		Queries: New(db),
		db:      db,
	}
}

// Open opens a database specified by the data source name.
// Format: host=foo port=5432 user=bar password=baz dbname=qux sslmode=disable"
func Open(dataSourceName string) (*sql.DB, error) {
	return sql.Open("postgres", dataSourceName)
}
