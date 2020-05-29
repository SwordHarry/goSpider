package engine

type ParserFunc func(contents []byte, url string) ParseResult

// 用于rpc：ParseFunc 需要序列化与反序列化
type Parser interface {
	Parse(contents []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}

type Request struct {
	Url string
	Parser
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

type NilParser struct {
}

func (n NilParser) Parse(contents []byte, url string) ParseResult {
	return ParseResult{}
}

func (n NilParser) Serialize() (name string, args interface{}) {
	return "NilParser", nil
}

// 包装 Parser
type FuncParser struct {
	FuncName   string
	ParserFunc ParserFunc
}

func (f *FuncParser) Parse(contents []byte, url string) ParseResult {
	return f.ParserFunc(contents, url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.FuncName, nil
}

func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{name, p}
}
