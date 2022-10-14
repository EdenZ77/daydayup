package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/pkg/errors"

	"github.com/olivere/elastic"
)

// EsDblResult 适用于 spam_dbl_v1 和 spam_dbl_all
type EsDblResult struct {
	ScrollId string          `json:"_scroll_id"`
	Hits     *OutsideDblHits `json:"hits"`
}

type OutsideDblHits struct {
	Total int64            `json:"total"`
	Hits  []*InsideDblHits `json:"hits"`
}

type InsideDblHits struct {
	Index  string    `json:"_index"`
	Id     string    `json:"_id"`
	Source DblSource `json:"_source"`
}

type DblSource struct {
	Accessible int    `json:"accessible"`
	InitSrc    string `json:"init_src"`
	InitCat    string `json:"init_cat"`
	IsDomain   int    `json:"is_domain"`
	Site       string `json:"site"`
	Cat        int    `json:"cat"`
	Source     string `json:"source"`
	Ctime      string `json:"ctime"`
	Utime      string `json:"utime"`
	IsDelete   int    `json:"is_delete"`
}

type Security struct {
	Type  string `json:"type"`
	Site  string `json:"site"`
	Url   string `json:"url"`
	Utime string `json:"utime"`
}

func NewESClient() (*elastic.Client, error) {
	var optitons []elastic.ClientOptionFunc
	optitons = append(optitons, elastic.SetURL("http://172.30.3.116:9200"))
	optitons = append(optitons, elastic.SetBasicAuth("admin", "skg2102stg"))
	return elastic.NewClient(optitons...)
}

func main() {
	scroll, err := TestScroll02()
	if err != nil {
		return
	}
	fmt.Println(scroll)
}

func TestScroll() ([]*Security, error) {
	esClient, _ := NewESClient()

	query := elastic.NewBoolQuery().Must(elastic.NewRangeQuery("utime").Gte("2000-11-26T07:08:33.165870"))

	svc := esClient.Scroll("spam_dbl_v1").Query(query).Sort("utime", true).Size(10)

	securitySlice := make([]*Security, 0, 16)
	for {
		res, err := svc.Do(context.Background())
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, errors.Wrap(err, "es scroll error")
		}
		for index, hit := range res.Hits.Hits {
			var esSecur Security
			if err := json.Unmarshal(*hit.Source, &esSecur); err != nil {
				return nil, errors.Wrap(err, "json Unmarshal error")
			}
			fmt.Println("=======", index)
			fmt.Println(esSecur)
			securitySlice = append(securitySlice, &esSecur)
		}
		if len(res.ScrollId) == 0 {
			return nil, errors.Errorf("expected scrollId in results; got %q", res.ScrollId)
		}
	}
	if err := svc.Clear(context.Background()); err != nil {
		return nil, errors.Wrap(err, "es clean scroll error")
	}
	return securitySlice, nil

}

func TestScroll02() ([]*DblSource, error) {
	esClient, _ := NewESClient()

	query := elastic.NewBoolQuery().Must(elastic.NewRangeQuery("utime").Lte("2021-03-28T01:08:33.165870"))
	query.Must(elastic.NewTermQuery("source", "sophos"))

	svc := esClient.Scroll("spam_dbl_v1").Query(query).Size(20)

	securitySlice := make([]*DblSource, 0, 32)
	for {
		res, err := svc.Do(context.Background())
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, errors.Wrap(err, "es scroll error")
		}
		for index, hit := range res.Hits.Hits {
			var esSecur DblSource
			if err := json.Unmarshal(*hit.Source, &esSecur); err != nil {
				return nil, errors.Wrap(err, "json Unmarshal error")
			}
			fmt.Println("=======", index)
			fmt.Println(esSecur)
			securitySlice = append(securitySlice, &esSecur)
		}
		if len(res.ScrollId) == 0 {
			return nil, errors.Errorf("expected scrollId in results; got %q", res.ScrollId)
		}
	}
	if err := svc.Clear(context.Background()); err != nil {
		return nil, errors.Wrap(err, "es clean scroll error")
	}
	return securitySlice, nil

}
