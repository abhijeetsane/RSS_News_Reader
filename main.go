package main

import (
	"fmt"
	"sync"

	"github.com/mmcdole/gofeed"
)

var rssList = [6]string{
	"http://rss.slashdot.org/Slashdot/slashdotMain",
	"https://feeds.twit.tv/floss.xml",
	"https://undeadly.org/cgi?action=rss",
	"http://www.NetBSD.org/changes/rss-netbsd.xml",
	"https://lobste.rs/t/programming.rss",
	"https://news.ycombinator.com/rss",
}

func printRssContents(wg *sync.WaitGroup, rssURL string) {
	defer wg.Done()
	rssParser := gofeed.NewParser()
	feed, err := rssParser.ParseURL(rssURL)
	if err != nil {
		fmt.Printf("\n\nURL data parsing failed for : %s \n", rssURL)
	} else {
		fmt.Printf("\nTitle of the RSS Feed is : %s \n", feed.Title)
		fmt.Printf("Number of items in feed is : %d \n", len(feed.Items))
		for i := 0; i < len(feed.Items); i++ {
			fmt.Printf("(%d) : %s \n", i, feed.Items[i].Title)
		}
	}
}
func main() {
	var wg sync.WaitGroup

	fmt.Printf("%s\n", "Rss NewReader")
	for lcv := 0; lcv < len(rssList); lcv++ {
		fmt.Println(rssList[lcv])
		wg.Add(1)
		go printRssContents(&wg, rssList[lcv])
	}
	wg.Wait()

}
