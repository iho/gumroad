// Code generated by sqlc. DO NOT EDIT.
// source: queries.sql

package pg

import (
	"context"
	"database/sql"
)

const createProduct = `-- name: CreateProduct :one
insert into public.products (
    name,
    price,
    description,
    summary,
    callToAction,
    coverImage,
    slug,
    isPablished,
    receipt,
    content,
    user_id
  )
values
  (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10,
    $11
  ) returning id, user_id, price, name, description, summary, calltoaction, coverimage, slug, ispablished, receipt, content, created_at, updated_at
`

type CreateProductParams struct {
	Name         string        `json:"name"`
	Price        sql.NullInt32 `json:"price"`
	Description  string        `json:"description"`
	Summary      string        `json:"summary"`
	Calltoaction string        `json:"calltoaction"`
	Coverimage   string        `json:"coverimage"`
	Slug         string        `json:"slug"`
	Ispablished  bool          `json:"ispablished"`
	Receipt      string        `json:"receipt"`
	Content      string        `json:"content"`
	UserID       sql.NullInt32 `json:"user_id"`
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, createProduct,
		arg.Name,
		arg.Price,
		arg.Description,
		arg.Summary,
		arg.Calltoaction,
		arg.Coverimage,
		arg.Slug,
		arg.Ispablished,
		arg.Receipt,
		arg.Content,
		arg.UserID,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Price,
		&i.Name,
		&i.Description,
		&i.Summary,
		&i.Calltoaction,
		&i.Coverimage,
		&i.Slug,
		&i.Ispablished,
		&i.Receipt,
		&i.Content,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
insert into public.users (username, "name", bio)
values
  ($1, $2, $3) returning id, username, name, bio, created_at, updated_at, last_active_at
`

type CreateUserParams struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Bio      string `json:"bio"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Username, arg.Name, arg.Bio)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Name,
		&i.Bio,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.LastActiveAt,
	)
	return i, err
}

const getProduct = `-- name: GetProduct :one
select
  id, user_id, price, name, description, summary, calltoaction, coverimage, slug, ispablished, receipt, content, created_at, updated_at
from public.products
where
  id = $1
`

func (q *Queries) GetProduct(ctx context.Context, id int32) (Product, error) {
	row := q.db.QueryRowContext(ctx, getProduct, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Price,
		&i.Name,
		&i.Description,
		&i.Summary,
		&i.Calltoaction,
		&i.Coverimage,
		&i.Slug,
		&i.Ispablished,
		&i.Receipt,
		&i.Content,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getProducts = `-- name: GetProducts :many
select
  id, user_id, price, name, description, summary, calltoaction, coverimage, slug, ispablished, receipt, content, created_at, updated_at
from public.products
where
  user_id = $1 or $1 is null
  and id > $2 limit $3
`

type GetProductsParams struct {
	UserID sql.NullInt32 `json:"user_id"`
	ID     int32         `json:"id"`
	Limit  int32         `json:"limit"`
}

func (q *Queries) GetProducts(ctx context.Context, arg GetProductsParams) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, getProducts, arg.UserID, arg.ID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Price,
			&i.Name,
			&i.Description,
			&i.Summary,
			&i.Calltoaction,
			&i.Coverimage,
			&i.Slug,
			&i.Ispablished,
			&i.Receipt,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUser = `-- name: GetUser :one
select
  id, username, name, bio, created_at, updated_at, last_active_at
from public.users
where
  id = $1
`

func (q *Queries) GetUser(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Name,
		&i.Bio,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.LastActiveAt,
	)
	return i, err
}

const publishProduct = `-- name: PublishProduct :one
update public.products
set
  isPablished = true
where
  id = $1 returning id, user_id, price, name, description, summary, calltoaction, coverimage, slug, ispablished, receipt, content, created_at, updated_at
`

func (q *Queries) PublishProduct(ctx context.Context, id int32) (Product, error) {
	row := q.db.QueryRowContext(ctx, publishProduct, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Price,
		&i.Name,
		&i.Description,
		&i.Summary,
		&i.Calltoaction,
		&i.Coverimage,
		&i.Slug,
		&i.Ispablished,
		&i.Receipt,
		&i.Content,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
