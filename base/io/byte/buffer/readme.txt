bytes.Buffer是一个缓冲byte类型的缓冲器,里面存放着都是byte。
Buffer是一个变长的缓冲器，具有 Read 和Write方法。Buffer的零值是一个空的 buffer，但是可以使用
四种方式创建Buffer缓冲器：
======
var b bytes.Buffer  //直接定义一个 Buffer 变量，而不用初始化
b1 := new(bytes.Buffer)   //直接使用 new 初始化，可以直接使用
// 其它两种定义方式
func NewBuffer(buf []byte) *Buffer
func NewBufferString(s string) *Buffer
======
















