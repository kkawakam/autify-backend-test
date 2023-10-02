package main

import (
	"flag"
	"fmt"

	"github.com/kkawakam/autify-backend-test/internal/fetcher"
)

func main() {
	var printMetadata bool
	flag.BoolVar(&printMetadata, "metadata", false, "print metadata of fetched")
	flag.Parse()
	fmt.Println("Non-flag arguments:", flag.Args())
	fmt.Println("Print Metadata argument:", printMetadata)
	fetcher.Run(flag.Args(), printMetadata)
}
