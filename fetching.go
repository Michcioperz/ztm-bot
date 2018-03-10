package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/SlyMarbo/rss"
)

func fetchBasicFeed() ([]*rss.Item, error) {
	feed, err := rss.Fetch("http://www.ztm.waw.pl/rss.php?l=1&IDRss=6")
	if err != nil {
		return []*rss.Item{}, err
	}
	return feed.Items, nil
}

func enhanceBasicItem(it *rss.Item) (rss.Item, error) {
	doc, err := goquery.NewDocument(it.Link)
	if err != nil {
		return *it, err
	}
	content := doc.Find("#PageContent").Text()
	newIt := rss.Item{
		Title: it.Summary,
		Summary: it.Title,
		Content: content,
		Link: it.Link,
	}
	return newIt, nil
}

func enhanceBasicFeed(feed []*rss.Item) ([]*rss.Item, error) {
	newFeed := []*rss.Item{}
	for _, it := range feed {
		newIt, err := enhanceBasicItem(it)
		if err != nil {
			return newFeed, err
		}
		newFeed = append(newFeed, &newIt)
	}
	return newFeed, nil
}

func FetchEnhancedFeed() ([]*rss.Item, error) {
	basicFeed, err := fetchBasicFeed()
	if err != nil {
		return basicFeed, err
	}
	return enhanceBasicFeed(basicFeed)
}
