package contentparser

type ContentServer interface {
	Parse(string) (*Access, error)
}
