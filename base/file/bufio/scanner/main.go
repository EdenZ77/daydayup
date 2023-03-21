package main

/*
参考资料：https://books.studygolang.com/The-Golang-Standard-Library-by-Example/chapter01/01.4.html#142-scanner-%E7%B1%BB%E5%9E%8B%E5%92%8C%E6%96%B9%E6%B3%95

对于简单的读取一行，在 Reader 类型中，感觉没有让人特别满意的方法。于是，Go1.1增加了一个类型：Scanner。官方关于Go1.1增加该类型的说明如下：
在 bufio 包中有多种方式获取文本输入，ReadBytes、ReadString 和独特的 ReadLine，对于简单的目的这些都有些过于复杂了。
在 Go 1.1 中，添加了一个新类型，Scanner，以便更容易的处理如按行读取输入序列或空格分隔单词等，这类简单的任务。
它终结了如输入一个很长的有问题的行这样的输入错误，并且提供了简单的默认行为：基于行的输入，每行都剔除分隔标识。这里的代码展示一次输入一行：
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        fmt.Println(scanner.Text()) // Println will add back the final '\n'
    }
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "reading standard input:", err)
    }
输入的行为可以通过一个函数控制，来控制输入的每个部分（参阅 SplitFunc 的文档），但是对于复杂的问题或持续传递错误的，可能还是需要原有接口。

*/

func main() {

}
