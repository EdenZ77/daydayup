package main

import (
	"fmt"
	"strings"
)

func main() {
	// 给定的 DN 字符串
	//dn := "CN=dc-user1,OU=dc-zhy-ou,OU=Domain Controllers,DC=swg-ad,DC=gatorcloud,DC=skyguardmis,DC=com"
	//dn := "CN=dc-user2,OU=bj-dc-ou,OU=Domain Controllers,DC=company,DC=net"
	//dn := "CN=web-server,OU=Web,OU=Servers,OU=Resources,DC=corp,DC=local"
	//dn := "CN=printer1,OU=Printers,DC=office,DC=internal,DC=com"
	//dn := "CN=user1,OU=Sales,OU=Departments,DC=na,DC=global,DC=company,DC=org"
	dn := "CN=admin,DC=admin,DC=local"
	// 构建所有可能的上级 OU 和 DC 组合
	dns := buildDNCombinations(dn)

	// 构建 IN 子句
	inClause := strings.Join(dns, "','")
	fmt.Println("inClause=========")
	fmt.Println(inClause)
}

func buildDNCombinations(dn string) []string {
	// 分割 DN 字符串为单独的部分
	parts := strings.Split(dn, ",")

	var combinations []string

	// 找到第一个 DC 的位置
	dcStartIndex := -1
	for i, part := range parts {
		if strings.HasPrefix(part, "DC=") {
			dcStartIndex = i
			break
		}
	}

	if dcStartIndex == -1 {
		return combinations
	}

	// 获取完整的 DC 路径
	dcPath := strings.Join(parts[dcStartIndex:], ",")

	// 从最具体的 OU 开始，逐层添加组合
	for i := 0; i < dcStartIndex; i++ {
		if strings.HasPrefix(parts[i], "OU=") {
			pathParts := append(parts[i:dcStartIndex], parts[dcStartIndex:]...)
			combinations = append(combinations, strings.Join(pathParts, ","))
		}
	}

	// 添加完整的 DC 路径
	combinations = append(combinations, dcPath)

	return combinations
}
