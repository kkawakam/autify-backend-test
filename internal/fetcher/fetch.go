package fetch

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

// Runs a 'task' which will be the following:
// 1. Fetch the HTML from a given url
// 2. Persist the HTML to disk
// 3. *Optionally* parse the HTML to extract metadata
// 4. *Optionally* persist images to disk
func task(url string, wg *sync.WaitGroup) {
	println(url)
	wg.Done()
}

func Run(urls []string, isPrintMetadata bool) {
	// Create a new directory that will contain the latest html files
	if len(urls) > 0 {
		currentTime := strconv.FormatInt(time.Now().UnixMilli(), 10)
		err := os.Mkdir(fmt.Sprintf("results-%s", currentTime), os.ModeDir)
		if err != nil {
			fmt.Printf("Unable to create directory %s for saving html", currentTime)
		}
	}

	// Create goroutines per URL
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go task(url, &wg)
	}
	wg.Wait()
}
