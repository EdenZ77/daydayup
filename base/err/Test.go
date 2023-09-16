package main

import (
	"fmt"
)

func main() {

	err := A()
	if err != nil {
		fmt.Println("======打印普通信息")
		fmt.Println(PrintMessage(err))
		fmt.Println("======打印附带堆栈信息")
		fmt.Println(PrintStack(err))
		//fmt.Println("======打印根因errors.Cause")
		//fmt.Printf("original error:%T ==> %v\n", errors.Cause(err), errors.Cause(err))

		//fmt.Println("======errors.Is和errors.As判断")
		//if errors.Is(err, FSError) {
		//	fmt.Println("err == FSError")
		//}
		//if errors.As(err, &FSError) {
		//	fmt.Println("err as  FSError")
		//
		//}

		return
	}
	fmt.Println("ok")
}
func PrintMessage(err error) string {
	return err.Error()
}

func PrintStack(err error) string {
	return fmt.Sprintf("%+v", err)
}
