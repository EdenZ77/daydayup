package main

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	DefaultTimeFormat = "2006-01-02 15:04:05"
)

type LocalTime time.Time

type DataVersionStatMsg struct {
	Type    string     `json:"type"`
	Version string     `json:"version"`
	Utime   *LocalTime `json:"utime"`
}

func (t LocalTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(DefaultTimeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, DefaultTimeFormat)
	b = append(b, '"')
	return b, nil
}

func (t *LocalTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+DefaultTimeFormat+`"`, string(data), time.Local)
	*t = LocalTime(now)
	return
}

func (t LocalTime) String() string {
	return time.Time(t).Format(DefaultTimeFormat)
}

func main() {
	var statMsg DataVersionStatMsg
	value := "{\"utime\":\"2022-06-16 18:20:51\",\"type\":\"SpamSA\",\"version\":\"2206160001\"}"
	if err := json.Unmarshal([]byte(value), &statMsg); err != nil {
		fmt.Println(err)
	}
	fmt.Println(statMsg.Utime)
}
