#!/bin/sh

tables=("payments_202007" "payments_202008" "payments_202009")

for i in "${tables[@]}"
do
   :
  for x in {0..2}
  do

#  echo "if _, err := tx.Exec(\`""create table payments_2020_$i""_""hash_$x partition of $i for values with (modulus 180, remainder $x);""\`); err != nil {
#            return err
#  }"

  echo "create table payments_2020_$i""_""hash_$x partition of $i for values with (modulus 3, remainder $x);"
  done

done


