### 交叉编译
#### Mac 下编译 Linux 和 Windows 64位可执行程序
```go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o mtlogin
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o goutil.exe
```

#### Linux 下编译 Mac 和 Windows 64位可执行程序
```go
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o mtlogin
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o mtlogin.exe
```

#### Windows 下编译 Mac 和 Linux 64位可执行程序
```go
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build -o mtlogin

SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o mtlogin
```

### 运行
+ mtlogin 和 .env 文件放到服务器同一目录
+ `chmod 777 ./mtlogin`增加文件执行权限
+ 服务器执行命令`date`查看服务器的时间
+ `cd /home/mtlogin` 进入目录
+ `vim .env` 修改配置文件中的时间
+ `nohup ./mtlogin &`启动程序后台运行
+ 查看运行日志`tail -f nohup.out`
+ 查看进程运行 `ps aux | grep mt`
