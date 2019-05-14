package contentparser

type Apache struct {
}

func (a Apache) Parse(line string) (*Access, error) {
	// log.Printf("Apache Replay %s", line)w
	// return *Access{Original: line, Time: }
	return &Access{Original: line}, nil
}
