package main

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	data, err := ioutil.ReadFile("file/AD_v2205100003.bin.ext")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	//
	fmt.Println(string(data[0:24]))

	out, err := os.Create("file/test111.zip")
	if err != nil {
		fmt.Println("Create error: ", err)
		return
	}

	defer out.Close()

	writer := zip.NewWriter(out)

	//var files = []struct {
	//	Name, Body string
	//}{
	//	{"1.txt", "01 1 file"}, // 这里可以用其他方式获取文件的名称和内容即可
	//	//{"2.txt", "01 2 file"},
	//}

	fileWriter, err := writer.Create("rr")
	if err != nil {
		if os.IsPermission(err) {
			fmt.Println("权限不足: ", err)
			return
		}
		fmt.Printf("Create file %s error: %s\n", "rr", err.Error())
		return
	}

	fileWriter.Write(data[24:])

	if err := writer.Close(); err != nil {
		fmt.Println("Close error: ", err)
	}
}
