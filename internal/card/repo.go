package card

import (
  "db-poc/internal/bootstrap"
  "db-poc/internal/common"
  "fmt"
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

func (r *Repo)FindExistingCard(vaultToken, merchantId string, m interface{}) error{
  q := r.Db.Model(m).Where(map[string]interface{}{
    "merchant_id":  merchantId,
    "vault_token": vaultToken,
  }).First(m)

  if q.RecordNotFound() {
    return fmt.Errorf("Record Not Found")
  }

  return nil
}
