package main

import (
	"flag"
	"log"
	"runtime/debug"

	"github.com/pavel-paulau/qb"
)

const _GOGC = 300

var (
	numWorkers, numDocs, docSize int64
	hostname                     string
)

func init() {
	debug.SetGCPercent(_GOGC)
}

func main() {
	flag.Int64Var(&numWorkers, "workers", 1, "number of workload threads")
	flag.Int64Var(&numDocs, "docs", 1e3, "number of documents to insert")
	flag.Int64Var(&docSize, "size", 512, "average size of the documents")

	flag.StringVar(&hostname, "hostname", "127.0.0.1", "Couchbase Server hostname")

	flag.Parse()

	err := initDatabase()
	if err != nil {
		log.Fatalf("database initialization failed: %v", err)
	}

	qb.Load(insert, numWorkers, numDocs, docSize)
}
