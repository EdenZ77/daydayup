
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

## bufio
### 基本介绍
io操作本身的效率并不低，低的是频繁的访问本地磁盘的文件。所以bufio就提供了缓冲区(分配一块内存)，读和写都先在缓冲区中，最后再读写文件，来降低访问本地磁盘的次数，从而提高效率。
![img.png](img.png)
bufio.Reader 是bufio中对io.Reader 的封装
```go
// Reader implements buffering for an io.Reader object.
type Reader struct {
    buf          []byte
    rd           io.Reader // reader provided by the client
    r, w         int       // buf read and write positions
    err          error
    lastByte     int // last byte read for UnreadByte; -1 means invalid
    lastRuneSize int // size of last rune read for UnreadRune; -1 means invalid
}
```
bufio.Read(p []byte) 相当于读取大小len(p)的内容，思路如下：

1. 当缓存区有内容的时，将缓存区内容全部填入p并清空缓存区
2. 当缓存区没有内容的时候且len(p)>len(buf),即要读取的内容比缓存区还要大，直接去文件读取即可
3. 当缓存区没有内容的时候且len(p)<len(buf),即要读取的内容比缓存区小，缓存区从文件读取内容充满缓存区，并将p填满（此时缓存区有剩余内容）
4. 以后再次读取时缓存区有内容，将缓存区内容全部填入p并清空缓存区（此时和情况1一样）

源码
```go
// Read reads data into p.
// It returns the number of bytes read into p.
// The bytes are taken from at most one Read on the underlying Reader,
// hence n may be less than len(p).
// To read exactly len(p) bytes, use io.ReadFull(b, p).
// At EOF, the count will be zero and err will be io.EOF.
func (b *Reader) Read(p []byte) (n int, err error) {
    n = len(p)
    if n == 0 {
        return 0, b.readErr()
    }
    if b.r == b.w {
        if b.err != nil {
            return 0, b.readErr()
        }
        if len(p) >= len(b.buf) {
            // Large read, empty buffer.
            // Read directly into p to avoid copy.
            n, b.err = b.rd.Read(p)
            if n < 0 {
                panic(errNegativeRead)
            }
            if n > 0 {
                b.lastByte = int(p[n-1])
                b.lastRuneSize = -1
            }
            return n, b.readErr()
        }
        // One read.
        // Do not use b.fill, which will loop.
        b.r = 0
        b.w = 0
        n, b.err = b.rd.Read(b.buf)
        if n < 0 {
            panic(errNegativeRead)
        }
        if n == 0 {
            return 0, b.readErr()
        }
        b.w += n
    }
	
    // copy as much as we can
    n = copy(p, b.buf[b.r:b.w])
    b.r += n
    b.lastByte = int(b.buf[b.r-1])
    b.lastRuneSize = -1
    return n, nil
}
```
reader内部通过维护一个r, w 即读入和写入的位置索引来判断是否缓存区内容被全部读出。

bufio.Writer 是bufio中对io.Writer 的封装
```go
// Writer implements buffering for an io.Writer object.
// If an error occurs writing to a Writer, no more data will be
// accepted and all subsequent writes, and Flush, will return the error.
// After all data has been written, the client should call the
// Flush method to guarantee all data has been forwarded to
// the underlying io.Writer.
type Writer struct {
    err error
    buf []byte
    n   int
    wr  io.Writer
}
```
bufio.Write(p []byte) 的思路如下

1. 判断buf中可用容量是否可以放下 p
2. 如果能放下，直接把p拼接到buf后面，即把内容放到缓冲区
3. 如果缓冲区的可用容量不足以放下，且此时缓冲区是空的，直接把p写入文件即可
4. 如果缓冲区的可用容量不足以放下，且此时缓冲区有内容，则用p把缓冲区填满，把缓冲区所有内容写入文件，并清空缓冲区
5. 判断p的剩余内容大小能否放到缓冲区，如果能放下（此时和步骤1情况一样）则把内容放到缓冲区
6. 如果p的剩余内容依旧大于缓冲区，（注意此时缓冲区是空的，情况和步骤3一样）则把p的剩余内容直接写入文件

以下是源码
```go
// Write writes the contents of p into the buffer.
// It returns the number of bytes written.
// If nn < len(p), it also returns an error explaining
// why the write is short.
func (b *Writer) Write(p []byte) (nn int, err error) {
    for len(p) > b.Available() && b.err == nil {
        var n int
        if b.Buffered() == 0 {
            // Large write, empty buffer.
            // Write directly from p to avoid copy.
            n, b.err = b.wr.Write(p)
        } else {
            n = copy(b.buf[b.n:], p)
            b.n += n
            b.Flush()
        }
        nn += n
        p = p[n:]
    }
    if b.err != nil {
        return nn, b.err
    }
    n := copy(b.buf[b.n:], p)
    b.n += n
    nn += n
    return nn, nil
}
```
说明：

b.wr 存储的是一个io.writer对象，实现了Write()的接口，所以可以使用b.wr.Write(p) 将p的内容写入文件。

b.flush() 会将缓存区内容写入文件，当所有写入完成后，因为缓存区会存储内容，所以需要手动flush()到文件。

b.Available() 为buf可用容量，等于len(buf) - n。

### 常用方法和函数
```go
// NewReaderSize 将 rd 封装成一个带缓存的 bufio.Reader 对象，
// 缓存大小由 size 指定（如果小于 16 则会被设置为 16）。
// 如果 rd 的基类型就是有足够缓存的 bufio.Reader 类型，则直接将
// rd 转换为基类型返回。
func NewReaderSize(rd io.Reader, size int) *Reader

// NewReader 相当于 NewReaderSize(rd, 4096)
func NewReader(rd io.Reader) *Reader

// Peek 返回缓存的一个切片，该切片引用缓存中前 n 个字节的数据，
// 该操作不会将数据读出，只是引用，引用的数据在下一次读取操作之
// 前是有效的。如果切片长度小于 n，则返回一个错误信息说明原因。
// 如果 n 大于缓存的总大小，则返回 ErrBufferFull。
func (b *Reader) Peek(n int) ([]byte, error)

// Read 从 b 中读出数据到 p 中，返回读出的字节数和遇到的错误。
// 如果缓存不为空，则只能读出缓存中的数据，不会从底层 io.Reader
// 中提取数据，如果缓存为空，则：
// 1、len(p) >= 缓存大小，则跳过缓存，直接从底层 io.Reader 中读
// 出到 p 中。
// 2、len(p) < 缓存大小，则先将数据从底层 io.Reader 中读取到缓存
// 中，再从缓存读取到 p 中。
func (b *Reader) Read(p []byte) (n int, err error)

// Buffered 返回缓存中未读取的数据的长度。
func (b *Reader) Buffered() int

// ReadBytes 功能同 ReadSlice，只不过返回的是缓存的拷贝。
func (b *Reader) ReadBytes(delim byte) (line []byte, err error)

// ReadString 功能同 ReadBytes，只不过返回的是字符串。
func (b *Reader) ReadString(delim byte) (line string, err error)
```

```go
// NewWriterSize 将 wr 封装成一个带缓存的 bufio.Writer 对象，
// 缓存大小由 size 指定（如果小于 4096 则会被设置为 4096）。
// 如果 wr 的基类型就是有足够缓存的 bufio.Writer 类型，则直接将
// wr 转换为基类型返回。
func NewWriterSize(wr io.Writer, size int) *Writer

// NewWriter 相当于 NewWriterSize(wr, 4096)
func NewWriter(wr io.Writer) *Writer

// WriteString 功能同 Write，只不过写入的是字符串
func (b *Writer) WriteString(s string) (int, error)

// WriteRune 向 b 写入 r 的 UTF-8 编码，返回 r 的编码长度。
func (b *Writer) WriteRune(r rune) (size int, err error)

// Flush 将缓存中的数据提交到底层的 io.Writer 中
func (b *Writer) Flush() error

// Available 返回缓存中未使用的空间的长度
func (b *Writer) Available() int

// Buffered 返回缓存中未提交的数据的长度
func (b *Writer) Buffered() int

// Reset 将 b 的底层 Writer 重新指定为 w，同时丢弃缓存中的所有数据，复位
// 所有标记和错误信息。相当于创建了一个新的 bufio.Writer。
func (b *Writer) Reset(w io.Writer)
```

### ReadSlice、ReadBytes、ReadString 和 ReadLine 方法
参考资料：https://books.studygolang.com/The-Golang-Standard-Library-by-Example/chapter01/01.4.html

之所以将这几个方法放在一起，是因为他们有着类似的行为。事实上，后三个方法最终都是调用ReadSlice来实现的。所以，我们先来看看ReadSlice方法。(感觉这一段直接看源码较好)
ReadSlice方法签名如下：
```go
func (b *Reader) ReadSlice(delim byte) (line []byte, err error)
```

