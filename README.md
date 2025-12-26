# Go HTTP Web Server

一个使用 Go 语言编写的简单 HTTP web 服务。

A simple HTTP web service written in Go.

## 功能特性 (Features)

- ✅ HTTP 服务器，监听端口 8080
- ✅ 多个 API 端点
- ✅ JSON 响应支持
- ✅ 请求日志记录
- ✅ 错误处理
- ✅ 健康检查端点

## 快速开始 (Quick Start)

### 运行服务器 (Run the server)

```bash
go run main.go
```

服务器将在 `http://localhost:8080` 启动。

### 构建可执行文件 (Build executable)

```bash
go build -o server main.go
./server
```

## API 端点 (API Endpoints)

### 1. 首页 (Homepage)
```
GET /
```
返回一个简单的 HTML 页面，显示所有可用的端点。

### 2. Hello API
```
GET /api/hello
```
返回问候消息的 JSON 响应。

**响应示例:**
```json
{
  "message": "Hello from Go HTTP Server! 你好！",
  "time": "2024-01-01T12:00:00Z"
}
```

### 3. 时间 API (Time API)
```
GET /api/time
```
返回当前服务器时间。

**响应示例:**
```json
{
  "current_time": "2024-01-01T12:00:00Z",
  "unix_time": 1704110400,
  "timezone": "UTC"
}
```

### 4. Echo API
```
POST /api/echo
Content-Type: application/json
```
回显你发送的 JSON 数据。

**请求示例:**
```json
{
  "name": "张三",
  "message": "Hello World"
}
```

**响应示例:**
```json
{
  "received": {
    "name": "张三",
    "message": "Hello World"
  },
  "time": "2024-01-01T12:00:00Z"
}
```

### 5. 健康检查 (Health Check)
```
GET /health
```
用于检查服务器是否正常运行。

**响应示例:**
```json
{
  "status": "healthy",
  "time": "2024-01-01T12:00:00Z"
}
```

## 测试 API (Testing the API)

使用 curl 命令测试：

```bash
# 测试 Hello API
curl http://localhost:8080/api/hello

# 测试时间 API
curl http://localhost:8080/api/time

# 测试 Echo API
curl -X POST http://localhost:8080/api/echo \
  -H "Content-Type: application/json" \
  -d '{"name":"测试","message":"你好"}'

# 测试健康检查
curl http://localhost:8080/health
```

## 项目结构 (Project Structure)

```
.
├── main.go           # 主程序文件
├── go.mod           # Go modules 配置
├── .gitignore       # Git 忽略文件
└── README.md        # 项目文档
```

## 技术栈 (Tech Stack)

- **语言**: Go 1.21+
- **标准库**: net/http, encoding/json

## 开发说明 (Development Notes)

该服务器包含以下特性：

1. **路由处理**: 使用 `http.ServeMux` 处理不同的路由
2. **中间件**: 实现了日志记录中间件
3. **JSON 处理**: 支持 JSON 请求和响应
4. **错误处理**: 统一的错误响应格式
5. **HTTP 方法验证**: 验证请求方法是否正确

## License

MIT
