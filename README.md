# bookmark-go

Restful API to manage bookmarks.

## Framework
- gin
- gorm

## Development
```sh
git clone https://github.com/GallenHu/bookmark-go.git bookmarkgo # use folder name "bookmarkgo"
go run script/initdb.go # 初始化数据库(仅第一次运行)
glide install           # 安装依赖
# 根据 setting.go 设置环境变量
export CONFIG_DBPWD=123456
export CONFIG_DBNAME=bookmark
go run main.go          # -c 指定配置文件路径
```

## Run with Docker
```sh
docker run -d --name bookmarkgo \
    --link mysql5.7:mysql \
    --link my-redis:redis \
    -p 3001:3001 \
    -e CONFIG_DBHOST=mysql5.7\
    -e CONFIG_REDISHOST=my-redis\
    -d bookmarkgo:v1
```

## Usage

Captcha
```
GET /api/v1/captcha HTTP/1.1
Host: 127.0.0.1:3000
Cache-Control: no-cache
Postman-Token: 384472ef-36af-212d-631c-ae3912dee06c
```

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
