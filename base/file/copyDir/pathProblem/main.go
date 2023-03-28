package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func walk(fp string) {
	filepath.Walk(fp, func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}

func main() {
	var fp = "D:\\test"
	walk(fp)

	fmt.Println("-----------------------------------------------")
	fp = "D:/test"
	walk(fp)
}
