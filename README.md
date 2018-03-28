# bookmark-go

## Dev
init database:
```sh
go run script/initdb.go
```

start:
```sh
glide                               # 安装依赖
go run main.go -c="`pwd`/conf.json" # -c 指定配置文件路径
```

## ref
https://bulma.io/documentation/form/general/