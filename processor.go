package main

import (
	"strings"
)

type LogProcessor struct {
	*LogStatistics
	uniqueURLs map[string]bool
	uniqueIPs  map[string]bool
}

func NewLogStatisticsProcessor() *LogProcessor {
	return &LogProcessor{
		&LogStatistics{},
		make(map[string]bool),
		make(map[string]bool),
	}
}

func (processor *LogProcessor) Process(hit Hit) {
	processor.Lines++
	if hit.Status >= 200 && hit.Status < 300 {
		processor.Views++
	}
	processor.Traffic += hit.ContentLength
	if _, exists := processor.uniqueIPs[hit.IP]; !exists {
		processor.uniqueIPs[hit.IP] = true
		processor.Visitors++
	}
	if _, exists := processor.uniqueURLs[hit.URL]; !exists {
		processor.uniqueURLs[hit.URL] = true
		processor.Urls++
	}

	if isGoogleCrawler(hit.UserAgent) {
		processor.Crawlers.Google++
	} else if isBingCrawler(hit.UserAgent) {
		processor.Crawlers.Bing++
	} else if isYandexCrawler(hit.UserAgent) {
		processor.Crawlers.Yandex++
	}
}

func isGoogleCrawler(userAgent string) bool {
	return strings.Contains(userAgent, "Googlebot")
}

func isBingCrawler(userAgent string) bool {
	return strings.Contains(userAgent, "msnbot")
}

func isYandexCrawler(userAgent string) bool {
	return strings.Contains(userAgent, "YandexBot")
}
