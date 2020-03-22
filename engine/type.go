package engine

type ParseResult struct {
	Request []Request
	Item []interface{}
}

type Request struct {
	Url string
	ParseFunc func([]byte)	ParseResult
}

func NilParse ([]byte) ParseResult{
	return ParseResult{}
}