# Wenwo RPC 文档

版本: v1.0  
时间: 2021/12/18

## 1. 概要

本文档描述了 Wenwo 各个系统间通讯的协议和数据格式。

## 2. 协议
* GRPC
* Proto 3.0
* 不开启 TLS

## 3. 地址
http://60.205.1.144:8081/git/gingernet/wenworpc.git

## 4. 错误约定
### 数据结构:
```protobuf
message Error {
  int32  code      = 1; // 0: 成功; 404: NotFound; : 失败;
  string brief     = 2;
  string detail    = 3;
  bool   can_retry = 4; // true: 可以重试，false: 不能, 在写入失败的时候，该标识位提示客户端来决定是否可以重试。
}
```
### 目前定义了三种状态码，
* 0: 成功;
* 404: NotFound;
* 1: 其他失败;

### 关于字段 can_trey: 
表示该 API 是否可以重试。   
在写入失败的时候，该标识位提示客户端来决定是否可以重试。  
在某些特殊的 API，重试可能会导致问题的，返回 false。 

### 例子: 通过 id 查地址:
```protobuf
service EsLook {
  rpc getBlogs(BlogEsRequest) returns (BlogEsResponse) {}
}
```

如果该地址不存在 响应值为(采用 JSON 格式描述)
```json
{
  "code": 404,
  "brief": "Not found user!",
  "detail": "",
  "can_retry": false
}
```

## 其他公共的数据结构定义
返回空的时候采用自定义的 Empty。  
常见的 Id 请求和返回格式：IdRequest, IdResponse, IdsResponse  
更多请参考 wenworpc/btc.proto
```protobuf
import "wenworpc/common.proto";
option go_package = "git.wenwo.com/wenwo/wenworpc/go-wenworpc/scrapy";
option java_package = "com.wenworpc.scrapy";
package wenworpc.scrapy;

message ArticleScrapy {
  string title = 1;
  string author = 2;
  string platform = 3;
  string url = 4;
  string abstruct = 5;
  string content = 6;
}

message ArticleScrapyRequest{
  string consumer_token = 1;
  uint32 page = 2;
  uint32 pagesize = 3;
}

message ArticleScrapyResponse{
  Error error = 1;
  ArticleScrapy aes = 2;
}
```

## 5. 业务监听的端口(请自行更新)


[完]
