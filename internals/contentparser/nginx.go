package contentparser

import (
	"regexp"
	"time"
)

type Nginx struct {
}

func (n Nginx) Parse(line string) (*Access, error) {
	r := regexp.MustCompile(`(.+) .{1,2} .{1,2} .(?P<datetime>\d{2}\/[a-z A-Z]{3}\/\d{4}:\d{2}:\d{2}:\d{2}) (\+\d{4}). \"(GET|POST|HEAD|OPTIONS|DELETE|PUT|TRACE) (?P<url>.+)(http|HTTP\/1\.1|HTTP\/1\.0)"`)
	m := r.FindStringSubmatch(line)

	t, _ := time.Parse("02/Jan/2006:15:04:05", m[2])
	method := m[4]
	uri := m[5]

	return &Access{Original: line, Time: t, Method: method, URI: uri}, nil
}
