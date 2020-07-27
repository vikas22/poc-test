package migration

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up20200720133320, Down20200720133320)
}

func Up20200720133320(tx *sql.Tx) error {
	_, err := tx.Exec(`CREATE TABLE cards (
  id char(14)  NOT NULL  PRIMARY KEY ,
  merchant_id varchar(14)  NOT NULL,
  name varchar(255)  NOT NULL,
  expiry_month char(2)  NOT NULL,
  expiry_year char(4)  NOT NULL,
  iin char(6)  NOT NULL,
  last4 char(4)  NOT NULL,
  length char(2)  NOT NULL,
  vault varchar(20)  DEFAULT NULL,
  vault_token varchar(50)  DEFAULT NULL,
  global_fingerprint varchar(50)  DEFAULT NULL,
  country char(2)  DEFAULT NULL,
  created_at int NOT NULL,
  updated_at int NOT NULL
);`)

	return err
}

func Down20200720133320(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
