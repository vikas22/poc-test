package card

import "db-poc/internal/common"

// Authentication - Store all the Authentication data for card
type Card struct {
  common.Model
  Id                 string
  MerchantId        string
  Name               string
  ExpiryMonth       string
  ExpiryYear        string
  Iin                string
  Last4              string
  Length             string
  Vault              string
  VaultToken        string
  GlobalFingerprint string
  Country            string
}

// TableName - Get the authentication table name (impl gorm tabler interface)
func (*Card) TableName() string {
  return "cards"
}

