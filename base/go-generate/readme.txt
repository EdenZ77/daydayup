背景:
大家经常碰到命名错误码、状态码的同时，又要同步写码对应的翻译，有没有感觉很无聊。这里举一个例子
// 定义错误码
const (
   ERR_CODE_OK             = 0 // OK
   ERR_CODE_INVALID_PARAMS = 1 // 无效参数
   ERR_CODE_TIMEOUT        = 2 // 超时
)
// 定义错误码与描述信息的映射
var mapErrDesc = map[int]string{
   ERR_CODE_OK:             "OK",
   ERR_CODE_INVALID_PARAMS: "无效参数",
   ERR_CODE_TIMEOUT:        "超时",
}
// 根据错误码返回描述信息
func GetDescription(errCode int) string {
   if desc, exist := mapErrDesc[errCode]; exist {
      return desc
   }
   return fmt.Sprintf("error code: %d", errCode)
}
func main() {
   fmt.Println(GetDescription(ERR_CODE_OK))
}
=====================================================
这是一种重复性操作，没有什么技术含量，另外很可能忘记写映射。我只想写错误码，对应的描述信息直接用注释里的就行，所以这里介绍一下对应的工具。

go generate是 Go 自带的工具，使用命令go generate执行，go generate是利用源代码中的注释工作的。

stringer不是Go自带工具，需要手动安装。执行如下命令即可
go get golang.org/x/tools/cmd/stringer

使用：
    一种是在errcode中，第一行增加注释//go:generate stringer -type ErrCode -linecomment
    另一种是直接命令行执行stringer -type ErrCode -linecomment
执行完毕会发现自动生成新文件
