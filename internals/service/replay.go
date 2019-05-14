package service

import (
	"log"
	"os"

	"github.com/jeancsil/accesslog-replayer/internals/contentparser"
)

type Replay struct {
	ContentServer contentparser.ContentServer
}

func (r Replay) Run(l string) {
	access, err := r.ContentServer.Parse(l)

	if err != nil {
		log.Fatal("error parsing the log")
	}

	lg := log.New(os.Stdout, "", 0)

	lg.Printf("%s\t%s %s", access.Time.Format("2006-01-02 15:04:05"), access.Method, access.URI)

	c := make(chan *contentparser.Access)

	go sendRequest(access, c)
	// go func() {
	// 	c <- access
	// }()
}

func sendRequest(a *contentparser.Access, c chan *contentparser.Access) {
	log.Println("Sending request to URI:", a.URI)
	c <- a
}
