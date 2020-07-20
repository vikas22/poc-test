package payment

import (
  "db-poc/internal/bootstrap"
  "db-poc/internal/common"
)

var (
	repo *Repo
)

// Repo - Repository for authentication
type Repo struct {
	common.Repo
}

// GetRepo - gets singleton instance of the same.
func GetRepo() *Repo {
	if repo == nil {
		repo = &Repo{common.Repo{
			Db: bootstrap.Db.Instance(),
		}}
	}

	return repo
}
