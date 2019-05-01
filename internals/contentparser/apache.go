package contentparser

import "log"

type Apache struct {
}

func (a Apache) Parse(line string) {
	log.Printf("Apache Replay %s", line)
}
