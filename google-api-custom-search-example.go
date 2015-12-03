package main

import (
	"flag"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	customsearch "google.golang.org/api/customsearch/v1"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var q = flag.String("q", "", "Query string to test against.")
var domain = flag.String("d", "", "Domain and path to look for: domain.com/about-us")
var id = flag.String("id", "", "Search Engine ID")

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: google-api-custom-search-example -id [Search Engine ID] -q [Query String] -d [domain.com/path to look for]\n\n")
	os.Exit(2)
}

type Result struct {
	Position int64
	Result   *customsearch.Result
}

func main() {

	// Parse those command line flags
	flag.Parse()

	// All flags are required.
	if flag.NFlag() != 3 {
		usage()
	}

	data, err := ioutil.ReadFile("search-key.json")
	if err != nil {
		log.Fatal(err)
	}
	//Get the config from the json key file with the correct scope
	conf, err := google.JWTConfigFromJSON(data, "https://www.googleapis.com/auth/cse")
	if err != nil {
		log.Fatal(err)
	}

	// Initiate an http.Client. The following GET request will be
	// authorized and authenticated on the behalf of
	// your service account.
	client := conf.Client(oauth2.NoContext)

	cseService, err := customsearch.New(client)
	search := cseService.Cse.List(*q)
	search.Cx(*id)

	//Thinking about searching a particular place?
	//search.Gl("Chattanooga, TN")

	result := doSearch(search)

	if result.Position == 0 {
		log.Fatal("No results found in the top 10 pages.\n")
	}

	fmt.Printf("Result found for \"%s\"!\n", *domain)
	fmt.Printf("Position: %d\n", result.Position)
	fmt.Printf("Url: %s\n", result.Result.Link)
	fmt.Printf("Title: %s\n", result.Result.Title)
	fmt.Printf("Snippet: %s\n", result.Result.Snippet)
}

func doSearch(search *customsearch.CseListCall) (result Result) {

	start := int64(1)

	// CSE Limits you to 10 pages of results with max 10 results per page
	for start < 100 {
		search.Start(start)
		call, err := search.Do()
		if err != nil {
			log.Fatal(err)
		}

		position, csResult := findDomain(call.Items, start)

		if csResult != nil {
			result = Result{
				Position: position,
				Result:   csResult,
			}
			return
		}

		// No more search results?
		if call.SearchInformation.TotalResults < start {
			return
		}
		start = start + 10
	}

	return
}

func findDomain(results []*customsearch.Result, start int64) (position int64, result *customsearch.Result) {
	for index, r := range results {
		if strings.Contains(r.Link, *domain) {
			return int64(index) + start, r
		}
	}
	return 0, nil
}
