// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
package graph

import (
	"github.com/iho/gumroad/pg"
)

type Resolver struct {
	Repository pg.Querier
}
