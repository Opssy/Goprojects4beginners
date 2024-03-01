package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"golang.org/x/net/html"
)

func hasHref(t html.Token) (ok bool, href string){	
	for _, a := range t.Attr{
		if a.Key == "href"{
			href = a.Val
			ok = true
		}
	}
	return
}

func crawl( url string, ch chan string, chFinished chan bool){
	resp, err := http.Get(url)

	defer func(){
		chFinished  <- true
	}()

	if err != nil{
		fmt.Println("ERR: failed to crawl:", url)
		return
	}
	b := resp.Body

	defer b.Close()

	z :=  html.NewTokenizer(b)

	for{
		tt := z.Next()
		switch {
		case tt == html.ErrorToken:
			return
		case tt == html.StartTagToken:
		  t := z.Token()

		  isAnchor := t.Data == "a"
		  if !isAnchor{
		  	continue
		  }
		  ok, url := hasHref(t)
		  if !ok{
			continue
		  }
		  hasProto := strings.Index(url, "http") == 0
		  if hasProto{
			ch <- url
		  }
		}
	}
}
func main() {

	foundUrl := make(map[string]bool)

	seedUrls := os.Args[1:]

	chUrls := make(chan string)
	chFinished := make(chan bool)

	for _, url := range seedUrls{
		go crawl(url, chUrls, chFinished)	
	}
	for c := 0 ; c<len(seedUrls);{
		select{
		case url := <-chUrls:
			foundUrl[url] = true
		case  <- chFinished:
			c++
		}
	}
   fmt.Println("\nFound", len(foundUrl), "unique urls:\n")

   for url, _ :=  range foundUrl{
	   fmt.Println(" - " + url)
   }

   close(chUrls)
}
