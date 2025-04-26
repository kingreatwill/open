

## model context protocol
https://modelcontextprotocol.io/introduction

MCP 是一种开放协议，它标准化了应用程序向 LLM 提供上下文的方式。可以将 MCP 视为 AI 应用程序的 USB-C 端口。正如 USB-C 提供了一种将设备连接到各种外围设备和配件的标准化方式一样，MCP 提供了一种将 AI 模型连接到不同数据源和工具的标准化方式。

## Find Awesome MCP Servers and Clients
https://mcp.so/


### smithery
https://smithery.ai/

npx -y @smithery/cli install elasticsearch-mcp-server --client claude --key xxx
npx -y @smithery/cli install mysql-mcp-server --client cursor --key xxx
npx -y @smithery/cli@latest install @benborla29/mcp-server-mysql --client claude

### mysql
https://github.com/benborla/mcp-server-mysql
{
  "mcpServers": {
    "MySQL": {
      "command": "npx",
      "args": [
        "mcprunner",
        "MYSQL_HOST=127.0.0.1",
        "MYSQL_PORT=3306",
        "MYSQL_USER=root",
        "MYSQL_PASS=root",
        "MYSQL_DB=demostore",
        "ALLOW_INSERT_OPERATION=true",
        "ALLOW_UPDATE_OPERATION=true",
        "ALLOW_DELETE_OPERATION=false",
        "--",
        "npx",
        "-y",
        "@benborla29/mcp-server-mysql"
      ]
    }
  }
}