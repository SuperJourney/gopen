配置文件路劲

~/.gopen.toml



# gopen
gopen api


https://github.com/SuperJourney/gopen

go get github.com/sashabaranov/go-openai


go install github.com/smallnest/gen@latest

sqlite3 



```sqlite

```


// 
## DDL

-- 创建应用表
CREATE TABLE app (
    app_id INTEGER PRIMARY KEY AUTOINCREMENT, -- 应用ID，自增长整数，作为主键
    name TEXT NOT NULL DEFAULT '' -- 应用名称，文本类型，不允许为空，设置默认值为空字符串
);

-- 创建应用属性表
CREATE TABLE app_property (
    title TEXT NOT NULL DEFAULT '', -- 应用属性标题，文本类型，不允许为空，设置默认值为空字符串
    usage TEXT NOT NULL DEFAULT '', -- 应用属性用途，文本类型，不允许为空，设置默认值为空字符串
    prompt INTEGER NOT NULL, -- 推荐词ID，整数类型，不允许为空
    create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- 创建时间，时间戳类型，设置默认值为当前时间
    FOREIGN KEY (recommended_word_id) REFERENCES recommended_word (recommended_word_id) -- 外键关联推荐词表
);

-- 创建推荐词表
CREATE TABLE prompt (
    prompt_id INTEGER PRIMARY KEY AUTOINCREMENT, -- 推荐词ID，自增长整数，作为主键
    type int
    content TEXT NOT NULL DEFAULT '' -- 推荐词内容，文本类型，不允许为空，设置默认值为空字符串 
);


["","","",""]

-- 公用提词器

text
文本

-- 请根据指定关键字进行联想，口红,特征： 女性，红色，给出指定标题， 标题长度5-10个字， 不要包含除了标题的其他内容

类型：标题
prompt拼接：

type : 1 
请根据相关提示进行联系，$1  给出指定 $2 ， $2 长度 $3 个字，不要包含除了 $2 的其他内容




关于chatgpt的研究： 

The following is a conversation with an AI assistant. The assistant is helpful, creative, clever, and very friendly


```
# Note: you need to be using OpenAI Python v0.27.0 for the code below to work
import openai

openai.ChatCompletion.create(
  model="gpt-3.5-turbo",
  messages=[
        {"role": "system", "content": "You are a helpful assistant."},
        {"role": "user", "content": "Who won the world series in 2020?"},
        {"role": "assistant", "content": "The Los Angeles Dodgers won the World Series in 2020."},
        {"role": "user", "content": "Where was it played?"}
    ]
)
```