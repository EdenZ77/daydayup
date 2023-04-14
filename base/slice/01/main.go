package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
)

func main() {
	//arr := []int32{8, 8, 1, 1, 1, 1, 3, 3, 4, 4, 5}
	//removeDuplic := RemoveDuplic(arr)
	//fmt.Println(removeDuplic)

	jsonStruct()
}

func RemoveDuplic(slice []int32) []int32 {
	var resultList []int32
	tempMap := map[int32]byte{}
	for _, item := range slice {
		tempLen := len(tempMap)
		tempMap[item] = 0
		if len(tempMap) != tempLen {
			resultList = append(resultList, item)
		}
	}
	return resultList
}

func jsonInter() {
	msg := map[string]interface{}{
		"cat":       3,
		"sourceCat": []string{"22", "33"},
		"source":    "sophos-add",
		"uuid":      "uuid==",
		"url":       "urlStr==",
		"url_type":  "sophosType==",
	}
	msgJsonStr, err := jsoniter.MarshalToString(msg)
	if err != nil {
		return
	}
	fmt.Println(msgJsonStr)
}

type RabbitmqMsg struct {
	Cat       int32    `json:"cat"`
	SourceCat []string `json:"sourceCat"`
	Source    string   `json:"source"`
	Uuid      string   `json:"uuid"`
	Url       string   `json:"url"`
	UrlType   string   `json:"url_type"`
}

func jsonStruct() {
	msg := RabbitmqMsg{
		Cat:       3,
		SourceCat: []string{"22", "33"},
		Source:    "sophos-add",
		Uuid:      "uuid==",
		Url:       "urlStr==",
		UrlType:   "sophosType==",
	}
	msgJsonStr, err := jsoniter.MarshalToString(msg)
	if err != nil {
		return
	}
	fmt.Println(msgJsonStr)
}
