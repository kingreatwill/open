
## 日志收集组件
[LoongCollector](https://github.com/alibaba/loongcollector)	
FluentBit	
Vector	
Filebeat

### 日志分析组件(可以在控制台展示分析日志, 还可以AI分析)
go install github.com/control-theory/gonzo/cmd/gonzo@latest

```
# Read logs directly from files
gonzo -f application.log

# Read from multiple files
gonzo -f application.log -f error.log -f debug.log

# Use glob patterns to read multiple files
gonzo -f "/var/log/*.log"
gonzo -f "/var/log/app/*.log" -f "/var/log/nginx/*.log"

# Follow log files in real-time (like tail -f)
gonzo -f /var/log/app.log --follow
gonzo -f "/var/log/*.log" --follow

# Analyze logs from stdin (traditional way)
cat application.log | gonzo

# Stream logs from kubectl
kubectl logs -f deployment/my-app | gonzo

# Follow system logs
tail -f /var/log/syslog | gonzo

# Analyze Docker container logs
docker logs -f my-container 2>&1 | gonzo

# With AI analysis (requires API key)
export OPENAI_API_KEY=sk-your-key-here
gonzo -f application.log --ai-model="gpt-4"
```