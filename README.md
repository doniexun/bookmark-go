# bookmark-go

Restful API to manage bookmarks.

## Framework
- gin
- gorm

## Development
```sh
go run script/initdb.go # 初始化数据库(仅第一次运行)
glide install           # 安装依赖
go run main.go          # -c 指定配置文件路径
```

## Usage

Signup
```
POST /api/v1/user HTTP/1.1
Host: 127.0.0.1:3000
Content-Type: application/json
Cache-Control: no-cache
Postman-Token: 19e9ff02-afa3-9a04-0e5e-7bf9fd718561

{
	"mail": "bbb@a.a",
	"password": "bbccss",
	"tick": "Xv8Cm6KBrqYrNaBjGFF4",
	"captcha": "13878"
}
```