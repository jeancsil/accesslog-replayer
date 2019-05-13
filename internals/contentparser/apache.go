package contentparser

type Apache struct {
}

func (a Apache) Parse(line string) Access {
	// log.Printf("Apache Replay %s", line)w
	// return *Access{Original: line, Time: }
	return *NewAccess(line)
}
