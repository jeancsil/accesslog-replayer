package service

import "github.com/jeancsil/accesslog-replayer/internals/contentparser"

type Replay struct {
	ContentServer contentparser.ContentServer
}

func (r Replay) Run(log string) {
	r.ContentServer.Parse(log)
}
