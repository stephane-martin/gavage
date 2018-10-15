package main

import (
	"net"

	"cloud.google.com/go/civil"
)

type LogLine struct {
	Date          civil.Date
	Time          civil.Time
	TimeTaken     float64
	ClientIP      net.IP
	Username      string
	ExceptionID   string
	FilterResult  string
	Categories    []string
	Status        int64
	Action        string
	Method        string
	ContentType   string
	Scheme        string
	Host          string
	Port          int64
	Path          string
	Query         string
	Extension     string
	UserAgent     string
	SentBytes     int64
	ReceivedBytes int64
}

func (l *LogLine) toES() *ESLogLine {

}

type ESLogLine struct {
	Timestamp     string
	ClientIP      string
	SentBytes     int64
	ReceivedBytes int64
	Categories    string
	UserAgent     string
	Host          string
	Method        string
	Extension     string
	Path          string
	Port          int64
	Query         string
	Scheme        string
	Username      string
	Date          string
	DayOfWeek     int64
	ContentType   string
	Action        string
	FilterResult  string
	Status        int64
	Spamitude     float64
	TimeTaken     float64
	TLD           string
	TopDomain     string
	WorkDay       bool
	ExceptionID   string
	Year          int64
	YearMonth     string
}
