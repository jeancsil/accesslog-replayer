package service

import (
	"fmt"

	"github.com/jeancsil/accesslog-replayer/internals/contentparser"
)

type Replay struct {
	ContentServer contentparser.ContentServer
}

func (r Replay) Run(log string) {
	a := r.ContentServer.Parse(log)

	fmt.Println("==========")
	fmt.Println("Method:", a.Method)
	fmt.Println("Date time:", a.Time)
	fmt.Println("URI:", a.URI)
	fmt.Println("")
	// fmt.Printf("Parsed line: %s", a.Original)
}
