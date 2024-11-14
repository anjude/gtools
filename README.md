### GTools
gtools，为你的terminal（终端）插上翅膀~

### 1. 简单使用
1. 准备好golang环境，go bin放到环境变量下
2. 执行安装命令
```shell
# 安装到go bin目录
go install github.com/anjude/gtools@lastest

# 执行智能对话
gtools -c "你好"

# 查看帮助
gtools -h
```

![img.png](doc/img.png)

### 2. 开发
1. 在handler目录新增一个自定义目录，可以参考version目录和version_handler.go文件
2. 新增的handler实现base.IHandler接口
3. 在handler/option_map.go文件定义命令和handler的关联
4. 重新编译文件（go build -o gtools.exe main.go）即可
```go
type IHandler interface {
	// GetCommand 获取handler的命令
	GetCommand() (command string)
	// GetDesc 获取该命令的详情
	GetDesc() (desc string)
	// GetArgs 获取该命令需要的参数
	GetArgs(args []string) (curArgs []string, nextArgs []string)
	// Handle 该命令的具体执行逻辑
	Handle(args []string)
}
```
### 3. 联系作者
公众号：豆小匠Coding

![](doc/qrcode_for_coder_bean.jpg)