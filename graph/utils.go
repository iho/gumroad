package graph

import (
	"github.com/iho/gumroad/graph/model"
	"github.com/iho/gumroad/pg"
)

func converUserToExtendedUser(user *pg.User) *model.ExtendedUser {
	return &model.ExtendedUser{
		Balance:  user.Balance.Int32,
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		Bio:      user.Bio,
	}
}
