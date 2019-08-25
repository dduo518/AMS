### grpc 服务[read](https://github.com/chong0808/advanced-go-programming-book/blob/master/ch4-rpc/ch4-02-pb-intro.md)
>>> 用Protobuf定义语言无关的PRC服务接口才是它真正的价值所在！
#### 安装protobuf 工具

- [下载protobuf](https://github.com/protocolbuffers/protobuf/releases)
##### mac 全局安装
解压之后 到指定文件夹 $filepath/protobuf
```
cd $filepath/protobuf
$ cp -r bin/ /usr/local/bin/
$ cp -r include/ /usr/local/include/
// 测试
$ protoc
```


###### 使用protoc-gen-go内置的grpc插件生成GRPC代码：
```
//// protoc --go_out=plugins=grpc:{输出目录}  {proto文件}
$ protoc --go_out=plugins=grpc:. *.proto
```



##### 部署

- 交叉编译

    mac上编译 linux 可执行文件
````$xslt
     $ CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
````
   
- linux 压缩与解压
 ```$xslt
// 压缩
$ tar zcvf FileName.tar.gz DirName
// 解压
$ tar zxvf FileName.tar.gz
 
````

- linux 
```$xslt
// 查看进程
$ ps -ef|grep java|grep
```

- 上传文件到服务器

 ```$xslt
$ scp $fromDir/file $username@xx.xx.xx.xx:$toDir

// example 
 scp /Users/z_chong/go/src/vote_go/templates.tar.gz root@120.24.170.180:/root/vote

```

- 守护进程

1. supervisor 
```$xslt
// 启动 
$ sudo echo_supervisord_conf > /etc/supervisord.conf

// 编辑/etc/supervisord.conf，在最后增加运行程序设置
[program:golang-http-server] // 经名字
command=/home/golang/simple_http_server //启动命令 比如 root/AMS/build/AMS -config conf.prod.yaml
autostart=true
autorestart=true
startsecs=10
stdout_logfile=/var/log/simple_http_server.log // 日志输出文件
stdout_logfile_maxbytes=1MB
stdout_logfile_backups=10
stdout_capture_maxbytes=1MB
stderr_logfile=/var/log/simple_http_server.log
stderr_logfile_maxbytes=1MB
stderr_logfile_backups=10
stderr_capture_maxbytes=1MB
```
2.几个配置说明：
````$xslt
command：表示运行的命令，填入完整的路径即可。
autostart：表示是否跟随supervisor一起启动。
autorestart：如果该程序挂了，是否重新启动。
stdout_logfile：终端标准输出重定向文件。
stderr_logfile：终端错误输出重定向文件
````
tips：如果修改了配置文件，可以用kill -HUP重新加载配置文件
```$xslt
$ cat /tmp/supervisord.pid | xargs sudo kill -HUP
```
查看supervisor运行状态
```$xslt
supervisorctl
```