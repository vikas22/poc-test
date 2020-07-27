package payment

import (
  "db-poc/internal/common"
  "db-poc/pkg/utils"
  "fmt"
  "sync"
)

var (
	core *Core
	once sync.Once
)

// Core have methods implementing buisness usecase
type Core struct {
	common.Core
	repo *Repo
}

// GetCore - returns the core instance
func GetCore() (*Core, error) {
	once.Do(func() {
		core = &Core{
			repo: GetRepo(),
		}
	})

	return core, nil
}

// ResetRepo resets authentication repo, used in testing
func ResetRepo() {
	repo = nil
	if core != nil {
		core.repo = GetRepo()
	}
}


func (c *Core) CreatePayment(payment Payment) error {
  if err := c.repo.Create(&payment); err != nil {
    fmt.Println("Error in creating payments:", err)
  }
  return nil
}

func (c *Core) Create(cardId string) error {
  payment := Payment{CardId:cardId, PaymentId: utils.NewID()}
  fmt.Println(payment)
  if err := c.repo.Create(&payment); err != nil {
    fmt.Println("Error in creating payments:", err)
  }
  return nil
}
