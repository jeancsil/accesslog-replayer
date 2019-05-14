package contentparser

import "time"

type Access struct {
	Time     time.Time
	URI      string
	Method   string
	Original string
}

// func NewAccess(line string) *Access {
// 	// parse line
// 	a := new(Access)
// 	a.Original = line
// 	// a.Time = ??;
// 	a.Method = "POST"
// 	a.URI = "/login"
// 	return a
// }
