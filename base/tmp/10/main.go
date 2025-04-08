package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	//testStr := "macdeMacBook-Air\\luming;@macdeMacBook-Air\\skyguard"
	//domain, logonName := ParseUserName(testStr)
	//fmt.Println(domain, "==", logonName)

	//str1 := "111;222;333"
	//str1 := "111;;333"
	//str1 := "111;;"
	//
	//split := strings.Split(str1, ";")
	//fmt.Println(split)
	//uuids := []string{}
	//uuidsStr := strings.Join(nil, ",")
	//fmt.Println(uuidsStr)

	cidr, err := IPAndMaskToCIDR("172.30.3.121", "255.255.255.252")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("CIDR:", cidr) // 输出: 172.30.3.121/24
}

// ParseUserName 解析用户名称为domain和logonName
func ParseUserName(userName string) (string, string) {
	// 这是因为uss表中，存在部分数据不规则，具有;符号
	split := strings.Split(userName, ";")
	// 根据\进行切分
	userNameSplits := strings.Split(split[0], "\\")
	// 某些数据只有一个logonName
	if len(userNameSplits) == 1 {
		return "", userNameSplits[0]
	}
	return userNameSplits[0], userNameSplits[1]
}

// 根据 IP 和子网掩码返回 CIDR 格式（如 "172.30.3.121/24"）
func IPAndMaskToCIDR(ipStr, maskStr string) (string, error) {
	// 解析 IP 地址
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return "", fmt.Errorf("invalid IP address: %s", ipStr)
	}

	// 解析子网掩码
	mask := net.ParseIP(maskStr)
	if mask == nil {
		return "", fmt.Errorf("invalid subnet mask: %s", maskStr)
	}

	// 将子网掩码转换为 4 字节格式
	maskBytes := mask.To4()
	if maskBytes == nil {
		return "", fmt.Errorf("subnet mask is not IPv4: %s", maskStr)
	}

	// 计算前缀长度（如 255.255.255.0 -> 24）
	prefixLen, _ := net.IPv4Mask(maskBytes[0], maskBytes[1], maskBytes[2], maskBytes[3]).Size()
	if prefixLen == 0 {
		return "", fmt.Errorf("invalid subnet mask: %s", maskStr)
	}

	// 返回 CIDR 格式
	return fmt.Sprintf("%s/%d", ipStr, prefixLen), nil
}
