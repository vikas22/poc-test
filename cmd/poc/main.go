package main

import (
  "db-poc/internal/bootstrap"
  "db-poc/internal/poc"
  "flag"
  "fmt"
  "github.com/prometheus/client_golang/prometheus/promhttp"
  "net/http"
)

var op string
var db string
var records int64
var concurrency int64

const (
	DefaultOp = "benchmark"
	DefaultDb = "postgresql"
  DefaultRecords = 100000
  DefaultConcurrency = 100
)

func init() {
  http.Handle("/metrics", promhttp.Handler())
  go http.ListenAndServe(":2193", nil)
}

func main() {
	bootstrap.Init()
	fmt.Println(bootstrap.Db)
	fmt.Println(bootstrap.Config)
	fmt.Println("Db Poc")

	flag.StringVar(&op, "op", DefaultOp, "Operation to be run")
	flag.StringVar(&db, "db", DefaultDb, "Database")
  flag.Int64Var(&records, "r", DefaultRecords, "Records")
  flag.Int64Var(&concurrency, "c", DefaultConcurrency, "Records")

	flag.Parse()

  switch op {
  case "bootstrap":
    bootstrap.Bootstrap()
  default:
    poc.RunPoc(records, concurrency)
  }

}
