
## 读取文件
在 Golang 中，读取文件有四种方法，分别为：使用 ioutil.ReadFile 读取文件，使用 file.Read 读取文件，使用 bufio.NewReader 读取文件，使用 ioutil.ReadAll 读取文件。

## OpenFile
在 Golang 中，OpenFile 是一个更底层的打开 文件 的 函数，该函数可以使用指定的选项与指定的模式来打开文件。
如果文件打开成功，返回一个可用于 IO 的文件对象，如果打开失败，则返回一个 error 错误。
```go
func OpenFile(name string, flag int, perm FileMode) (*File, error) 
```
参数
* 参数	描述
* name	要打开的文件路径。
* flag	打开文件的方式，见下表。
* perm	打开文件的模式。

返回值
* 返回值	描述
* File	打开文件返回的文件句柄。
* error	打开失败，则返回错误信息。

OpenFile函数flag参数
* 打开方式	说明
* O_RDONLY	只读方式打开
* O_WRONLY	只写方式打开
* O_RDWR	读写方式打开
* O_APPEND	追加方式打开
* O_CREATE	不存在，则创建
* O_EXCL	如果文件存在，且标定了O_CREATE的话，则产生一个错误
* O_TRUNG	如果文件存在，且它成功地被打开为只写或读写方式，将其长度裁剪唯一。（覆盖）
* O_NOCTTY	如果文件名代表一个终端设备，则不把该设备设为调用进程的控制设备
* O_NONBLOCK	如果文件名代表一个FIFO,或一个块设备，字符设备文件，则在以后的文件及I/O操作中置为非阻塞模式。
* O_SYNC	当进行一系列写操作时，每次都要等待上次的I/O操作完成再进行。

## 写入文件
在 Golang 中，写 文件 有四种方法，分别为：使用 io.WriteString 写文件，使用 ioutil.WriteFile 写文件，使用 file.Write 写文件，使用 writer.WriteString 写文件。


