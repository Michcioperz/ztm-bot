package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/SlyMarbo/rss"
	"fmt"
	"strings"
	"net/url"
	"strconv"
	"log"
)

func fetchRSSFeed() ([]*rss.Item, error) {
	feed, err := rss.Fetch("http://www.ztm.waw.pl/rss.php?l=1&IDRss=6")
	if err != nil {
		return []*rss.Item{}, err
	}
	return feed.Items, nil
}

type ZTMevent struct {
	Id int
	Lines []string
	Title string
	Content string
}

type EventParseError struct {
	i *rss.Item
	reason string
}

func (v EventParseError) Error() string {
	return fmt.Sprintf("problem with event parsing due to \"%s\" event=%+v", v.reason, v.i)
}

func (v *ZTMevent) String() string {
	return fmt.Sprintf("%+v\n", *v)
}

const utrudnieniaStr = "Utrudnienia w komunikacji linii: "

func getId(it *rss.Item) (int, error) {
	u, err := url.Parse(it.Link)
	if err != nil {
		return 0, err
	}
	idKeys, ok := u.Query()["i"]

	if !ok || len(idKeys) < 1 {
		return 0, EventParseError{it, "Problem getting id"}
	}
	idStr := idKeys[0]
	return strconv.Atoi(idStr)
}

func getContent(it *rss.Item) (string, error) {
	document, err := goquery.NewDocument(it.Link)
	if err != nil {
		return "", err
	}
	return document.Find("#PageContent").Text(), nil
}

func getLines(it *rss.Item) ([]string, error) {
	if !strings.HasPrefix(it.Title, utrudnieniaStr) {
		return []string{}, EventParseError{it, "Weird Title"}
	}
	linesStr := it.Title[len(utrudnieniaStr):]
	return strings.Split(linesStr, ", "), nil
}

func parseRssItem(it *rss.Item) (*ZTMevent, error) {
	log.Print("Link=", it.Link)
	log.Print("Title=", it.Title)
	log.Print("Summery=", it.Summary)
	defer log.Print("--------")
	id, err := getId(it)
	if err != nil {
		return nil, err
	}
	content, err := getContent(it)
	if err != nil {
		return nil, err
	}
	lines, err := getLines(it)
	if err != nil {
		return nil, err
	}

	newIt := ZTMevent{
		Id: id,
		Lines: lines,
		Title: it.Summary,
		Content: content,
	}
	return &newIt, nil
}

func parseRSSFeed(feed []*rss.Item) ([]*ZTMevent, error) {
	var newFeed []*ZTMevent
	for _, it := range feed {
		newIt, err := parseRssItem(it)
		if err != nil {
			return newFeed, err
		}
		newFeed = append(newFeed, newIt)
	}
	return newFeed, nil
}

func FetchZTMevents() ([]*ZTMevent, error) {
	basicFeed, err := fetchRSSFeed()
	if err != nil {
		return []*ZTMevent{}, err
	}
	return parseRSSFeed(basicFeed)
}
