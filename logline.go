package main

import (
	"encoding/json"
	"math/rand"
	"strings"
	"time"

	"cloud.google.com/go/civil"
	"github.com/corpix/uarand"
)

type LogLine struct {
	Timestamp     time.Time
	TimeTaken     float64
	ClientIP      string
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
	Spamitude     float64
}

func ParseHost(host string) (topdomain string, tld string) {
	if len(host) == 0 {
		return "", ""
	}
	parts := strings.Split(host, ".")
	nbparts := len(parts)
	if nbparts < 2 {
		return host, ""
	}
	return parts[nbparts-2] + "." + parts[nbparts-1], parts[nbparts-1]
}

func (l *LogLine) Random(start, end time.Time, r *rand.Rand) *LogLine {
	if l == nil {
		l = new(LogLine)
	}
	if r == nil {
		r = rand.New(rand.NewSource(time.Now().UnixNano()))
	}
	uar := uarand.UARand{Randomizer: r}
	l.Timestamp = start.Add(time.Nanosecond * time.Duration(r.Int63n(end.Sub(start).Nanoseconds())))
	l.TimeTaken = r.ExpFloat64()
	l.ClientIP = FakeClientIP(r)
	l.Username = FakeUserName(r)
	l.ExceptionID = FakeExceptionID(r)
	l.FilterResult = FakeFilterResult(r)
	l.Categories = FakeCategories(r)
	l.Status = FakeStatus(r)
	l.Action = FakeAction(r)
	l.Method = FakeMethod(r)
	l.ContentType = FakeContentType(r)
	l.Scheme = FakeScheme(r)
	l.Host = FakeDomain(r)
	l.Port = 80
	if l.Scheme == "https" {
		l.Port = 443
	}
	l.Path = FakePath(r)
	l.Query = FakeQuery(r)
	l.Extension = FakeExtension(r)
	l.UserAgent = uar.GetRandom()
	l.SentBytes = FakeSentBytes(r)
	l.ReceivedBytes = FakeReceivedBytes(r)
	l.Spamitude = r.ExpFloat64() * 20
	return l
}

func (l *LogLine) ToES(res *ESLogLine) *ESLogLine {
	if l == nil {
		l = l.Random(time.Now().Add(-time.Hour), time.Now(), nil)
	}
	if res == nil {
		res = new(ESLogLine)
	}
	res.SentBytes = l.SentBytes
	res.ReceivedBytes = l.ReceivedBytes
	res.Categories = strings.Join(l.Categories, ";")
	res.ClientIP = l.ClientIP

	res.Scheme = l.Scheme
	res.Host = l.Host

	res.TopDomain, res.TLD = ParseHost(res.Host)
	res.Method = l.Method
	res.Extension = l.Extension
	res.Path = l.Path
	res.Port = l.Port
	res.Query = l.Query
	res.UserAgent = l.UserAgent
	res.ContentType = l.ContentType
	res.TimeTaken = l.TimeTaken

	res.Username = l.Username
	res.Action = l.Action
	res.FilterResult = l.FilterResult
	res.ExceptionID = l.ExceptionID
	res.Status = l.Status

	res.Spamitude = l.Spamitude

	res.Timestamp = l.Timestamp.Format(time.RFC3339)
	res.DayOfWeek = int64(l.Timestamp.Weekday()) + 1
	res.WorkDay = res.DayOfWeek != 6 && res.DayOfWeek != 7
	dt := civil.DateTimeOf(l.Timestamp)
	res.Date = dt.Date.String()
	res.Time = dt.Time.String()[0:8]
	res.Year = int64(dt.Date.Year)
	res.YearMonth = res.Date[0:7]
	return res
}

func (l *LogLine) toJSON() string {
	return l.ToES(nil).toJSON()
}

type ESLogLine struct {
	SentBytes     int64  `json:"cs-bytes"`
	ReceivedBytes int64  `json:"sc-bytes"`
	Categories    string `json:"cs-categories"`
	ClientIP      string `json:"c-ip"`

	Scheme      string  `json:"cs-uri-scheme"`
	Host        string  `json:"cs-host"`
	TLD         string  `json:"tld"`
	TopDomain   string  `json:"topdomain"`
	Method      string  `json:"cs-method"`
	Extension   string  `json:"cs-uri-extension"`
	Path        string  `json:"cs-uri-path"`
	Port        int64   `json:"cs-uri-port"`
	Query       string  `json:"cs-uri-query"`
	UserAgent   string  `json:"cs-h-user-agent"`
	ContentType string  `json:"rs-h-content-type"`
	TimeTaken   float64 `json:"time-taken"`

	Username     string `json:"cs-username"`
	Action       string `json:"s-action"`
	FilterResult string `json:"sc-filter-result"`
	ExceptionID  string `json:"x-exception-id"`
	Status       int64  `json:"sc-status"`

	Spamitude float64 `json:"spamitude"`

	WorkDay   bool   `json:"work"`
	Timestamp string `json:"@timestamp"`
	DayOfWeek int64  `json:"dayofweek"`
	Date      string `json:"date"`
	Time      string `json:"time"`
	Year      int64  `json:"year"`
	YearMonth string `json:"yearmonth"`
}

func (l *ESLogLine) toJSON() string {
	b, _ := json.Marshal(l)
	return string(b)
}
