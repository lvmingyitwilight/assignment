socket粘包的解包方式：

1. fix length：按照固定长度进行解码，缺点是消息长度不足的时候需要补齐，存在浪费空间的情况
2. delimiter based：根据起始和终止分隔符进行解码，内容中如果出现分隔符自身则需要转义
3. length field based frame decoder：根据固定表示长度的字段来获取单个包的长度，然后对后续内容进行解析