package migration

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up20200720125828, Down20200720125828)
}

func Up20200720125828(tx *sql.Tx) error {
	_, err := tx.Exec(`CREATE TABLE payments (
    id char(14)  NOT NULL,
    merchant_id char(14)  NOT NULL,
    amount int NOT NULL,
    currency char(3)  NOT NULL,
    base_amount int NOT NULL,
    method varchar(255)  NOT NULL,
    status varchar(255)  NOT NULL,
    two_factor_auth varchar(20)  DEFAULT NULL,
    order_id char(14)  DEFAULT NULL,
    invoice_id char(14)  DEFAULT NULL,
    transfer_id char(14)  DEFAULT NULL,
    payment_link_id char(14)  DEFAULT NULL,
    receiver_id char(14)  DEFAULT NULL,
    receiver_type varchar(255)  DEFAULT NULL,
    international BOOLEAN DEFAULT NULL,
    amount_authorized int NOT NULL DEFAULT 0,
    amount_refunded int NOT NULL DEFAULT 0,
    base_amount_refunded int NOT NULL DEFAULT 0,
    amount_transferred int NOT NULL DEFAULT 0,
    amount_paidout int NOT NULL DEFAULT 0,
    refund_status varchar(255)  DEFAULT NULL,
    description varchar(255)  DEFAULT NULL,
    card_id char(14)  DEFAULT NULL,
    subscription_id char(14)  DEFAULT NULL,
    bank char(6)  DEFAULT NULL,
    wallet varchar(15)  DEFAULT NULL,
    vpa varchar(100)  DEFAULT NULL,
    on_hold BOOLEAN NOT NULL DEFAULT false,
    on_hold_until int DEFAULT NULL,
    emi_plan_id char(14)  DEFAULT NULL,
    emi_subvention varchar(10)  DEFAULT NULL,
    error_code varchar(100)  DEFAULT NULL,
    internal_error_code varchar(255)  DEFAULT NULL,
    error_description varchar(255)  DEFAULT NULL,
    cancellation_reason varchar(255)  DEFAULT NULL,
    customer_id varchar(14)  DEFAULT NULL,
    global_customer_id varchar(14)  DEFAULT NULL,
    app_token varchar(14)  DEFAULT NULL,
    token_id varchar(14)  DEFAULT NULL,
    global_token_id varchar(14)  DEFAULT NULL,
    email varchar(255)  DEFAULT NULL,
    contact varchar(20)  DEFAULT NULL,
    notes text  DEFAULT NULL,
    transaction_id char(14)  DEFAULT NULL,
    authorized_at int DEFAULT NULL,
    auto_captured BOOLEAN NOT NULL DEFAULT false,
    captured_at int DEFAULT NULL,
    gateway varchar(255)  DEFAULT NULL,
    terminal_id char(14)  DEFAULT NULL,
    authentication_gateway varchar(255)  DEFAULT NULL,
    batch_id varchar(255)  DEFAULT NULL,
    reference1 varchar(255)  DEFAULT NULL,
    reference2 varchar(255)  DEFAULT NULL,
    reference3 BOOLEAN DEFAULT false,
    cps_route BOOLEAN DEFAULT false,
    reference5 int DEFAULT NULL,
    reference6 int DEFAULT NULL,
    reference9 int DEFAULT NULL,
    signed BOOLEAN NOT NULL DEFAULT false,
    verified BOOLEAN DEFAULT false,
    gateway_captured BOOLEAN DEFAULT false,
    verify_bucket BOOLEAN DEFAULT false,
    verify_at int DEFAULT NULL,
    callback_url text  DEFAULT NULL,
    fee int DEFAULT NULL,
    mdr int DEFAULT NULL,
    tax int DEFAULT NULL,
    otp_attempts BOOLEAN DEFAULT false,
    otp_count BOOLEAN DEFAULT false,
    recurring BOOLEAN NOT NULL DEFAULT false,
    save BOOLEAN NOT NULL DEFAULT false,
    late_authorized BOOLEAN DEFAULT NULL,
    convert_currency BOOLEAN DEFAULT NULL,
    disputed BOOLEAN NOT NULL DEFAULT false,
    recurring_type varchar(255)  DEFAULT NULL,
    auth_type varchar(14)  DEFAULT NULL,
    acknowledged_at int DEFAULT NULL,
    refund_at int DEFAULT NULL,
    fee_bearer int DEFAULT 0,
    reference13 char(14)  DEFAULT NULL,
    reference14 char(14)  DEFAULT NULL,
    settled_by varchar(255)  DEFAULT NULL,
    reference16 varchar(255)  DEFAULT NULL,
    reference17 text  DEFAULT NULL,
    created_at int NOT NULL,
    partition_at DATE NOT NULL,
    updated_at int,
    CONSTRAINT payments_pkey PRIMARY KEY (payment_id, partition_at))
    PARTITION BY RANGE(partition_at)
    WITH (OIDS=FALSE);`)


	return err
}

func Down20200720125828(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
