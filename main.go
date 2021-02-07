package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func main() {
	// Go to https://hacker-news.firebaseio.com/v0/topstories.json, get the top 25
	var response, err = http.Get("https://hacker-news.firebaseio.com/v0/topstories.json")
	if err != nil {
		panic(err)
	}
	var nums []int
	err = json.Unmarshal(getResponseBody(response), &nums)
	if err != nil {
		panic(err)
	}
	nums = nums[:25]
	fmt.Println(nums)
	// For each of those, go to https://hacker-news.firebaseio.com/v0/item/$id.json and get the title
	for i := 0; i < len(nums); i++ {

	}
	startTime := time.Now()
	wg := sync.WaitGroup{}
	newsChannel := make(chan NewsItem, 25)
	for index, value := range nums {
		go getHackerNewsItem(index, value, &wg, newsChannel)
	}
	wg.Wait()
	close(newsChannel)

	for item := range newsChannel {
		fmt.Printf("%v: %v\n", item.Index, item.Title)
	}
	endTime := time.Since(startTime)
	fmt.Printf("Took: %v\n", endTime)
}

func getResponseBody(r *http.Response) []byte {
	bytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		panic(err)
	}
	return bytes
}

func getHackerNewsItem(index, value int, wg *sync.WaitGroup, channel chan NewsItem) {
	wg.Add(1)
	var response, err = http.Get(fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json", value))
	if err != nil {
		panic(err)
	}
	var news NewsItem
	b := getResponseBody(response)
	err = json.Unmarshal(b, &news)
	if err != nil {
		panic(err)
	}
	news.Index = index
	wg.Done()
	channel <- news
}

type NewsItem struct {
	Title string `json:"title"`
	Index int
}


