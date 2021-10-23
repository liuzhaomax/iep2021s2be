@cd "%~dp0%"
@cd..
:@cls
:@cmd
:整理检查依赖
@go mod tidy

:go mod init # 初始化当前目录为模块根目录，生成go.mod, go.sum文件
:go mod download # 下载依赖包
:go mod tidy #整理检查依赖，如果缺失包会下载或者引用的不需要的包会删除
:go mod vendor #复制依赖到vendor目录下面
:go mod 可看完整所有的命令
:go mod graph 以文本模式打印模块需求图
:go mod verify 验证依赖是否正确
:go mod edit 编辑go.mod文件