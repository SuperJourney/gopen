配置文件路劲

~/.gopen.toml

swagger

# gopen
gopen api


https://github.com/SuperJourney/gopen

go get github.com/sashabaranov/go-openai


go install github.com/smallnest/gen@latest




["","","",""]

-- 公用提词器

text
文本

-- 请根据指定关键字进行联想，口红,特征： 女性，红色，给出指定标题， 标题长度5-10个字， 不要包含除了标题的其他内容

类型：标题
prompt拼接：

type : 1 
请根据相关提示进行联系，$1  给出指定 $2 ， $2 长度 $3 个字，不要包含除了 $2 的其他内容




```json
{
  "context": [
    {
      "content": "以下对话场景你需要充当一个商品设计小助手, 我会在下文提供少量商品的描述，请根据对话内容给出极具创造力，吸引力的标题,标题长度不得超过20个字",
      "role": "system"
    },
    {
      "content": "口红， 方便面色，涂上很好看",
      "role": "user"
    },
    {
      "content": "瞬间抢眼！方便面色轻抹口红",
      "role": "assistant"
    }
  ],
  "name": "商品标题",
  "type": 1
}
```