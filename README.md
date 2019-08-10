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
