package contentparser

import (
	"log"
)

type Nginx struct {
}

func (n Nginx) Parse(line string) Access {
	log.Printf("NGINX replay %s", line)

	// parse and create the Access obj
	return *NewAccess(line)
}
