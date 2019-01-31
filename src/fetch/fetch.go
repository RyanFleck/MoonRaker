package fetch

import (
	"log"
	"logger"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Page stores all Div entities for a ID.
type Page struct {
	ID       string
	Entities map[string]Div
}

// NewPage initializes a new page struct.
func NewPage(id string) Page {
	logger.Init("NewPage", "New Page @"+id)
	page := Page{
		ID: id,
	}
	page.Entities = make(map[string]Div)
	return page
}

// Div Stores one div record, with parsed information.
type Div struct {
	Text      string
	Emails    []string
	Buzzwords []string
}

// Scrape fills the Entity map with Div structures.
func (p Page) Scrape() {

	// Setup logging.
	l := logger.Logger{Name: "Scrape"}
	l.Log("Scraping https://news.ycombinator.com/item?id=" + p.ID)

	// Setup http client.
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// http client -> set up request
	request, err := http.NewRequest("GET", "https://news.ycombinator.com/item?id="+p.ID, nil)
	if err != nil {
		l.Err("Could not setup request @" + p.ID)
	}
	request.Header.Set("User-Agent", "MoonRaker Go Program")

	// http client -> send request
	l.Log("Sending request to client.")
	response, err := client.Do(request)
	if err != nil {
		l.Err("Client failed to fetch @" + p.ID)
	}

	// Wait for response, goquery
	l.Log("Using GoQuery to find all top-level comments.")
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}
	document.Find("table.comment-tree tbody tr td table tbody tr td.default").Each(func(index int, element *goquery.Selection) {
		body, err := element.Html()
		if err != nil {
			l.Log(body)
		}
	})
}

// Test info
func Test() {

	l := logger.Logger{
		Name: "Fetch",
	}

	l.Log("It works.")

	l.Log("Testing page struct.")
	p := NewPage("18800645")
	l.Log("Adding a map div.")
	l.Log("Scraping " + p.ID)

	p.Entities["one"] = Div{
		Text:   "one",
		Emails: []string{"EmailOne", "EmailTwo"},
	}

	p.Entities["two"] = Div{
		Text:   "two",
		Emails: []string{"EmailOne", "EmailTwo"},
	}

	// Alright
	for k, v := range p.Entities {
		l.Log("Company: " + k)
		for i := range v.Emails {
			l.Log("\tEmail: " + v.Emails[i])
		}
	}

	l.Log("Great, can loop through entries in map and access properties.")

	p.Scrape()
}
