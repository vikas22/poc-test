package card

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


func (c *Core) CreateCard(card Card) error{
  if err := c.repo.Create(&card); err != nil {
    fmt.Println("Error in creating card:", err)
  }
  return nil
}
func (c *Core) Create(vaultToken, merchantId string) error {
  card := Card{Id: utils.NewID(), VaultToken:vaultToken, MerchantId:merchantId}
  fmt.Println(card)
  if err := c.repo.Create(&card); err != nil {
    fmt.Println("Error in creating card:", err)
  }
  return nil
}


