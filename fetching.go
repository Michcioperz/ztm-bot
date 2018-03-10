package main

import (
	"errors"
	//"github.com/PuerkitoBio/goquery"
	"github.com/SlyMarbo/rss"
)

func fetchBasicFeed() ([]*rss.Item, error) {
	feed, err := rss.Fetch("http://www.ztm.waw.pl/rss.php?l=1&IDRss=6")
	if err != nil {
		return []*rss.Item{}, err
	}
	return feed.Items, nil
}

func enhanceBasicItem(it rss.Item) (*rss.Item, error) {
	return nil, errors.New("not implemented yet")
}

func enhanceBasicFeed(feed []*rss.Item) ([]*rss.Item, error) {
	newFeed := []*rss.Item{}
	for _, it := range feed {
		newIt, err := enhanceBasicItem(*it)
		if err != nil {
			return newFeed, err
		}
		newFeed = append(newFeed, newIt)
	}
	return newFeed, nil
}
