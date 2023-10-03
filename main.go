package main

import (
	"flag"

	"github.com/kkawakam/autify-backend-test/internal/fetcher"
)

func main() {
	var printMetadata bool
	flag.BoolVar(&printMetadata, "metadata", false, "record metadata for recorded sites")
	flag.Parse()
	fetcher.Run(flag.Args(), printMetadata)
}
