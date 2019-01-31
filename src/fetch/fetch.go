package fetch

import (
	"logger"
)

// Page stores all Div entities for a URL.
type Page struct {
	URL      string
	Entities map[string]Div
}

// NewPage initializes a new page struct.
func NewPage(url string) Page {
	logger.Init("NewPage", "New Page @"+url)
	page := Page{
		URL: url,
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

// Test info
func Test() {

	l := logger.Logger{
		Name: "Fetch",
	}

	l.Log("It works.")

	l.Log("Testing page struct.")
	p := NewPage("https://news.ycombinator.com/item?id=18499843")
	l.Log("Adding a map div.")
	l.Log("Scraping " + p.URL)
	p.Entities["one"] = Div{
		Text:   "one",
		Emails: []string{"EmailOne", "EmailTwo"},
	}
}
