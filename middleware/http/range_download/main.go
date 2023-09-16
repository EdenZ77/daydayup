package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	fileURL := "https://ds.cloudapis.bigdata.gatorcloud.skyguardmis.com/securitySiteFulldb/security_fulldb_v2308040301.bin"
	filePath := "file.zip"

	err := downloadFile(fileURL, filePath)
	if err != nil {
		fmt.Println("下载文件出错:", err)
		return
	}

	fmt.Println("文件下载完成。")
}

func downloadFile(url string, filePath string) error {
	// 检查文件是否已存在
	_, err := os.Stat(filePath)
	if err == nil {
		fmt.Println("文件已存在，进行断点续传...")
	}

	// 打开文件以追加或创建
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// 获取文件已下载部分的大小
	fileInfo, _ := file.Stat()
	startOffset := fileInfo.Size()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	// 设置 Range 请求头以实现断点续传
	rangeHeader := fmt.Sprintf("bytes=%d-", startOffset)
	req.Header.Set("Range", rangeHeader)

	// 创建 HTTP 请求
	httpClient := NewHttpClient()
	// 发送请求并获取响应
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	contentLength := resp.ContentLength
	// 文件中的字节数：818554942
	// 当重新下载时，文件总字节数就是还没有被下载的数量
	fmt.Printf("当前文件剩余的总字节数=%d\n", contentLength)
	defer resp.Body.Close()

	// 检查服务器是否支持断点续传
	if resp.StatusCode != http.StatusPartialContent {
		fmt.Println("服务器不支持断点续传。")
		_, err = file.Seek(0, io.SeekStart) // 从文件开始处重新下载
		if err != nil {
			return err
		}
	} else {
		fmt.Println("服务器支持断点续传。")
	}

	// 创建一个缓冲区，并限制每次读取的字节数为100 000
	count := startOffset
	fmt.Printf("==============文件中已下载字节数fileSize=%d\n", startOffset)
	buffer := make([]byte, 100000)
	for {
		n, err := resp.Body.Read(buffer)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		// 写入文件的字节数限制为实际读取的字节数
		writeByte, err := file.Write(buffer[:n])
		if err != nil {
			return err
		}
		count = count + int64(writeByte)
		fmt.Printf("已下载字节count=%d\n", count)
	}
	// 刷盘
	if err := file.Sync(); err != nil {
		return err
	}
	info, err := file.Stat()
	if err != nil {
		return err
	}
	endFileSize := info.Size()
	fmt.Printf("下载完成后，文件大小：%d\n", endFileSize)

	return nil
}

func NewHttpClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr, Timeout: time.Hour * 7}
	return client
}
