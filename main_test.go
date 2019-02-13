package main

import (
	. "gopkg.in/check.v1"
	"strings"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestLogScanner(c *C) {
	c.Check(
		processLog(`84.242.208.111- - [11/May/2013:06:31:00 +0200] "POST /chat.php HTTP/1.1" 200 354 "http://bim-bom.ru/" "Mozilla/5.0 (Windows NT 6.1; rv:20.0) Gecko/20100101 Firefox/20.0"
91.224.96.130 - - [11/May/2013:04:09:02 -0700] "GET /mod.php HTTP/1.0" 301 - 0 12413 "http://wiki.org/index.php#lang=en" "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; WOW64; Trident/5.0)"
77.21.132.156 - - [24/May/2013:23:37:48 +0200] "POST /app/engine/api.php HTTP/1.1" 200 80 "http://lag.ru/index.php" "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.31 (KHTML, like Gecko) Chrome/26.0.1410.64 Safari/537.31"
77.21.132.156 - - [24/May/2013:23:37:48 +0200] "POST /app/modules/randomgallery.php HTTP/1.1" 200 46542 "http://lag.ru/index.php" "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.31 (KHTML, like Gecko) Chrome/26.0.1410.64 Safari/537.31"
77.21.132.156 - - [24/May/2013:23:37:50 +0200] "POST /chat.php?id=a65 HTTP/1.1" 200 366 "http://lag.ru/index.php" "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.31 (KHTML, like Gecko) Chrome/26.0.1410.64 Safari/537.31"
77.21.132.156 - - [24/May/2013:23:37:58 +0200] "-" 400 0 "-" "-"
77.21.132.156 - - [24/May/2013:23:37:58 +0200] "-" 400 0 "-" "-"
77.21.132.156 - - [24/May/2013:23:38:03 +0200] "POST /app/engine/api.php HTTP/1.1" 200 80 "http://lag.ru/index.php" "Mozilla/5.0 (Windows NT 6.1; WOW64) YandexBot/537.31 (KHTML, like Gecko) Chrome/26.0.1410.64 Safari/537.31"
77.21.132.156 - - [24/May/2013:23:37:58 +0200] "-" 400 0 "-" "-"
77.21.132.156 - - [24/May/2013:23:37:58 +0200] "-" 400 0 "-" "-"
77.21.132.156 - - [24/May/2013:23:38:03 +0200] "POST /app/engine/api.php HTTP/1.1" 200 80 "http://lag.ru/index.php" "Mozilla/5.0 (Windows NT 6.1; WOW64) Googlebot/537.31 (KHTML, like Gecko) Chrome/26.0.1410.64 Safari/537.31"
77.21.132.156 - - [24/May/2013:23:38:05 +0200] "POST /chat.php?id=a65 HTTP/1.1" 200 31 "http://lag.ru/index.php" "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.31 (KHTML, like Gecko) Chrome/26.0.1410.64 Safari/537.31"
77.21.132.156 - - [24/May/2013:23:38:23 +0200] "POST /app/modules/randomgallery.php HTTP/1.1" 200 46542 "http://lag.ru/index.php" "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.31 (KHTML, like Gecko) Chrome/26.0.1410.64 Safari/537.31"
84.242.208.111- - [11/May/2013:06:31:00 +0200] "POST /chat.php HTTP/1.1" 200 354 "http://bim-bom.ru/" "Mozilla/5.0 (Windows NT 6.1; rv:20.0) Gecko/20100101 Firefox/20.0"
91.234.91.31 - - [11/May/2013:04:09:02 -0700] "GET /mod.php HTTP/1.0" 301 - 0 12413 "http://wiki.org/index.php#lang=en" "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; WOW64; Trident/5.0)"
77.21.132.156 - - [24/May/2013:23:37:48 +0200] "POST /app/engine/api.php HTTP/1.1" 200 80 "http://lag.ru/index.php" "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.31 (KHTML, like Gecko) Chrome/26.0.1410.64 Safari/537.31"
77.21.132.156 - - [24/May/2013:23:37:48 +0200] "POST /app/modules/randomgallery.php HTTP/1.1" 200 46542 "http://lag.ru/index.php" "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.31 (KHTML, like Gecko) Chrome/26.0.1410.64 Safari/537.31"
77.21.132.156 - - [24/May/2013:23:37:50 +0200] "POST /chat.php?id=a65 HTTP/1.1" 200 366 "http://lag.ru/index.php" "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.31 (KHTML, like Gecko) Chrome/26.0.1410.64 Safari/537.31"
77.21.132.156 - - [24/May/2013:23:37:58 +0200] "-" 400 0 "-" "-"
77.21.132.156 - - [24/May/2013:23:37:58 +0200] "-" 400 0 "-" "-"
77.21.132.156 - - [24/May/2013:23:37:58 +0200] "-" 400 0 "-" "-"
77.21.132.156 - - [24/May/2013:23:37:58 +0200] "-" 400 0 "-" "-"
77.21.132.156 - - [24/May/2013:23:38:03 +0200] "POST /app/engine/api.php HTTP/1.1" 200 80 "http://lag.ru/index.php" "Mozilla/5.0 (Windows NT 6.1; WOW64) Googlebot/537.31 (KHTML, like Gecko) Chrome/26.0.1410.64 Safari/537.31"
77.21.132.156 - - [24/May/2013:23:38:05 +0200] "POST /chat.php?id=a65 HTTP/1.1" 200 31 "http://lag.ru/index.php" "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.31 (KHTML, like Gecko) Chrome/26.0.1410.64 Safari/537.31"
77.21.132.156 - - [24/May/2013:23:38:23 +0200] "POST /app/modules/randomgallery.php HTTP/1.1" 200 46542 "http://lag.ru/index.php" "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.31 (KHTML, like Gecko) Chrome/26.0.1410.64 Safari/537.31"`),
		Equals,
		LogStatistics{
			Views:    15,
			Visitors: 4,
			Urls:     6,
			Traffic:  188070,
			Lines:    25,
			Crawlers: CrawlersStatistics{
				Google: 2,
				Bing:   0,
				Yandex: 1,
			},
		},
	)
}

func processLog(log string) LogStatistics {
	logScanner := NewLogScanner(strings.NewReader(log))
	logScanner.ScanAll()
	return *logScanner.LogStatistics
}
