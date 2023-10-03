package fetcher

import (
	"errors"
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/html"
)

func getHtml(url string) (string, *time.Time, error) {
	fetchTime := time.Now()
	res, err := http.Get(url)
	if err != nil {
		return "", nil, err
	}
	defer res.Body.Close()

	mediatype, _, err := mime.ParseMediaType(res.Header.Get("content-type"))
	if err != nil {
		return "", nil, err
	}
	if mediatype != "text/html" {
		return "", nil, errors.New(url + " does not contain text/html response content type")
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", nil, err
	}
	return string(bodyBytes), &fetchTime, nil
}

func writeHtmlToDisk(outputDirectory string, host string, path string, body string) error {
	filename := host + path + ".html"
	filename = strings.ReplaceAll(filename, "/", "-") // Remove all backslashes from the path
	f, err := os.Create(outputDirectory + "/" + filename)
	if err != nil {
		return err
	}
	defer f.Close()

	f.WriteString(body)
	f.Sync()

	return nil
}

type metadata struct {
	NumLinks int
	Images   int
}

func recordMetadata(rawHtml string) (*metadata, error) {
	root, err := html.Parse(strings.NewReader(rawHtml))
	if err != nil {
		return nil, err
	}
	bfs := []*html.Node{root}
	linkCount := 0
	imgCount := 0
	for len(bfs) > 0 {
		cur := bfs[0]

		if cur.Type == html.ElementNode {
			if cur.Data == "a" {
				linkCount++
			} else if cur.Data == "img" {
				imgCount++
			}
		}

		bfs = bfs[1:]
		iter := cur.FirstChild
		for iter != nil {
			bfs = append(bfs, iter)
			iter = iter.NextSibling
		}
	}

	return &metadata{
		NumLinks: linkCount,
		Images:   imgCount,
	}, nil
}

// Runs a 'task' which will be the following:
// 1. Fetch the HTML from a given url
// 2. Persist the HTML to disk
// 3. *Optionally* parse the HTML to extract metadata
func task(outputDirectory string, rawUrl string, isPrintMetadata bool, wg *sync.WaitGroup) {
	defer wg.Done()

	// Check if we got an url
	u, err := url.Parse(rawUrl)
	if err != nil || (u.Scheme == "" && u.Host == "") {
		log.Print("Invalid URL:", rawUrl)
		return
	}

	// 1. Fetch HTML
	body, fetchTime, err := getHtml(rawUrl)
	if err != nil {
		log.Println(err)
		return
	}

	// 2. Persist HTML to disk
	err = writeHtmlToDisk(outputDirectory, u.Host, u.Path, body)
	if err != nil {
		log.Println(err)
		return
	}

	// 3. Parse HTML and generate metadata
	if isPrintMetadata {
		parsedMetadata, err := recordMetadata(body)
		if err != nil {
			log.Println("Unable to parse html retrieved from " + rawUrl)
			return
		}
		println("site: ", u.Host+u.Path)
		println("num_links: ", parsedMetadata.NumLinks)
		println("images: ", parsedMetadata.Images)
		println("last_fetch: ", fetchTime.Format("Mon Jan 02 2006 15:04 MST"))
	}
}

// This the entrypoint and is responsible for
// 1. Creating an output directory that will contain the html that was retrieved
// 2. Spawn a separate goroutine for each url
func Run(urls []string, isPrintMetadata bool, outputDirectory string) {
	currentTime := strconv.FormatInt(time.Now().UnixMilli(), 10)
	resultDirectory := fmt.Sprintf("%s/results-%s", outputDirectory, currentTime)
	if len(urls) > 0 {
		var wg sync.WaitGroup
		defer wg.Wait()

		// Create a new directory that will contain the latest html files
		err := os.MkdirAll(resultDirectory, 0700)
		if err != nil {
			fmt.Printf("Unable to create directory %s for saving html", currentTime)
		}

		// Create goroutines per URL
		for _, url := range urls {
			wg.Add(1)
			go task(resultDirectory, url, isPrintMetadata, &wg)
		}
	}
}
