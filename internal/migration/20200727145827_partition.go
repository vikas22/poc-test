package migration

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up20200727145827, Down20200727145827)
}

func Up20200727145827(tx *sql.Tx) error {
	if _, err := tx.Exec(`CREATE TABLE payments_202007 PARTITION OF payments
   FOR VALUES FROM ('2020-07-01') TO ('2020-08-01') PARTITION BY HASH(payment_id) WITH (OIDS=FALSE);`); err != nil {
		return err
	}

  if _, err := tx.Exec(`CREATE TABLE payments_202008 PARTITION OF payments
   FOR VALUES FROM ('2020-08-01') TO ('2020-09-01') PARTITION BY HASH(payment_id) WITH (OIDS=FALSE);`); err != nil {
    return err
  }

  if _, err := tx.Exec(`CREATE TABLE payments_202009 PARTITION OF payments
   FOR VALUES FROM ('2020-9-01') TO ('2020-10-01') PARTITION BY HASH(payment_id) WITH (OIDS=FALSE);`); err != nil {
    return err
  }
	return nil

}

func Down20200727145827(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
