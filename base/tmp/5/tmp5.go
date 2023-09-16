package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type ResponseStatus struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Reputation struct {
	Category  string   `json:"category"`
	Score     float32  `json:"score"`
	Tag       []string `json:"tag"`
	Timestamp string   `json:"timestamp"`
	SourceRef []string `json:"source_ref"`
}

type Label struct {
	Type       string        `json:"type"`
	Value      string        `json:"value"`
	Geo        string        `json:"geo"`
	Reputation []*Reputation `json:"reputation"`
}

type DataItem struct {
	Type           string   `json:"type"`
	Id             string   `json:"id"`
	Title          string   `json:"title"`
	CreatedBy      string   `json:"created_by"`
	CreatedTime    string   `json:"created_time"`
	ModifiedBy     string   `json:"modified_by"`
	ModifiedTime   string   `json:"modified_time"`
	Pattern        string   `json:"pattern"`
	StartTime      string   `json:"start_time"`
	EndTime        string   `json:"end_time"`
	SuggestedOfCoa string   `json:"suggested_of_coa"`
	Labels         []*Label `json:"labels"`
}

type IocsResponse struct {
	ResponseStatus *ResponseStatus `json:"response_status"`
	NextPage       string          `json:"nextpage"`
	ResponseData   []*DataItem     `json:"response_data"`
}

func main() {

	f, err := os.Open("C:\\Users\\zhuqiqi\\Desktop\\XXS\\iocs.json")
	//f, err := os.Open("C:\\Users\\zhuqiqi\\Desktop\\XXS\\test.json")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	buf := make([]byte, 0, 5*1024*1024)
	scanner.Buffer(buf, 5*1024*1024)
	for scanner.Scan() {
		var resp IocsResponse
		err = json.Unmarshal([]byte(scanner.Text()), &resp)
		if err != nil {
			fmt.Printf("Unmarshal error=== %+v", err)
			return
		}
		fmt.Println(resp.NextPage)
		fmt.Println("原始文本====")
		//fmt.Println(scanner.Text())
		//b, err2 := json.Marshal(resp)
		//if err2 != nil {
		//	return
		//}
		//fmt.Println("序列化====")
		//fmt.Println(string(b))
	}
	if scanner.Err() != nil {
		fmt.Printf("err: %+v", scanner.Err())
	}

}
