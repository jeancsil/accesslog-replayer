package contentparser

import "log"

type Nginx struct {
}

func (a Nginx) Parse(line string) {
	log.Printf("NGINX replay %s", line)
}
