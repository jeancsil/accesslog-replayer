package contentparser

import "time"

type Access struct {
	Time     time.Time
	URL      string
	Type     string
	Original string
}

func NewAccess(line string) *Access {
	// parse line
	a := new(Access)
	a.Original = line
	// a.Time = ??;
	a.Type = "POST"
	a.URL = "http://www.trivago.com.br/"
	return a
}
