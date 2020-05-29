package config

const (
	// itemSaver service port
	ItemSaverPort = 1234
	WorkerPort0   = 9000

	// ElasticSearch index
	ElasticIndexForMaoyan   = "maoyan"
	ElasticIndexForCncn     = "cncn"
	ElastciIndexForGushiwen = "gushiwen"
	// ElasticSearch type
	ElasticType = "movie"

	// ParserName
	// maoyan
	ParseMovie = "ParseMovie"
	// cncn
	ParseCityList = "ParseCityList"
	ParseCity     = "ParseCity"
	ParseFood     = "ParseFood"
	ParserStore   = "ParserStore"
	NilParser     = "NilParser"

	// gushiwen
	ParseThemeList = "ParseThemeList"
	ParseSentence  = "ParseSentence"

	// service Name
	ItemSaverRpc    = "ItemSaverService.Save"
	CrawlServiceRpc = "CrawlService.Process"

	// rate limit
	Qps = 10
)
