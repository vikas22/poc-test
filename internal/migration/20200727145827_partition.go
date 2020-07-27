package migration

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up20200727145827, Down20200727145827)
}

func Up20200727145827(tx *sql.Tx) error {
	if _, err := tx.Exec(`CREATE TABLE payment_y2020m07 PARTITION OF payments
    FOR VALUES FROM ('2020-07-01') TO ('2020-08-01');`); err != nil {
		return err
	}

	if _, err := tx.Exec(`CREATE TABLE payment_y2020m08 PARTITION OF payments
    FOR VALUES FROM ('2020-08-01') TO ('2020-09-01');`); err != nil {
		return err
	}

	if _, err := tx.Exec(`CREATE TABLE payment_y2020m09 PARTITION OF payments
    FOR VALUES FROM ('2020-09-01') TO ('2020-10-01');`); err != nil {
		return err
	}

	return nil

}

func Down20200727145827(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
