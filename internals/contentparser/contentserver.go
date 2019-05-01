package contentparser

import "time"

type ContentServer interface {
	Parse(string)
}

type parsedData struct {
	Time time.Time
	URL  string
	Type string
}
