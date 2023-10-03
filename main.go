package main

import (
	"flag"

	"github.com/kkawakam/autify-backend-test/internal/fetcher"
)

func main() {
	var printMetadata bool
	var outputDirectory string
	flag.BoolVar(&printMetadata, "metadata", false, "record metadata for recorded sites")
	flag.StringVar(&outputDirectory, "output_directory", "", "where results where be persisted NOTE: will not create the directory")
	flag.Parse()
	fetcher.Run(flag.Args(), printMetadata, outputDirectory)
}
