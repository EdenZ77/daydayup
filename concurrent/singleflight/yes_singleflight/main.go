package main

import (
	"errors"
	"golang.org/x/sync/singleflight"
	"log"
	"sync"
)

var errorNotExist = errors.New("not exist")
var gsf singleflight.Group

// 模拟从cache中获取值，cache中无该值
func getDataFromCache(key string) (string, error) {
	return "", errorNotExist
}

// 模拟从数据库中获取值
func getDataFromDB(key string) (string, error) {
	log.Printf("get %s from database", key)
	return "data", nil
}

// 获取数据
func getData(key string) (string, error) {
	data, err := getDataFromCache(key)
	if err == errorNotExist {
		//模拟从db中获取数据
		// Do 执行函数, 对同一个 key 多次调用的时候，在第一次调用没有执行完的时候只会执行一次 fn 其他的调用会阻塞住等待这次调用返回
		// v, err 是传入的 fn 的返回值
		// shared 表示是否真正执行了 fn 返回的结果，还是返回的共享的结果
		v, err, _ := gsf.Do(key, func() (interface{}, error) {
			return getDataFromDB(key)
			//set cache
		})
		if err != nil {
			log.Println(err)
			return "", err
		}

		//TOOD: set cache
		data = v.(string)
	} else if err != nil {
		return "", err
	}
	return data, nil
}

func main() {
	var wg sync.WaitGroup
	wg.Add(10)

	//模拟10个并发
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			data, err := getData("key")
			if err != nil {
				log.Print(err)
				return
			}
			log.Println(data)
		}()
	}
	wg.Wait()
}
