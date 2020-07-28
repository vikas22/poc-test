package card

import (
  "db-poc/internal/bootstrap"
  "db-poc/internal/common"
  "db-poc/pkg/prom_metrics"
  "fmt"
  "time"
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
  now := time.Now()
  q := r.Db.Model(m).Where(map[string]interface{}{
    "merchant_id":  merchantId,
    "vault_token": vaultToken,
  }).First(m)

  if q.RecordNotFound() {
    prom_metrics.DbRequestDuration("card_fetch_not_found", true, now)
    return fmt.Errorf("Record Not Found")
  }

  prom_metrics.DbRequestDuration("card_fetch_found", true, now)

  return nil
}

