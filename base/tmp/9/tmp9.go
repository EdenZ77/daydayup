package main

import (
	"encoding/json"
	"fmt"
)

type mystring struct {
	name string
}

func main() {
	//s := []*mystring{{"11"}, {"33"}, {"22"}}
	//fmt.Println(s, len(s))
	//s = append(s, nil)
	//fmt.Println(s, len(s))

	//s := []int32{33, 44, 55}
	////sprintf := fmt.Sprintf("%s \t %d", "qq", s)
	//sprintf := fmt.Sprintf("%d", s)
	//replace := strings.Replace(sprintf, " ", ",", -1)
	//fmt.Println(replace)

	//s := "http://dropbox.com/s/sh8sdkotfpbrrmg/fattura n. 177 del  26-07-2016 .pdf.zip?dl=1 [3585]"
	//s := "http://dropbox.com/s/sh8sdkotfpbrrmg/fattura n. 177 del  26-07-2016 .pdf.zi [3585]"
	//s := "http://dropbox.com?s/sh8sdkotfpbrrmg/fattura n. 177 del  26-07-2016 .pdf.zi [3585]"
	//s := "http://dropbox.com?s/sh8sdkotfpbrrmg/fattura n. 177 del  26-07-2016 .pdf.zip?dl=1 [3585]"
	//s := "http://dropbox.com [3585]"
	//strArr := strings.Fields(strings.TrimSpace(s))
	//strArr = strArr[:len(strArr)-1]
	//url := strings.Join(strArr, "")
	//fmt.Println(url)
	//index := strings.Index(url, "://")
	//noProto := url
	//if index != -1 {
	//	noProto = noProto[index+3:]
	//}
	//fmt.Println(noProto)
	//// =======
	//indexXie := strings.Index(noProto, "/")
	//indexWen := strings.Index(noProto, "?")
	//noProtoAndPath := noProto
	//if indexXie == -1 && indexWen != -1 {
	//	noProtoAndPath = noProtoAndPath[:indexWen]
	//} else if indexXie != -1 && indexWen == -1 {
	//	noProtoAndPath = noProtoAndPath[:indexXie]
	//} else if indexXie != -1 && indexWen != -1 {
	//	if indexXie > indexWen {
	//		noProtoAndPath = noProtoAndPath[:indexWen]
	//	} else {
	//		noProtoAndPath = noProtoAndPath[:indexXie]
	//	}
	//}
	//fmt.Println(noProtoAndPath)

	//url := "https://www.outcastcrossfit.com"
	//
	//index := strings.Index(url, "://")
	//noProto := url
	//if index != -1 {
	//	noProto = noProto[index+3:]
	//}
	//fmt.Println("noProto: ", noProto)
	//indexXie := strings.Index(noProto, "/")
	//indexWen := strings.Index(noProto, "?")
	//noProtoAndPath := noProto
	//if indexXie == -1 && indexWen != -1 {
	//	noProtoAndPath = noProtoAndPath[:indexWen]
	//} else if indexXie != -1 && indexWen == -1 {
	//	noProtoAndPath = noProtoAndPath[:indexXie]
	//} else if indexXie != -1 && indexWen != -1 {
	//	if indexXie > indexWen {
	//		noProtoAndPath = noProtoAndPath[:indexWen]
	//	} else {
	//		noProtoAndPath = noProtoAndPath[:indexXie]
	//	}
	//}
	//fmt.Println("noProtoAndPath: ", noProtoAndPath)

	//cats := strings.Split("[2306,2562,2563]", ",")
	//cats := strings.Split("[2306]", ",")
	//fmt.Println(len(cats))
	//for _, cat := range cats {
	//	fmt.Println(cat)
	//}
	var catArr []int32
	_ = json.Unmarshal(([]byte)("[]"), &catArr)
	fmt.Println(catArr)
}

//const (
//	ErrParams = 10400
//	ErrServer = 10500
//)
//
//var msg = map[int]string{
//	ErrParams: "参数错误",
//	ErrServer: "系统内部错误",
//}
