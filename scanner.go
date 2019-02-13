package main

import (
	"bufio"
	"io"
)

type LogScanner struct {
	*LogProcessor
	*bufio.Scanner
}

func NewLogScanner(reader io.Reader) *LogScanner {
	return &LogScanner{
		NewLogStatisticsProcessor(),
		bufio.NewScanner(reader),
	}
}

func (scanner *LogScanner) ScanAll() {
	var hit Hit
	for scanner.Scan() {
		hit = ParseLogLine(scanner.Text())
		scanner.Process(hit)
	}
}