package engine

type Request struct {
	Url string
	ParserFunc
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url     string
	Id      string
	Type    string
	Payload interface{} // data
}

type ParserFunc func(contents []byte, url string) ParseResult

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
