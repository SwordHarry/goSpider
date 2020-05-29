# goSpider
使用纯 go 进行爬虫，没有使用其余第三方框架，使用正则表达式进行页面信息爬取（后期可以考虑 goquery）
## 1. 依赖环境
#### go v 1.14.2
#### docker v19.03.8
#### docker -> elasticsearch v 5.x

## 2. 业务

刚开始爬的珍爱网 zhenai，但是被防，换成 cncn 美食网，被限流，再换成猫眼电影期待电影榜和古诗文网，最终使用古诗文网(`gushiwen.org`)进行爬虫练习

## 3. 架构

### 架构图

![goSpider](./goSpider.png)

### engine
总控节点

### itemSaver
为调用 ElasticSearch 的客户端，进行爬虫数据存储，TODO：去重部分可以做成去重服务

### worker
由 fetcher 和 parser 组成
#### parser
需要进行 序列化 和 反序列化 的包装操作

### 分布式微服务，jsonrpc
使用 jsonrpc 将各模块分装成服务，再利用


## 4. 进一步的工作
- 爬取更多网站，使用 css 选择器
- 对抗反爬技术/遵循 robots 协议
- 模拟登录，爬取动态网页
- 分布式去重
- 优化ElasticSearch查询质量，中文分词
- 使用脚本部署；使用 docker + k8s 部署
- 集成服务发现框架 consul
- 用 logstash 汇总和分析日志

