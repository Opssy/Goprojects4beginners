package main

import (
 "fmt"
 "log"
 "strings"

 "github.com/gocolly/colly/v2"
)
func main() {
   // create a new colly collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("hackerspaces.org", "wiki.hackerspaces.org"),
	)

   c.OnHTML("a[href]", func(e *colly.HTMLElement) {

	 link := e.Attr("href")

	 fmt.Printf("Link found: %q -> %s\n", e.Text, link)
      
	 
   })
}