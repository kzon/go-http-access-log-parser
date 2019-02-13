package main

import (
	"regexp"
	"strconv"
	"time"
)

type Hit struct {
	IP            string
	Time          time.Time
	Status        int
	ContentLength int
	URL           string
	Method        string
	Referrer      string
	UserAgent     string
}

var logRecordRegexp = regexp.MustCompile(`^([\d.]+).*"(\S*) ?([\w\S]*).*" (\d+) ([\d\-]+).*"(.+)".*"(.+)"`)

func ParseLogLine(line string) Hit {
	hit := Hit{}
	match := logRecordRegexp.FindStringSubmatch(line)

	hit.IP = match[1]
	hit.Method = match[2]
	hit.URL = match[3]

	status, _ := strconv.Atoi(match[4])
	hit.Status = status

	contentLength, _ := strconv.Atoi(match[5])
	hit.ContentLength = contentLength

	hit.Referrer = match[6]
	hit.UserAgent = match[7]

	return hit
}
