


https://github.com/ollama/ollama 93.7k

[Ollama](https://ollama.com/) 是一个开源工具，它允许用户在本地设备上轻松运行和自定义大型语言模型，例如[llama3](https://ollama.com/library/llama3)和[gemma2](https://ollama.com/library/gemma2) 和[Qwen2.5](https://ollama.com/library/qwen2.5)。

`ollama run gemma:2b`
Ollama启动服务: `ollama serve`, 会在本地的11434端口启动一个服务
```
curl http://localhost:11434/api/chat -d '{
  "model": "gemma:2b",
  "messages": [
    { "role": "user", "content": "你好" }
  ]
}'
```

详情接口: https://github.com/ollama/ollama/blob/main/docs/api.md#generate-a-chat-completion

### 翻译
https://github.com/UlionTse/translators

https://github.com/LibreTranslate/LibreTranslate