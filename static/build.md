Mac 下编译 Linux 和 Windows 64位可执行程序

    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
    CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go

Linux 下编译 Mac 和 Windows 64位可执行程序

    CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go
    CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go

Windows 下编译 Mac 和 Linux 64位可执行程序

    SET CGO_ENABLED=0
    SET GOOS=darwin
    SET GOARCH=amd64
    go build main.go
    
    SET CGO_ENABLED=0
    SET GOOS=linux
    SET GOARCH=amd64
    go build main.go

    gox -osarch="linux/amd64"

pprof查看性能

    http://127.0.0.1:11451/debug/pprof/
    # 下载goroutine分析文件
    127.0.0.1:11451/debug/pprof/goroutine
    go tool pprof -http=":8081" goroutine
    # 下载内存分析文件
    127.0.0.1:11451/debug/pprof/heap
    go tool pprof -http=":8081" heap
    # 下载CPU分析文件
    http://127.0.0.1:11451/debug/pprof/profile?seconds=5 # 采样5s之内的cpu使用情况。
    go tool pprof -http=":8081" profile
    # 跟踪当前程序执行
    http://127.0.0.1:11451/debug/pprof/trace?seconds=5
    go tool trace -http=":8081" trace
    