package migration

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up20200720171940, Down20200720171940)
}

func Up20200720171940(tx *sql.Tx) error {
  _, err := tx.Exec(`CREATE INDEX VAULT_TOKEN_MERCHANT_ID ON CARDS (merchant_id, vault_token);`)
	return err
}

func Down20200720171940(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
