package model

type SearchResult struct {
	Hits     int64
	Start    int
	Items    []interface{}
	Query    string // 查询字符串
	PrevFrom int    // 上一页
	NextFrom int    // 下一页
}
