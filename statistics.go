package main

import "encoding/json"

type LogStatistics struct {
	Views    int
	Visitors int
	Urls     int
	Traffic  int
	Lines    int
	Crawlers CrawlersStatistics
}

type CrawlersStatistics struct {
	Google int
	Bing   int
	Yandex int
}

func (stat *LogStatistics) String() string {
	statJson, err := json.Marshal(stat)
	FatalIfError(err)
	return string(statJson)
}

