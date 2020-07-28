package migration

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up20200728162603, Down20200728162603)
}

func Up20200728162603(tx *sql.Tx) error {
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_0 partition of payments_202007 for values with (modulus 180, remainder 0) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_1 partition of payments_202007 for values with (modulus 180, remainder 1) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_2 partition of payments_202007 for values with (modulus 180, remainder 2) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_3 partition of payments_202007 for values with (modulus 180, remainder 3) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_4 partition of payments_202007 for values with (modulus 180, remainder 4) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_5 partition of payments_202007 for values with (modulus 180, remainder 5) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_6 partition of payments_202007 for values with (modulus 180, remainder 6) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_7 partition of payments_202007 for values with (modulus 180, remainder 7) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_8 partition of payments_202007 for values with (modulus 180, remainder 8) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_9 partition of payments_202007 for values with (modulus 180, remainder 9) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_10 partition of payments_202007 for values with (modulus 180, remainder 10) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_11 partition of payments_202007 for values with (modulus 180, remainder 11) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_12 partition of payments_202007 for values with (modulus 180, remainder 12) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_13 partition of payments_202007 for values with (modulus 180, remainder 13) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_14 partition of payments_202007 for values with (modulus 180, remainder 14) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_15 partition of payments_202007 for values with (modulus 180, remainder 15) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_16 partition of payments_202007 for values with (modulus 180, remainder 16) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_17 partition of payments_202007 for values with (modulus 180, remainder 17) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_18 partition of payments_202007 for values with (modulus 180, remainder 18) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_19 partition of payments_202007 for values with (modulus 180, remainder 19) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_20 partition of payments_202007 for values with (modulus 180, remainder 20) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_21 partition of payments_202007 for values with (modulus 180, remainder 21) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_22 partition of payments_202007 for values with (modulus 180, remainder 22) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_23 partition of payments_202007 for values with (modulus 180, remainder 23) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_24 partition of payments_202007 for values with (modulus 180, remainder 24) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_25 partition of payments_202007 for values with (modulus 180, remainder 25) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_26 partition of payments_202007 for values with (modulus 180, remainder 26) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_27 partition of payments_202007 for values with (modulus 180, remainder 27) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_28 partition of payments_202007 for values with (modulus 180, remainder 28) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_29 partition of payments_202007 for values with (modulus 180, remainder 29) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_30 partition of payments_202007 for values with (modulus 180, remainder 30) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_31 partition of payments_202007 for values with (modulus 180, remainder 31) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_32 partition of payments_202007 for values with (modulus 180, remainder 32) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_33 partition of payments_202007 for values with (modulus 180, remainder 33) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_34 partition of payments_202007 for values with (modulus 180, remainder 34) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_35 partition of payments_202007 for values with (modulus 180, remainder 35) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_36 partition of payments_202007 for values with (modulus 180, remainder 36) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_37 partition of payments_202007 for values with (modulus 180, remainder 37) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_38 partition of payments_202007 for values with (modulus 180, remainder 38) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_39 partition of payments_202007 for values with (modulus 180, remainder 39) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_40 partition of payments_202007 for values with (modulus 180, remainder 40) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_41 partition of payments_202007 for values with (modulus 180, remainder 41) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_42 partition of payments_202007 for values with (modulus 180, remainder 42) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_43 partition of payments_202007 for values with (modulus 180, remainder 43) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_44 partition of payments_202007 for values with (modulus 180, remainder 44) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_45 partition of payments_202007 for values with (modulus 180, remainder 45) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_46 partition of payments_202007 for values with (modulus 180, remainder 46) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_47 partition of payments_202007 for values with (modulus 180, remainder 47) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_48 partition of payments_202007 for values with (modulus 180, remainder 48) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_49 partition of payments_202007 for values with (modulus 180, remainder 49) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_50 partition of payments_202007 for values with (modulus 180, remainder 50) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_51 partition of payments_202007 for values with (modulus 180, remainder 51) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_52 partition of payments_202007 for values with (modulus 180, remainder 52) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_53 partition of payments_202007 for values with (modulus 180, remainder 53) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_54 partition of payments_202007 for values with (modulus 180, remainder 54) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_55 partition of payments_202007 for values with (modulus 180, remainder 55) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_56 partition of payments_202007 for values with (modulus 180, remainder 56) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_57 partition of payments_202007 for values with (modulus 180, remainder 57) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_58 partition of payments_202007 for values with (modulus 180, remainder 58) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_59 partition of payments_202007 for values with (modulus 180, remainder 59) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_60 partition of payments_202007 for values with (modulus 180, remainder 60) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_61 partition of payments_202007 for values with (modulus 180, remainder 61) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_62 partition of payments_202007 for values with (modulus 180, remainder 62) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_63 partition of payments_202007 for values with (modulus 180, remainder 63) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_64 partition of payments_202007 for values with (modulus 180, remainder 64) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_65 partition of payments_202007 for values with (modulus 180, remainder 65) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_66 partition of payments_202007 for values with (modulus 180, remainder 66) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_67 partition of payments_202007 for values with (modulus 180, remainder 67) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_68 partition of payments_202007 for values with (modulus 180, remainder 68) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_69 partition of payments_202007 for values with (modulus 180, remainder 69) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_70 partition of payments_202007 for values with (modulus 180, remainder 70) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_71 partition of payments_202007 for values with (modulus 180, remainder 71) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_72 partition of payments_202007 for values with (modulus 180, remainder 72) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_73 partition of payments_202007 for values with (modulus 180, remainder 73) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_74 partition of payments_202007 for values with (modulus 180, remainder 74) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_75 partition of payments_202007 for values with (modulus 180, remainder 75) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_76 partition of payments_202007 for values with (modulus 180, remainder 76) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_77 partition of payments_202007 for values with (modulus 180, remainder 77) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_78 partition of payments_202007 for values with (modulus 180, remainder 78) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_79 partition of payments_202007 for values with (modulus 180, remainder 79) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_80 partition of payments_202007 for values with (modulus 180, remainder 80) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_81 partition of payments_202007 for values with (modulus 180, remainder 81) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_82 partition of payments_202007 for values with (modulus 180, remainder 82) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_83 partition of payments_202007 for values with (modulus 180, remainder 83) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_84 partition of payments_202007 for values with (modulus 180, remainder 84) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_85 partition of payments_202007 for values with (modulus 180, remainder 85) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_86 partition of payments_202007 for values with (modulus 180, remainder 86) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_87 partition of payments_202007 for values with (modulus 180, remainder 87) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_88 partition of payments_202007 for values with (modulus 180, remainder 88) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_89 partition of payments_202007 for values with (modulus 180, remainder 89) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_90 partition of payments_202007 for values with (modulus 180, remainder 90) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_91 partition of payments_202007 for values with (modulus 180, remainder 91) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_92 partition of payments_202007 for values with (modulus 180, remainder 92) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_93 partition of payments_202007 for values with (modulus 180, remainder 93) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_94 partition of payments_202007 for values with (modulus 180, remainder 94) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_95 partition of payments_202007 for values with (modulus 180, remainder 95) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_96 partition of payments_202007 for values with (modulus 180, remainder 96) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_97 partition of payments_202007 for values with (modulus 180, remainder 97) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_98 partition of payments_202007 for values with (modulus 180, remainder 98) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_99 partition of payments_202007 for values with (modulus 180, remainder 99) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_100 partition of payments_202007 for values with (modulus 180, remainder 100) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_101 partition of payments_202007 for values with (modulus 180, remainder 101) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_102 partition of payments_202007 for values with (modulus 180, remainder 102) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_103 partition of payments_202007 for values with (modulus 180, remainder 103) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_104 partition of payments_202007 for values with (modulus 180, remainder 104) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_105 partition of payments_202007 for values with (modulus 180, remainder 105) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_106 partition of payments_202007 for values with (modulus 180, remainder 106) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_107 partition of payments_202007 for values with (modulus 180, remainder 107) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_108 partition of payments_202007 for values with (modulus 180, remainder 108) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_109 partition of payments_202007 for values with (modulus 180, remainder 109) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_110 partition of payments_202007 for values with (modulus 180, remainder 110) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_111 partition of payments_202007 for values with (modulus 180, remainder 111) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_112 partition of payments_202007 for values with (modulus 180, remainder 112) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_113 partition of payments_202007 for values with (modulus 180, remainder 113) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_114 partition of payments_202007 for values with (modulus 180, remainder 114) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_115 partition of payments_202007 for values with (modulus 180, remainder 115) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_116 partition of payments_202007 for values with (modulus 180, remainder 116) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_117 partition of payments_202007 for values with (modulus 180, remainder 117) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_118 partition of payments_202007 for values with (modulus 180, remainder 118) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_119 partition of payments_202007 for values with (modulus 180, remainder 119) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_120 partition of payments_202007 for values with (modulus 180, remainder 120) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_121 partition of payments_202007 for values with (modulus 180, remainder 121) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_122 partition of payments_202007 for values with (modulus 180, remainder 122) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_123 partition of payments_202007 for values with (modulus 180, remainder 123) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_124 partition of payments_202007 for values with (modulus 180, remainder 124) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_125 partition of payments_202007 for values with (modulus 180, remainder 125) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_126 partition of payments_202007 for values with (modulus 180, remainder 126) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_127 partition of payments_202007 for values with (modulus 180, remainder 127) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_128 partition of payments_202007 for values with (modulus 180, remainder 128) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_129 partition of payments_202007 for values with (modulus 180, remainder 129) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_130 partition of payments_202007 for values with (modulus 180, remainder 130) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_131 partition of payments_202007 for values with (modulus 180, remainder 131) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_132 partition of payments_202007 for values with (modulus 180, remainder 132) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_133 partition of payments_202007 for values with (modulus 180, remainder 133) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_134 partition of payments_202007 for values with (modulus 180, remainder 134) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_135 partition of payments_202007 for values with (modulus 180, remainder 135) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_136 partition of payments_202007 for values with (modulus 180, remainder 136) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_137 partition of payments_202007 for values with (modulus 180, remainder 137) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_138 partition of payments_202007 for values with (modulus 180, remainder 138) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_139 partition of payments_202007 for values with (modulus 180, remainder 139) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_140 partition of payments_202007 for values with (modulus 180, remainder 140) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_141 partition of payments_202007 for values with (modulus 180, remainder 141) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_142 partition of payments_202007 for values with (modulus 180, remainder 142) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_143 partition of payments_202007 for values with (modulus 180, remainder 143) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_144 partition of payments_202007 for values with (modulus 180, remainder 144) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_145 partition of payments_202007 for values with (modulus 180, remainder 145) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_146 partition of payments_202007 for values with (modulus 180, remainder 146) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_147 partition of payments_202007 for values with (modulus 180, remainder 147) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_148 partition of payments_202007 for values with (modulus 180, remainder 148) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_149 partition of payments_202007 for values with (modulus 180, remainder 149) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_150 partition of payments_202007 for values with (modulus 180, remainder 150) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_151 partition of payments_202007 for values with (modulus 180, remainder 151) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_152 partition of payments_202007 for values with (modulus 180, remainder 152) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_153 partition of payments_202007 for values with (modulus 180, remainder 153) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_154 partition of payments_202007 for values with (modulus 180, remainder 154) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_155 partition of payments_202007 for values with (modulus 180, remainder 155) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_156 partition of payments_202007 for values with (modulus 180, remainder 156) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_157 partition of payments_202007 for values with (modulus 180, remainder 157) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_158 partition of payments_202007 for values with (modulus 180, remainder 158) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_159 partition of payments_202007 for values with (modulus 180, remainder 159) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_160 partition of payments_202007 for values with (modulus 180, remainder 160) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_161 partition of payments_202007 for values with (modulus 180, remainder 161) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_162 partition of payments_202007 for values with (modulus 180, remainder 162) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_163 partition of payments_202007 for values with (modulus 180, remainder 163) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_164 partition of payments_202007 for values with (modulus 180, remainder 164) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_165 partition of payments_202007 for values with (modulus 180, remainder 165) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_166 partition of payments_202007 for values with (modulus 180, remainder 166) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_167 partition of payments_202007 for values with (modulus 180, remainder 167) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_168 partition of payments_202007 for values with (modulus 180, remainder 168) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_169 partition of payments_202007 for values with (modulus 180, remainder 169) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_170 partition of payments_202007 for values with (modulus 180, remainder 170) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_171 partition of payments_202007 for values with (modulus 180, remainder 171) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_172 partition of payments_202007 for values with (modulus 180, remainder 172) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_173 partition of payments_202007 for values with (modulus 180, remainder 173) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_174 partition of payments_202007 for values with (modulus 180, remainder 174) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_175 partition of payments_202007 for values with (modulus 180, remainder 175) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_176 partition of payments_202007 for values with (modulus 180, remainder 176) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_177 partition of payments_202007 for values with (modulus 180, remainder 177) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_178 partition of payments_202007 for values with (modulus 180, remainder 178) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202007_hash_179 partition of payments_202007 for values with (modulus 180, remainder 179) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_0 partition of payments_202008 for values with (modulus 180, remainder 0) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_1 partition of payments_202008 for values with (modulus 180, remainder 1) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_2 partition of payments_202008 for values with (modulus 180, remainder 2) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_3 partition of payments_202008 for values with (modulus 180, remainder 3) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_4 partition of payments_202008 for values with (modulus 180, remainder 4) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_5 partition of payments_202008 for values with (modulus 180, remainder 5) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_6 partition of payments_202008 for values with (modulus 180, remainder 6) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_7 partition of payments_202008 for values with (modulus 180, remainder 7) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_8 partition of payments_202008 for values with (modulus 180, remainder 8) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_9 partition of payments_202008 for values with (modulus 180, remainder 9) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_10 partition of payments_202008 for values with (modulus 180, remainder 10) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_11 partition of payments_202008 for values with (modulus 180, remainder 11) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_12 partition of payments_202008 for values with (modulus 180, remainder 12) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_13 partition of payments_202008 for values with (modulus 180, remainder 13) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_14 partition of payments_202008 for values with (modulus 180, remainder 14) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_15 partition of payments_202008 for values with (modulus 180, remainder 15) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_16 partition of payments_202008 for values with (modulus 180, remainder 16) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_17 partition of payments_202008 for values with (modulus 180, remainder 17) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_18 partition of payments_202008 for values with (modulus 180, remainder 18) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_19 partition of payments_202008 for values with (modulus 180, remainder 19) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_20 partition of payments_202008 for values with (modulus 180, remainder 20) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_21 partition of payments_202008 for values with (modulus 180, remainder 21) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_22 partition of payments_202008 for values with (modulus 180, remainder 22) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_23 partition of payments_202008 for values with (modulus 180, remainder 23) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_24 partition of payments_202008 for values with (modulus 180, remainder 24) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_25 partition of payments_202008 for values with (modulus 180, remainder 25) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_26 partition of payments_202008 for values with (modulus 180, remainder 26) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_27 partition of payments_202008 for values with (modulus 180, remainder 27) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_28 partition of payments_202008 for values with (modulus 180, remainder 28) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_29 partition of payments_202008 for values with (modulus 180, remainder 29) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_30 partition of payments_202008 for values with (modulus 180, remainder 30) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_31 partition of payments_202008 for values with (modulus 180, remainder 31) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_32 partition of payments_202008 for values with (modulus 180, remainder 32) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_33 partition of payments_202008 for values with (modulus 180, remainder 33) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_34 partition of payments_202008 for values with (modulus 180, remainder 34) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_35 partition of payments_202008 for values with (modulus 180, remainder 35) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_36 partition of payments_202008 for values with (modulus 180, remainder 36) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_37 partition of payments_202008 for values with (modulus 180, remainder 37) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_38 partition of payments_202008 for values with (modulus 180, remainder 38) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_39 partition of payments_202008 for values with (modulus 180, remainder 39) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_40 partition of payments_202008 for values with (modulus 180, remainder 40) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_41 partition of payments_202008 for values with (modulus 180, remainder 41) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_42 partition of payments_202008 for values with (modulus 180, remainder 42) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_43 partition of payments_202008 for values with (modulus 180, remainder 43) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_44 partition of payments_202008 for values with (modulus 180, remainder 44) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_45 partition of payments_202008 for values with (modulus 180, remainder 45) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_46 partition of payments_202008 for values with (modulus 180, remainder 46) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_47 partition of payments_202008 for values with (modulus 180, remainder 47) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_48 partition of payments_202008 for values with (modulus 180, remainder 48) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_49 partition of payments_202008 for values with (modulus 180, remainder 49) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_50 partition of payments_202008 for values with (modulus 180, remainder 50) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_51 partition of payments_202008 for values with (modulus 180, remainder 51) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_52 partition of payments_202008 for values with (modulus 180, remainder 52) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_53 partition of payments_202008 for values with (modulus 180, remainder 53) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_54 partition of payments_202008 for values with (modulus 180, remainder 54) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_55 partition of payments_202008 for values with (modulus 180, remainder 55) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_56 partition of payments_202008 for values with (modulus 180, remainder 56) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_57 partition of payments_202008 for values with (modulus 180, remainder 57) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_58 partition of payments_202008 for values with (modulus 180, remainder 58) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_59 partition of payments_202008 for values with (modulus 180, remainder 59) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_60 partition of payments_202008 for values with (modulus 180, remainder 60) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_61 partition of payments_202008 for values with (modulus 180, remainder 61) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_62 partition of payments_202008 for values with (modulus 180, remainder 62) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_63 partition of payments_202008 for values with (modulus 180, remainder 63) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_64 partition of payments_202008 for values with (modulus 180, remainder 64) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_65 partition of payments_202008 for values with (modulus 180, remainder 65) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_66 partition of payments_202008 for values with (modulus 180, remainder 66) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_67 partition of payments_202008 for values with (modulus 180, remainder 67) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_68 partition of payments_202008 for values with (modulus 180, remainder 68) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_69 partition of payments_202008 for values with (modulus 180, remainder 69) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_70 partition of payments_202008 for values with (modulus 180, remainder 70) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_71 partition of payments_202008 for values with (modulus 180, remainder 71) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_72 partition of payments_202008 for values with (modulus 180, remainder 72) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_73 partition of payments_202008 for values with (modulus 180, remainder 73) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_74 partition of payments_202008 for values with (modulus 180, remainder 74) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_75 partition of payments_202008 for values with (modulus 180, remainder 75) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_76 partition of payments_202008 for values with (modulus 180, remainder 76) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_77 partition of payments_202008 for values with (modulus 180, remainder 77) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_78 partition of payments_202008 for values with (modulus 180, remainder 78) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_79 partition of payments_202008 for values with (modulus 180, remainder 79) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_80 partition of payments_202008 for values with (modulus 180, remainder 80) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_81 partition of payments_202008 for values with (modulus 180, remainder 81) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_82 partition of payments_202008 for values with (modulus 180, remainder 82) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_83 partition of payments_202008 for values with (modulus 180, remainder 83) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_84 partition of payments_202008 for values with (modulus 180, remainder 84) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_85 partition of payments_202008 for values with (modulus 180, remainder 85) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_86 partition of payments_202008 for values with (modulus 180, remainder 86) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_87 partition of payments_202008 for values with (modulus 180, remainder 87) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_88 partition of payments_202008 for values with (modulus 180, remainder 88) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_89 partition of payments_202008 for values with (modulus 180, remainder 89) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_90 partition of payments_202008 for values with (modulus 180, remainder 90) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_91 partition of payments_202008 for values with (modulus 180, remainder 91) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_92 partition of payments_202008 for values with (modulus 180, remainder 92) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_93 partition of payments_202008 for values with (modulus 180, remainder 93) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_94 partition of payments_202008 for values with (modulus 180, remainder 94) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_95 partition of payments_202008 for values with (modulus 180, remainder 95) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_96 partition of payments_202008 for values with (modulus 180, remainder 96) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_97 partition of payments_202008 for values with (modulus 180, remainder 97) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_98 partition of payments_202008 for values with (modulus 180, remainder 98) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_99 partition of payments_202008 for values with (modulus 180, remainder 99) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_100 partition of payments_202008 for values with (modulus 180, remainder 100) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_101 partition of payments_202008 for values with (modulus 180, remainder 101) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_102 partition of payments_202008 for values with (modulus 180, remainder 102) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_103 partition of payments_202008 for values with (modulus 180, remainder 103) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_104 partition of payments_202008 for values with (modulus 180, remainder 104) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_105 partition of payments_202008 for values with (modulus 180, remainder 105) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_106 partition of payments_202008 for values with (modulus 180, remainder 106) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_107 partition of payments_202008 for values with (modulus 180, remainder 107) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_108 partition of payments_202008 for values with (modulus 180, remainder 108) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_109 partition of payments_202008 for values with (modulus 180, remainder 109) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_110 partition of payments_202008 for values with (modulus 180, remainder 110) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_111 partition of payments_202008 for values with (modulus 180, remainder 111) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_112 partition of payments_202008 for values with (modulus 180, remainder 112) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_113 partition of payments_202008 for values with (modulus 180, remainder 113) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_114 partition of payments_202008 for values with (modulus 180, remainder 114) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_115 partition of payments_202008 for values with (modulus 180, remainder 115) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_116 partition of payments_202008 for values with (modulus 180, remainder 116) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_117 partition of payments_202008 for values with (modulus 180, remainder 117) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_118 partition of payments_202008 for values with (modulus 180, remainder 118) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_119 partition of payments_202008 for values with (modulus 180, remainder 119) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_120 partition of payments_202008 for values with (modulus 180, remainder 120) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_121 partition of payments_202008 for values with (modulus 180, remainder 121) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_122 partition of payments_202008 for values with (modulus 180, remainder 122) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_123 partition of payments_202008 for values with (modulus 180, remainder 123) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_124 partition of payments_202008 for values with (modulus 180, remainder 124) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_125 partition of payments_202008 for values with (modulus 180, remainder 125) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_126 partition of payments_202008 for values with (modulus 180, remainder 126) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_127 partition of payments_202008 for values with (modulus 180, remainder 127) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_128 partition of payments_202008 for values with (modulus 180, remainder 128) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_129 partition of payments_202008 for values with (modulus 180, remainder 129) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_130 partition of payments_202008 for values with (modulus 180, remainder 130) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_131 partition of payments_202008 for values with (modulus 180, remainder 131) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_132 partition of payments_202008 for values with (modulus 180, remainder 132) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_133 partition of payments_202008 for values with (modulus 180, remainder 133) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_134 partition of payments_202008 for values with (modulus 180, remainder 134) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_135 partition of payments_202008 for values with (modulus 180, remainder 135) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_136 partition of payments_202008 for values with (modulus 180, remainder 136) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_137 partition of payments_202008 for values with (modulus 180, remainder 137) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_138 partition of payments_202008 for values with (modulus 180, remainder 138) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_139 partition of payments_202008 for values with (modulus 180, remainder 139) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_140 partition of payments_202008 for values with (modulus 180, remainder 140) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_141 partition of payments_202008 for values with (modulus 180, remainder 141) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_142 partition of payments_202008 for values with (modulus 180, remainder 142) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_143 partition of payments_202008 for values with (modulus 180, remainder 143) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_144 partition of payments_202008 for values with (modulus 180, remainder 144) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_145 partition of payments_202008 for values with (modulus 180, remainder 145) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_146 partition of payments_202008 for values with (modulus 180, remainder 146) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_147 partition of payments_202008 for values with (modulus 180, remainder 147) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_148 partition of payments_202008 for values with (modulus 180, remainder 148) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_149 partition of payments_202008 for values with (modulus 180, remainder 149) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_150 partition of payments_202008 for values with (modulus 180, remainder 150) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_151 partition of payments_202008 for values with (modulus 180, remainder 151) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_152 partition of payments_202008 for values with (modulus 180, remainder 152) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_153 partition of payments_202008 for values with (modulus 180, remainder 153) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_154 partition of payments_202008 for values with (modulus 180, remainder 154) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_155 partition of payments_202008 for values with (modulus 180, remainder 155) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_156 partition of payments_202008 for values with (modulus 180, remainder 156) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_157 partition of payments_202008 for values with (modulus 180, remainder 157) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_158 partition of payments_202008 for values with (modulus 180, remainder 158) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_159 partition of payments_202008 for values with (modulus 180, remainder 159) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_160 partition of payments_202008 for values with (modulus 180, remainder 160) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_161 partition of payments_202008 for values with (modulus 180, remainder 161) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_162 partition of payments_202008 for values with (modulus 180, remainder 162) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_163 partition of payments_202008 for values with (modulus 180, remainder 163) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_164 partition of payments_202008 for values with (modulus 180, remainder 164) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_165 partition of payments_202008 for values with (modulus 180, remainder 165) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_166 partition of payments_202008 for values with (modulus 180, remainder 166) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_167 partition of payments_202008 for values with (modulus 180, remainder 167) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_168 partition of payments_202008 for values with (modulus 180, remainder 168) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_169 partition of payments_202008 for values with (modulus 180, remainder 169) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_170 partition of payments_202008 for values with (modulus 180, remainder 170) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_171 partition of payments_202008 for values with (modulus 180, remainder 171) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_172 partition of payments_202008 for values with (modulus 180, remainder 172) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_173 partition of payments_202008 for values with (modulus 180, remainder 173) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_174 partition of payments_202008 for values with (modulus 180, remainder 174) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_175 partition of payments_202008 for values with (modulus 180, remainder 175) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_176 partition of payments_202008 for values with (modulus 180, remainder 176) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_177 partition of payments_202008 for values with (modulus 180, remainder 177) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_178 partition of payments_202008 for values with (modulus 180, remainder 178) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202008_hash_179 partition of payments_202008 for values with (modulus 180, remainder 179) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_0 partition of payments_202009 for values with (modulus 180, remainder 0) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_1 partition of payments_202009 for values with (modulus 180, remainder 1) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_2 partition of payments_202009 for values with (modulus 180, remainder 2) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_3 partition of payments_202009 for values with (modulus 180, remainder 3) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_4 partition of payments_202009 for values with (modulus 180, remainder 4) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_5 partition of payments_202009 for values with (modulus 180, remainder 5) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_6 partition of payments_202009 for values with (modulus 180, remainder 6) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_7 partition of payments_202009 for values with (modulus 180, remainder 7) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_8 partition of payments_202009 for values with (modulus 180, remainder 8) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_9 partition of payments_202009 for values with (modulus 180, remainder 9) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_10 partition of payments_202009 for values with (modulus 180, remainder 10) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_11 partition of payments_202009 for values with (modulus 180, remainder 11) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_12 partition of payments_202009 for values with (modulus 180, remainder 12) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_13 partition of payments_202009 for values with (modulus 180, remainder 13) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_14 partition of payments_202009 for values with (modulus 180, remainder 14) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_15 partition of payments_202009 for values with (modulus 180, remainder 15) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_16 partition of payments_202009 for values with (modulus 180, remainder 16) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_17 partition of payments_202009 for values with (modulus 180, remainder 17) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_18 partition of payments_202009 for values with (modulus 180, remainder 18) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_19 partition of payments_202009 for values with (modulus 180, remainder 19) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_20 partition of payments_202009 for values with (modulus 180, remainder 20) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_21 partition of payments_202009 for values with (modulus 180, remainder 21) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_22 partition of payments_202009 for values with (modulus 180, remainder 22) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_23 partition of payments_202009 for values with (modulus 180, remainder 23) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_24 partition of payments_202009 for values with (modulus 180, remainder 24) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_25 partition of payments_202009 for values with (modulus 180, remainder 25) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_26 partition of payments_202009 for values with (modulus 180, remainder 26) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_27 partition of payments_202009 for values with (modulus 180, remainder 27) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_28 partition of payments_202009 for values with (modulus 180, remainder 28) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_29 partition of payments_202009 for values with (modulus 180, remainder 29) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_30 partition of payments_202009 for values with (modulus 180, remainder 30) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_31 partition of payments_202009 for values with (modulus 180, remainder 31) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_32 partition of payments_202009 for values with (modulus 180, remainder 32) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_33 partition of payments_202009 for values with (modulus 180, remainder 33) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_34 partition of payments_202009 for values with (modulus 180, remainder 34) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_35 partition of payments_202009 for values with (modulus 180, remainder 35) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_36 partition of payments_202009 for values with (modulus 180, remainder 36) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_37 partition of payments_202009 for values with (modulus 180, remainder 37) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_38 partition of payments_202009 for values with (modulus 180, remainder 38) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_39 partition of payments_202009 for values with (modulus 180, remainder 39) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_40 partition of payments_202009 for values with (modulus 180, remainder 40) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_41 partition of payments_202009 for values with (modulus 180, remainder 41) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_42 partition of payments_202009 for values with (modulus 180, remainder 42) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_43 partition of payments_202009 for values with (modulus 180, remainder 43) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_44 partition of payments_202009 for values with (modulus 180, remainder 44) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_45 partition of payments_202009 for values with (modulus 180, remainder 45) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_46 partition of payments_202009 for values with (modulus 180, remainder 46) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_47 partition of payments_202009 for values with (modulus 180, remainder 47) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_48 partition of payments_202009 for values with (modulus 180, remainder 48) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_49 partition of payments_202009 for values with (modulus 180, remainder 49) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_50 partition of payments_202009 for values with (modulus 180, remainder 50) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_51 partition of payments_202009 for values with (modulus 180, remainder 51) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_52 partition of payments_202009 for values with (modulus 180, remainder 52) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_53 partition of payments_202009 for values with (modulus 180, remainder 53) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_54 partition of payments_202009 for values with (modulus 180, remainder 54) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_55 partition of payments_202009 for values with (modulus 180, remainder 55) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_56 partition of payments_202009 for values with (modulus 180, remainder 56) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_57 partition of payments_202009 for values with (modulus 180, remainder 57) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_58 partition of payments_202009 for values with (modulus 180, remainder 58) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_59 partition of payments_202009 for values with (modulus 180, remainder 59) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_60 partition of payments_202009 for values with (modulus 180, remainder 60) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_61 partition of payments_202009 for values with (modulus 180, remainder 61) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_62 partition of payments_202009 for values with (modulus 180, remainder 62) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_63 partition of payments_202009 for values with (modulus 180, remainder 63) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_64 partition of payments_202009 for values with (modulus 180, remainder 64) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_65 partition of payments_202009 for values with (modulus 180, remainder 65) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_66 partition of payments_202009 for values with (modulus 180, remainder 66) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_67 partition of payments_202009 for values with (modulus 180, remainder 67) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_68 partition of payments_202009 for values with (modulus 180, remainder 68) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_69 partition of payments_202009 for values with (modulus 180, remainder 69) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_70 partition of payments_202009 for values with (modulus 180, remainder 70) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_71 partition of payments_202009 for values with (modulus 180, remainder 71) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_72 partition of payments_202009 for values with (modulus 180, remainder 72) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_73 partition of payments_202009 for values with (modulus 180, remainder 73) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_74 partition of payments_202009 for values with (modulus 180, remainder 74) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_75 partition of payments_202009 for values with (modulus 180, remainder 75) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_76 partition of payments_202009 for values with (modulus 180, remainder 76) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_77 partition of payments_202009 for values with (modulus 180, remainder 77) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_78 partition of payments_202009 for values with (modulus 180, remainder 78) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_79 partition of payments_202009 for values with (modulus 180, remainder 79) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_80 partition of payments_202009 for values with (modulus 180, remainder 80) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_81 partition of payments_202009 for values with (modulus 180, remainder 81) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_82 partition of payments_202009 for values with (modulus 180, remainder 82) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_83 partition of payments_202009 for values with (modulus 180, remainder 83) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_84 partition of payments_202009 for values with (modulus 180, remainder 84) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_85 partition of payments_202009 for values with (modulus 180, remainder 85) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_86 partition of payments_202009 for values with (modulus 180, remainder 86) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_87 partition of payments_202009 for values with (modulus 180, remainder 87) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_88 partition of payments_202009 for values with (modulus 180, remainder 88) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_89 partition of payments_202009 for values with (modulus 180, remainder 89) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_90 partition of payments_202009 for values with (modulus 180, remainder 90) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_91 partition of payments_202009 for values with (modulus 180, remainder 91) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_92 partition of payments_202009 for values with (modulus 180, remainder 92) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_93 partition of payments_202009 for values with (modulus 180, remainder 93) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_94 partition of payments_202009 for values with (modulus 180, remainder 94) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_95 partition of payments_202009 for values with (modulus 180, remainder 95) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_96 partition of payments_202009 for values with (modulus 180, remainder 96) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_97 partition of payments_202009 for values with (modulus 180, remainder 97) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_98 partition of payments_202009 for values with (modulus 180, remainder 98) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_99 partition of payments_202009 for values with (modulus 180, remainder 99) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_100 partition of payments_202009 for values with (modulus 180, remainder 100) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_101 partition of payments_202009 for values with (modulus 180, remainder 101) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_102 partition of payments_202009 for values with (modulus 180, remainder 102) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_103 partition of payments_202009 for values with (modulus 180, remainder 103) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_104 partition of payments_202009 for values with (modulus 180, remainder 104) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_105 partition of payments_202009 for values with (modulus 180, remainder 105) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_106 partition of payments_202009 for values with (modulus 180, remainder 106) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_107 partition of payments_202009 for values with (modulus 180, remainder 107) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_108 partition of payments_202009 for values with (modulus 180, remainder 108) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_109 partition of payments_202009 for values with (modulus 180, remainder 109) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_110 partition of payments_202009 for values with (modulus 180, remainder 110) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_111 partition of payments_202009 for values with (modulus 180, remainder 111) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_112 partition of payments_202009 for values with (modulus 180, remainder 112) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_113 partition of payments_202009 for values with (modulus 180, remainder 113) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_114 partition of payments_202009 for values with (modulus 180, remainder 114) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_115 partition of payments_202009 for values with (modulus 180, remainder 115) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_116 partition of payments_202009 for values with (modulus 180, remainder 116) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_117 partition of payments_202009 for values with (modulus 180, remainder 117) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_118 partition of payments_202009 for values with (modulus 180, remainder 118) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_119 partition of payments_202009 for values with (modulus 180, remainder 119) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_120 partition of payments_202009 for values with (modulus 180, remainder 120) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_121 partition of payments_202009 for values with (modulus 180, remainder 121) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_122 partition of payments_202009 for values with (modulus 180, remainder 122) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_123 partition of payments_202009 for values with (modulus 180, remainder 123) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_124 partition of payments_202009 for values with (modulus 180, remainder 124) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_125 partition of payments_202009 for values with (modulus 180, remainder 125) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_126 partition of payments_202009 for values with (modulus 180, remainder 126) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_127 partition of payments_202009 for values with (modulus 180, remainder 127) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_128 partition of payments_202009 for values with (modulus 180, remainder 128) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_129 partition of payments_202009 for values with (modulus 180, remainder 129) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_130 partition of payments_202009 for values with (modulus 180, remainder 130) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_131 partition of payments_202009 for values with (modulus 180, remainder 131) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_132 partition of payments_202009 for values with (modulus 180, remainder 132) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_133 partition of payments_202009 for values with (modulus 180, remainder 133) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_134 partition of payments_202009 for values with (modulus 180, remainder 134) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_135 partition of payments_202009 for values with (modulus 180, remainder 135) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_136 partition of payments_202009 for values with (modulus 180, remainder 136) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_137 partition of payments_202009 for values with (modulus 180, remainder 137) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_138 partition of payments_202009 for values with (modulus 180, remainder 138) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_139 partition of payments_202009 for values with (modulus 180, remainder 139) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_140 partition of payments_202009 for values with (modulus 180, remainder 140) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_141 partition of payments_202009 for values with (modulus 180, remainder 141) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_142 partition of payments_202009 for values with (modulus 180, remainder 142) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_143 partition of payments_202009 for values with (modulus 180, remainder 143) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_144 partition of payments_202009 for values with (modulus 180, remainder 144) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_145 partition of payments_202009 for values with (modulus 180, remainder 145) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_146 partition of payments_202009 for values with (modulus 180, remainder 146) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_147 partition of payments_202009 for values with (modulus 180, remainder 147) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_148 partition of payments_202009 for values with (modulus 180, remainder 148) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_149 partition of payments_202009 for values with (modulus 180, remainder 149) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_150 partition of payments_202009 for values with (modulus 180, remainder 150) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_151 partition of payments_202009 for values with (modulus 180, remainder 151) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_152 partition of payments_202009 for values with (modulus 180, remainder 152) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_153 partition of payments_202009 for values with (modulus 180, remainder 153) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_154 partition of payments_202009 for values with (modulus 180, remainder 154) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_155 partition of payments_202009 for values with (modulus 180, remainder 155) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_156 partition of payments_202009 for values with (modulus 180, remainder 156) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_157 partition of payments_202009 for values with (modulus 180, remainder 157) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_158 partition of payments_202009 for values with (modulus 180, remainder 158) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_159 partition of payments_202009 for values with (modulus 180, remainder 159) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_160 partition of payments_202009 for values with (modulus 180, remainder 160) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_161 partition of payments_202009 for values with (modulus 180, remainder 161) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_162 partition of payments_202009 for values with (modulus 180, remainder 162) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_163 partition of payments_202009 for values with (modulus 180, remainder 163) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_164 partition of payments_202009 for values with (modulus 180, remainder 164) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_165 partition of payments_202009 for values with (modulus 180, remainder 165) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_166 partition of payments_202009 for values with (modulus 180, remainder 166) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_167 partition of payments_202009 for values with (modulus 180, remainder 167) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_168 partition of payments_202009 for values with (modulus 180, remainder 168) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_169 partition of payments_202009 for values with (modulus 180, remainder 169) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_170 partition of payments_202009 for values with (modulus 180, remainder 170) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_171 partition of payments_202009 for values with (modulus 180, remainder 171) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_172 partition of payments_202009 for values with (modulus 180, remainder 172) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_173 partition of payments_202009 for values with (modulus 180, remainder 173) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_174 partition of payments_202009 for values with (modulus 180, remainder 174) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_175 partition of payments_202009 for values with (modulus 180, remainder 175) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_176 partition of payments_202009 for values with (modulus 180, remainder 176) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_177 partition of payments_202009 for values with (modulus 180, remainder 177) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_178 partition of payments_202009 for values with (modulus 180, remainder 178) partition by hash(payment_id);`); err != nil {
    return err
  }
  if _, err := tx.Exec(`create table payments_2020_payments_202009_hash_179 partition of payments_202009 for values with (modulus 180, remainder 179) partition by hash(payment_id);`); err != nil {
    return err
  }

  return nil
}

func Down20200728162603(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
