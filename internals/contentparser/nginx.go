package contentparser

import (
	"log"
	"regexp"
	"strings"
	"time"
)

type Nginx struct {
}

func (n Nginx) Parse(line string) Access {
	// log.Printf("NGINX replay %s", line)

	// parse and create the Access obj
	// 127.0.0.1 - - [16/May/2018:02:20:37 +0000] "COPY /source/copy.txt HTTP/1.1" 204 0 "-" "curl/7.29.0" "-"

	var logRe = regexp.MustCompile(
		`^(?:(\S+(?:,\s\S+)*)\s)?` + // IP
			`(\S+)\s(\S+)\s` +
			`\[` +
			`((\d){2}/\w{3}/(\d){4}:(\d){2}:(\d){2}:(\d){2}\s)` + // Date
			`((.){5})` +
			`\]` +
			`\s` + // Whitespace char
			`(\S+)` + // Method
			`\s` + // Whitespace char
			`(\S+)` + // URI
			`\s` + // Whitespace char
			`(\S+)` + // HTTP Version
			`\s` + // Whitespace char
			`(\S+)` + // HTTP Response Code
			`\s` + // Whitespace char
			`(\S+)` + // Number of bytes ???
			`\s` + // Whitespace char
			`(\S+)` + // ?? "-"
			`\s` + // Whitespace char
			`(\S+)` + // User Agent
			`\s` + // Whitespace char
			`(\S+)`) // ?? "-"

	matches := logRe.FindStringSubmatch(line)
	if len(matches) < 1 {
		log.Fatal("failed to parse nginx (not matched): ", line)
	}

	// for index := 0; index < len(matches); index++ {
	// 	log.Println("", index, matches[index])
	// }

	time, _ := time.Parse("02/Jan/2006:15:04:05", strings.TrimSpace(matches[4]))
	method := strings.TrimLeft(matches[12], "\"")
	uri := strings.TrimSpace(matches[13])

	return Access{Original: line, Time: time, Method: method, URI: uri}
}
